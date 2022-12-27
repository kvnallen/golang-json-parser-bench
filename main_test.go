package main

import (
	"testing"
)

func Benchmark_parseGoJay(b *testing.B) {
	d := []byte(`{"id":1,"name":"gojay","email":"gojay@email.com"}`)
	for i := 0; i < b.N; i++ {
		parseGoJay(d)
	}
}

func Benchmark_parseJsonIter(b *testing.B) {
	d := []byte(`{"id":1,"name":"gojay","email":"gojay@email.com"}`)
	for i := 0; i < b.N; i++ {
		parseJsonIter(d)
	}
}

func Benchmark_parseJsonIterDefault(b *testing.B) {
	d := []byte(`{"id":1,"name":"gojay","email":"gojay@email.com"}`)
	for i := 0; i < b.N; i++ {
		parseJsonIterDefault(d)
	}
}

func Benchmark_parseJsonIterFastest(b *testing.B) {
	d := []byte(`{"id":1,"name":"gojay","email":"gojay@email.com"}`)
	for i := 0; i < b.N; i++ {
		parseJsonIterFastest(d)
	}
}

func Benchmark_parseSonic(b *testing.B) {
	d := []byte(`{"id":1,"name":"gojay","email":"gojay@email.com"}`)
	for i := 0; i < b.N; i++ {
		parseSonic(d)
	}
}

func Benchmark_parseJson(b *testing.B) {
	d := []byte(`{"id":1,"name":"gojay","email":"gojay@email.com"}`)
	for i := 0; i < b.N; i++ {
		parseJson(d)
	}
}
