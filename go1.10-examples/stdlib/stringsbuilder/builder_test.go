package builder

import "testing"

func BenchmarkBuildStringWithBytesBuffer(b *testing.B) {
	var	builder BuilderByBytesBuffer

	for i := 0; i < b.N; i++ {
		builder.WriteString("Hello, ")
		builder.WriteString("Go")
		builder.WriteString("-1.10")
		_ = builder.String()
	}

}
func BenchmarkBuildStringWithStringsBuilder(b *testing.B) {

	var	builder BuilderByStringsBuilder

	for i := 0; i < b.N; i++ {
		builder.WriteString("Hello, ")
		builder.WriteString("Go")
		builder.WriteString("-1.10")
		_ = builder.String()
	}
}
