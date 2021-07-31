package benchmark

import (
	"os"
	"strings"
	"testing"
)

func BenchmarkStringsEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(os.Args, " ")
	}
}

func BenchmarkLoopEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s, sep string

		for _, arg := range os.Args {
			s += sep + arg
			sep = " "
		}
	}
}
