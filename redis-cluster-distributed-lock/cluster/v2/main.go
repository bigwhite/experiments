package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

const (
	redisKeyExpiredEventSubj = `__keyevent@0__:expired`
	redisClusterMasters      = "localhost:30001,localhost:30002,localhost:30003,localhost:30004,localhost:30005,localhost:30006"
)

var (
	isLeader  int64
	m         atomic.Value
	id        string
	mutexName = "the-year-of-the-ox-2021"
)

func init() {
	if len(os.Args) < 2 {
		panic("args number is not correct")
	}
	id = os.Args[1]
}

func tryToBecomeLeader(client *goredislib.ClusterClient) (bool, func() (bool, error), error) {
	pool := goredis.NewPool(client)
	rs := redsync.New(pool)

	mutex := rs.NewMutex(mutexName)

	if err := mutex.Lock(); err != nil {
		return false, nil, err
	}

	return true, func() (bool, error) {
		return mutex.Unlock()
	}, nil
}

func doElectionAndMaintainTheStatus(c *goredislib.ClusterClient, quit <-chan struct{}) {
	ticker := time.NewTicker(time.Second * 5)
	var err error
	var ok bool
	var cf func() (bool, error)

	for {
		select {
		case <-ticker.C:
			if atomic.LoadInt64(&isLeader) == 0 {
				ok, cf, err = tryToBecomeLeader(c)
				if ok {
					log.Printf("prog-%s become leader successfully\n", id)
					atomic.StoreInt64(&isLeader, 1)
					defer cf()
				}
				if !ok || err != nil {
					log.Printf("prog-%s try to become leader failed: %s\n", id, err)
				}
			} else {
				log.Printf("prog-%s is the leader\n", id)
				// update the lock live time and maintain the leader status
				c.Expire(context.Background(), mutexName, 8*time.Second)
			}
		case <-quit:
			return
		}
	}
}

func doExpire(quit <-chan struct{}) {
	var ch = make(chan *goredislib.Message)
	nodes := strings.Split(redisClusterMasters, ",")

	for _, node := range nodes {
		node := node
		go func(quit <-chan struct{}) {
			c := goredislib.NewClient(&goredislib.Options{
				Addr: node})
			defer c.Close()

			// subscribe the expire event of redis
			ctx := context.Background()
			pubsub := c.Subscribe(ctx, redisKeyExpiredEventSubj)
			_, err := pubsub.Receive(ctx)
			if err != nil {
				log.Printf("prog-%s subscribe expire event of node[%s] failed: %s\n",
					id, node, err)
				return
			}
			log.Printf("prog-%s subscribe expire event of node[%s] ok\n", id, node)

			// Go channel which receives messages from redis db
			pch := pubsub.Channel()

			for {
				select {
				case event := <-pch:
					ch <- event
				case <-quit:
					return
				}
			}
		}(quit)
	}
	for {
		select {
		case event := <-ch:
			key := event.Payload
			if atomic.LoadInt64(&isLeader) == 0 {
				break
			}
			log.Printf("prog-%s 收到并处理一条过期消息[key:%s]", id, key)
		case <-quit:
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	var quit = make(chan struct{})
	client := goredislib.NewClusterClient(&goredislib.ClusterOptions{
		Addrs: strings.Split(redisClusterMasters, ",")})
	defer client.Close()

	go func() {
		doElectionAndMaintainTheStatus(client, quit)
		wg.Done()
	}()
	go func() {
		doExpire(quit)
		wg.Done()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	_ = <-c
	close(quit)
	log.Printf("recv exit signal...")
	wg.Wait()
	log.Printf("program exit ok")
}
