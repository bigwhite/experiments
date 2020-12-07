// +build !trace

package main

func trace() func() {
	return func() {

	}
}
