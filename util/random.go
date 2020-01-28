package util

import (
	"math/rand"
	"time"
)

// RandomString returns one of the given strings at random.
func RandomString(options []string) string {
	rand.Seed(time.Now().Unix())
	return options[rand.Intn(len(options))]
}
