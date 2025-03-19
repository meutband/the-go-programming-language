// go test -bench=.
// BenchmarkPopCount-16            409717357                2.816 ns/op
// BenchmarkPopCountLoop-16        158704407                7.583 ns/op
// PASS
package popcount

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for b.Loop() {
		PopCount(100)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for b.Loop() {
		PopCountLoop(100)
	}
}
