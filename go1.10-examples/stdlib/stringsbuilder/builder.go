package builder

import (
	"bytes"
	"strings"
)


type BuilderByBytesBuffer struct {
	b bytes.Buffer
}
func (b *BuilderByBytesBuffer) WriteString(s string) error {
	_, err := b.b.WriteString(s)
	return err
}
func (b *BuilderByBytesBuffer) String() string{
	return b.b.String()
}

type BuilderByStringsBuilder struct {
	b strings.Builder
}

func (b *BuilderByStringsBuilder) WriteString(s string) error {
	_, err := b.b.WriteString(s)
	return err
}
func (b *BuilderByStringsBuilder) String() string{
	return b.b.String()
}
