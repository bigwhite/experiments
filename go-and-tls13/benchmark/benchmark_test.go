package main

import (
	"crypto/tls"
	"testing"
)

func tls12_dial() error {
	conf := &tls.Config{
		InsecureSkipVerify: true,
		MaxVersion:         tls.VersionTLS12,
	}

	conn, err := tls.Dial("tcp", "localhost:8443", conf)
	if err != nil {
		return err
	}
	conn.Close()
	return nil
}

func tls13_dial() error {
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", "localhost:8443", conf)
	if err != nil {
		return err
	}
	conn.Close()
	return nil
}

func BenchmarkTls13(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		err := tls13_dial()
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkTls12(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		err := tls12_dial()
		if err != nil {
			panic(err)
		}
	}
}
