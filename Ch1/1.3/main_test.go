// This file runs benchmark testing for each echo function from main.go
//
// To run all benchmarks... run `go test -bench=.`
//
// To run specific benchmark... run `go test -bench=<Name>`
//   - Echo1 for BenchmarkEcho1
//   - ...
package main

import "testing"

var test_args []string = []string{"these", "are", "args", "to", "test", "echo", "functions"}

func BenchmarkEcho1(b *testing.B) {
	for b.Loop() {
		echo1(test_args)
	}
}

func BenchmarkEcho2(b *testing.B) {
	for b.Loop() {
		echo2(test_args)
	}
}

func BenchmarkEcho3(b *testing.B) {
	for b.Loop() {
		echo3(test_args)
	}
}
