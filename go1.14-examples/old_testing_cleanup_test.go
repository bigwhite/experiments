package main

import "testing"

func setup(t *testing.T) func() {
	t.Logf("setup before test")
	return func() {
		t.Logf("teardown/cleanup after test")
	}
}

func TestCase1(t *testing.T) {
	f := setup(t)
	defer f()
	t.Logf("test the testcase")
}
