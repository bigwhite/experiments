package main

import (
	"testing"
	"time"
)

func shutdownMaker(processTm int) func(time.Duration) error {
	return func(time.Duration) error {
		time.Sleep(time.Second * time.Duration(processTm))
		return nil
	}
}

func TestConcurrencyShutdown(t *testing.T) {
	f1 := shutdownMaker(2)
	f2 := shutdownMaker(6)

	err := ConcurrencyShutdown(time.Duration(10)*time.Second, ShutdownerFunc(f1), ShutdownerFunc(f2))
	if err != nil {
		t.Errorf("want nil, actual: %s", err)
		return
	}

	err = ConcurrencyShutdown(time.Duration(4)*time.Second, ShutdownerFunc(f1), ShutdownerFunc(f2))
	if err == nil {
		t.Error("want timeout, actual nil")
		return
	}
}

func TestSequentialShutdown(t *testing.T) {
	f1 := shutdownMaker(2)
	f2 := shutdownMaker(6)

	err := SequentialShutdown(time.Duration(10)*time.Second, ShutdownerFunc(f1), ShutdownerFunc(f2))
	if err != nil {
		t.Errorf("want nil, actual: %s", err)
		return
	}

	err = SequentialShutdown(time.Duration(5)*time.Second, ShutdownerFunc(f1), ShutdownerFunc(f2))
	if err == nil {
		t.Error("want timeout, actual nil")
		return
	}
}
