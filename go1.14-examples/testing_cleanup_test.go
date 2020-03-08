package main

import "testing"

func TestCase1(t *testing.T) {

	t.Run("A=1", func(t *testing.T) {
		t.Logf("subtest1 in testcase1")

	})
	t.Run("A=2", func(t *testing.T) {
		t.Logf("subtest2 in testcase1")
	})
	t.Cleanup(func() {
		t.Logf("cleanup1 in testcase1")
	})
	t.Cleanup(func() {
		t.Logf("cleanup2 in testcase1")
	})
}

func TestCase2(t *testing.T) {
	t.Cleanup(func() {
		t.Logf("cleanup1 in testcase2")
	})
	t.Cleanup(func() {
		t.Logf("cleanup2 in testcase2")
	})
}
