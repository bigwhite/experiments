package main

import (
	"context"
	"testing"
	"testing/synctest" // 这是Go 1.25的新版本
)

func Test(t *testing.T) {
        synctest.Test(t, func(t *testing.T) {
                _, cancel := context.WithCancel(t.Context())
                defer cancel()
        })
}
