package example05

import (
	example04 "learn_go/golang_intro_classical/hour15/6"
	"testing"
)

func BenchmarkStringFromAssignment(b *testing.B) {
	for n := 0; n < b.N; n++ {
		example04.StringFromAssignment(100)
	}
}

func BenchmarkStringFromBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		example04.StringFromBuffer(100)
	}
}
