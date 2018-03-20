package main

import (
	"fmt"
	"math/rand"
	"time"
)

const iters = 1e10

// A round of the game show where the contestant always swaps
// if offered the opportunity.
func runSwap() bool {
	car := rand.Intn(3)
	guess := rand.Intn(3)

	// If the contestant guessed correctly first time then swapping guarantees they're swapping to a non-winning choice.
	if car == guess {
		return false
	}

	// If the contestant hasn't guessed correctly first time round, then swapping must guarantee them the win as the host can't reveal a door that has the car behind it.
	// Coding it out like this makes it seem really obvious!
	return true
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	success := 0
	for i:=0; i<iters; i++ {
		if runSwap() {
			success += 1
		}
	}
	
	fmt.Println(success, iters)
}