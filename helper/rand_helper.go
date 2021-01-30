package helper

import (
	"math/rand"
	"time"
)

func GenerateUintNumber(n int) uint64 {
	rand.Seed(time.Now().UnixNano())
	return uint64(rand.Intn(n))
}

func GenerateFloatNumber() float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()
}

func GenerateFloatNumberRange(a float64) float64 {
	rand.Seed(time.Now().UnixNano())
	f := rand.Float64()
	return f * a
}

// Perm returns, as a slice of n ints, a pseudo-random permutation of the integers [0,n).
func Perm(n int) []float64 {
	m := make([]float64, n)
	for i := 0; i < n; i++ {
		j := rand.Intn(i + 1)
		m[i] = m[j]
		m[j] = float64(i)
	}
	return m
}
