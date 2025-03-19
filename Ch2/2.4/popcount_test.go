// go test -bench=.
// BenchmarkPopCount-16            409454588                2.992 ns/op
// BenchmarkPopCountShift-16       35474401                31.93 ns/op
// PASS
package popcount

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for b.Loop() {
		PopCount(100)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for b.Loop() {
		PopCountShift(100)
	}
}
