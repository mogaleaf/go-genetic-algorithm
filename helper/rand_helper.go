package helper

import (
	"math/rand"
	"time"
)

func GenerateUintNumber(n int) uint64 {
	rand.Seed(time.Now().UnixNano())
	return uint64(rand.Intn(n))
}
