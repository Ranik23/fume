package utils

import (
	"time"
	"math/rand"
)

func Contains(slice []string, element string) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}

func RandomInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(200 - 0 + 1) + 0 
}