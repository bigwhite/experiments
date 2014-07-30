/*
 * Copyright 2012 iWobi.net. All rights reserved.
 *
 * test cases for main package
 */

package main

import (
	"fmt"
	"iwobi.net/libmemcached"
	"testing"
)

func TestXxx(t *testing.T) {

}

func BenchmarkIncr(b *testing.B) {
	b.StopTimer()
	//mc, err := libmemcached.Connect("10.10.126.187", "11212")
	mc, err := libmemcached.Connect("localhost", "11211")
	if err != nil {
		fmt.Println(err)
		fmt.Printf("[counterproxy] - Connect Memcached Server failed!\n")
		return
	}
	defer libmemcached.Close(mc)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, err := incr(mc, "12580", "1", "120")
		if err != nil {
			fmt.Println(err)
			fmt.Printf("[counterproxy] - Incr err!\n")
			return
		}
		//        fmt.Println("new value = ", out)
	}
}
