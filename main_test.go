package main

import (
	"fmt"
	"fts/internal/strext"
	"github.com/kljensen/snowball/english"
	"testing"
)

func Test(t *testing.T) {
	for str := range strext.Fields("hello world    ") {
		fmt.Println(str)
	}
}

func BenchmarkStem(b *testing.B) {
	b.ReportAllocs()

	for range b.N {
		_ = english.Stem("accumulator", false)
	}
}
