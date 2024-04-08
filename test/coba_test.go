package test

import (
	"fmt"
	"testing"
)

type Person struct {
	Age int
}

func ReturnSliceWithPointers(size int) []*Person {
	res := make([]*Person, size)

	for i := 0; i < size; i++ {
		res[i] = &Person{}
	}

	return res
}

func ReturnSliceWithStructs(size int) []Person {
	res := make([]Person, size)

	for i := 0; i < size; i++ {
		res[i] = Person{}
	}

	return res
}

func Benchmark_ReturnSliceWithPointers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReturnSliceWithPointers(10000)
	}
}

func Benchmark_ReturnSliceWithStructs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReturnSliceWithStructs(10000)
	}
}

func TestIni1(t *testing.T) {
	for i := 0; i < 2; i++ {
		var p Person
		fmt.Println(&p)
	}
}
