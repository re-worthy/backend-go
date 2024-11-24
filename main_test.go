package main

import (
	random "math/rand"
	"testing"
)

func TestRandom(t *testing.T) {
	if x := random.Intn(100); x < 30 {
		t.Fatalf("You got bumped... by number: %d", x)
	}
}
