// go test -bench=.
// BenchmarkPopCount-16            411001372                2.865 ns/op
// BenchmarkPopCountClear-16       408850522                2.967 ns/op
// PASS
package popcount

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for b.Loop() {
		PopCount(100)
	}
}

func BenchmarkPopCountClear(b *testing.B) {
	for b.Loop() {
		PopCountClear(100)
	}
}
