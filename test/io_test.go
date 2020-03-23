package test

import (
	"strconv"
	"testing"
)

func BenchmarkLoop(b *testing.B) {
	n := 0
	for i := 0; i < b.N; i++ {
		RadishGo([]byte(strconv.Itoa(n)), "localhost:8080")
		n++
	}
}
