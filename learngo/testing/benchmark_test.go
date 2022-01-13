package testing

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"testing"
)

func BenchmarkSHA1(b *testing.B) {
	data := []byte("Go benchmarking test")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha1.Sum(data)
	}
	b.StopTimer()
}

func BenchmarkSHA256(b *testing.B) {
	data := []byte("Go benchmarking test")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha256.Sum256(data)
	}
	b.StopTimer()
}

func BenchmarkSHA512(b *testing.B) {
	data := []byte("Go benchmarking test")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha512.Sum512(data)
	}
	b.StopTimer()
}
