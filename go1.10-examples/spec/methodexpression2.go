package main

import "errors"

var err = errors.New("I am an error")

func main() {
	(interface {
		Error() string
	}).Error(err)
}
