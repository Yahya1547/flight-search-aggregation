package utils

import (
	"math/rand"
	"time"
)

func RandomDelay(minMs, maxMs int) {
	delay := rand.Intn(maxMs-minMs+1) + minMs
	time.Sleep(time.Duration(delay) * time.Millisecond)
}