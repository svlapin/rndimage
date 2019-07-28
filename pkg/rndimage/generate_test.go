package rndimage

import "testing"

func BenchmarkGenerate1000(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Generate(1000, 1000)
	}
}
