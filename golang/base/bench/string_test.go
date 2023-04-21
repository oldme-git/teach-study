package bench

import (
	"bytes"
	"testing"
)

func BenchmarkNormal(b *testing.B) {
	var (
		s  = "oldme"
		ns string
	)
	for i := 0; i < 100000; i++ {
		ns += s
	}
}

func BenchmarkBuf(b *testing.B) {
	var (
		s   = "oldme"
		buf bytes.Buffer
	)
	for i := 0; i < 100000; i++ {
		buf.WriteString(s)
	}
}
