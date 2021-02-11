package main

import (
	"context"
	"log"
	"os"
	"os/signal"
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

func tryToBecomeLeader() (bool, func() (bool, error), error) {
	client := goredislib.NewClient(&goredislib.Options{
		Addr: "localhost:6379",
	})
	pool := goredis.NewPool(client)
	rs := redsync.New(pool)

	mutex := rs.NewMutex(mutexName)

	if err := mutex.Lock(); err != nil {
		client.Close()
		return false, nil, err
	}

	return true, func() (bool, error) {
		return mutex.Unlock()
	}, nil
}

func doElectionAndMaintainTheStatus(quit <-chan struct{}) {
	ticker := time.NewTicker(time.Second * 5)
	var err error
	var ok bool
	var cf func() (bool, error)

	c := goredislib.NewClient(&goredislib.Options{
		Addr: "localhost:6379",
	})
	defer c.Close()
	for {
		select {
		case <-ticker.C:
			if atomic.LoadInt64(&isLeader) == 0 {
				ok, cf, err = tryToBecomeLeader()
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
	// subscribe the expire event of redis
	c := goredislib.NewClient(&goredislib.Options{
		Addr: "localhost:6379"})
	defer c.Close()

	ctx := context.Background()
	pubsub := c.Subscribe(ctx, redisKeyExpiredEventSubj)
	_, err := pubsub.Receive(ctx)
	if err != nil {
		log.Printf("prog-%s subscribe expire event failed: %s\n", id, err)
		return
	}
	log.Printf("prog-%s subscribe expire event ok\n", id)

	// Go channel which receives messages from redis db
	ch := pubsub.Channel()
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

	go func() {
		doElectionAndMaintainTheStatus(quit)
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
