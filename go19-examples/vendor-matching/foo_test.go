package foo

import "testing"

func TestEcho(t *testing.T) {
	in := "hello"
	out := Echo(in)
	if out != "hello" {
		t.Fatal("expected: %s, we got: %s", "hello", out)
	}
}
