package main

import (
	"crypto/sha256"
	"testing"

	"golang.org/x/crypto/scrypt"
)

func BenchmarkSHA256(b *testing.B) {
	b.ReportAllocs()
	data := []byte("hello world")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sha256.Sum256(data)
	}
}

func BenchmarkScrypt(b *testing.B) {
	b.ReportAllocs()
	const keyLen = 32
	data := []byte("hello world")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		scrypt.Key(data, data, 16384, 8, 1, keyLen)
	}
}
