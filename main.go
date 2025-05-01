package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 5
	usage    = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, harder it gets.

Wanna play?
`
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	// Modern seeding method
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	for turn := 0; turn < maxTurns; turn++ {
		n := r.Intn(guess + 1)
		fmt.Printf("Turn %d: Generated %d\n", turn+1, n)

		if n == guess {
			switch r.Intn(5) {
			case 0:
				fmt.Println("🎉 YOU WIN! Fortune is on your side!")
			case 1:
				fmt.Println("👏 Bravo! You nailed it!")
			case 2:
				fmt.Println("🌈 Luck has smiled on you today!")
			case 3:
				fmt.Println("🔥 Incredible! You guessed it right!")
			case 4:
				fmt.Println("🏆 Victory! You’ve beaten the odds!")
			}

			if turn == 0 {
				fmt.Println("🌟 Amazing! You guessed it on the first try!")
			}
			return
		}
	}

	switch r.Intn(5) {
	case 0:
		fmt.Println("☠️ YOU LOST... Better luck next time!")
	case 1:
		fmt.Println("💀 So close, yet so far!")
	case 2:
		fmt.Println("😢 Not this time. Try again?")
	case 3:
		fmt.Println("👻 The lucky number slipped away...")
	case 4:
		fmt.Println("🙈 No match. Don’t give up!")
	}
}
