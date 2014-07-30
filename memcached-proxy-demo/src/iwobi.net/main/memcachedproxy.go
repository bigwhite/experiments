/*
 * Copyright 2012 iWobi.net. All rights reserved.
 *
 * A counterproxy server
 */

package main

import (
	"fmt"
	"io"
	"iwobi.net/libmemcached"
	"iwobi.net/proto"
	"iwobi.net/utils/ringbuf"
	"net"
	"os"
	"runtime"
	"time"
)

const (
	EXPIRE_TIME = "120" /* default key expired time */
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func incr(mc *libmemcached.MemcachedConn, key string, value string, tmstamp string, exptime string) (string, error) {
	newKey := key + tmstamp /* e.g. 1258020120906105105 */
	//fmt.Printf("new key = %s\n", newKey)

	var out string
	err := libmemcached.Incr(mc, newKey, value, &out)
	if err != nil {
		if err == libmemcached.ErrNotFound {
			err = libmemcached.Add(mc, newKey, "0", exptime)
			if err != nil {
				if err == libmemcached.ErrExists {
					err = libmemcached.Incr(mc, newKey, value, &out)
					if err != nil {
						fmt.Println(err)
						return "", err
					}
				} else {
					fmt.Println(err)
					return "", err
				}
			} else {
				err = libmemcached.Incr(mc, newKey, value, &out)
				if err != nil {
					fmt.Println(err)
					return "", err
				}
			}
		}
	}

	return out, nil
}

func conHandler(c net.Conn) {
	defer c.Close()
	fmt.Printf("[counterproxy] - Accept connect: %v\n", c)

	mc, err := libmemcached.Connect("localhost", "11211")
	if err != nil {
		fmt.Println(err)
		fmt.Printf("[counterproxy] - Connect Memcached Server failed!\n")
		return
	}
	defer libmemcached.Close(mc)

	var buf [512]byte
	rb := ringbuf.New(512)

	for {
		n, err := c.Read(buf[:])

		if err != nil {
			fmt.Println(err)
			if err == io.EOF {
				fmt.Printf("[counterproxy] - Client [%v] shutdown!\n", c)
			}
			break
		}

		_, err = rb.Write(buf[:], n)
		if err != nil {
			fmt.Println(err)
			break
		}

		packages, err := proto.Parse(rb)
		if err != nil && err != proto.ErrInComplete {
			fmt.Println(err)
			fmt.Printf("[counterproxy] - Close the connection: %v\n", c)
			return
		}

		for i := range packages {
			t := time.Now()
			tmstamp := fmt.Sprintf("%d%02d%02d%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
			out, err := incr(mc, packages[i].Key, packages[i].Val, tmstamp, EXPIRE_TIME)
			if err != nil {
				fmt.Println(err)
				fmt.Printf("[counterproxy] - Incr err!\n")
				return
			}

			_, err = c.Write([]byte(tmstamp + " " + out + "\r\n"))
			if err != nil {
				fmt.Println(err)
				fmt.Printf("[counterproxy] - Send response err!\n")
				return
			}
		}
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	service := ":12580"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "[counterproxy] - Accept error: %s", err.Error())
			continue
		}
		go conHandler(conn)
	}
}
