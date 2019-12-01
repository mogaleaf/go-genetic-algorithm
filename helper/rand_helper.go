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
