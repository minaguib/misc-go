package main

import (
	"math/rand"
	"testing"
)

const SEED = 99

// Native size

func BenchmarkInt(b *testing.B) {
	r := rand.New(rand.NewSource(SEED))
	for i := 0; i < b.N; i++ {
		r.Int()
	}
}

func BenchmarkIntn(b *testing.B) {
	r := rand.New(rand.NewSource(SEED))
	for i := 0; i < b.N; i++ {
		r.Intn(1 << 29)
	}
}

// 31-bit

func BenchmarkInt31(b *testing.B) {
	r := rand.New(rand.NewSource(SEED))
	for i := 0; i < b.N; i++ {
		r.Int31()
	}
}

func BenchmarkInt31n(b *testing.B) {
	r := rand.New(rand.NewSource(SEED))
	for i := 0; i < b.N; i++ {
		r.Int31n(1 << 29)
	}
}

// 32-bit

func BenchmarkUint32(b *testing.B) {
	r := rand.New(rand.NewSource(SEED))
	for i := 0; i < b.N; i++ {
		r.Uint32()
	}
}

func BenchmarkFloat32(b *testing.B) {
	r := rand.New(rand.NewSource(SEED))
	for i := 0; i < b.N; i++ {
		r.Float32()
	}
}

// 63-bit

func BenchmarkInt63(b *testing.B) {
	r := rand.New(rand.NewSource(SEED))
	for i := 0; i < b.N; i++ {
		r.Int63()
	}
}

func BenchmarkInt63n(b *testing.B) {
	r := rand.New(rand.NewSource(SEED))
	for i := 0; i < b.N; i++ {
		r.Int63n(1 << 62)
	}
}

// 64-bit

func BenchmarkUint64(b *testing.B) {
	r := rand.New(rand.NewSource(SEED))
	for i := 0; i < b.N; i++ {
		r.Uint64()
	}
}

func BenchmarkFloat64(b *testing.B) {
	r := rand.New(rand.NewSource(SEED))
	for i := 0; i < b.N; i++ {
		r.Float64()
	}
}
