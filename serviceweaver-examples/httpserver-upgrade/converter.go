package main

import (
	"context"
	"strings"

	"github.com/ServiceWeaver/weaver"
)

// Converter component.
type Converter interface {
	ToUpper(context.Context, string) (string, error)
}

// Implementation of the Converter component.
type converter struct {
	weaver.Implements[Converter]
}

func (r *converter) ToUpper(_ context.Context, s string) (string, error) {
	return strings.ToUpper(s), nil
}
