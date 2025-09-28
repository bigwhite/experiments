package main

import (
	"context"
	"testing"
	"testing/synctest" // 假设这是Go 1.24的旧版本
)

func Test(t *testing.T) {
        synctest.Run(func() {
                _, cancel := context.WithCancel(t.Context())
                defer cancel()
        })
}
