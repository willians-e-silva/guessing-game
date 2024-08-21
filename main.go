package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, Welcome to the guessing game")
	fmt.Println("A number between 0 and 100 will be generated, try to guess it!")

	scanner := bufio.NewScanner(os.Stdin)

	guesses := [10]int64{}
	x := rand.Int64N(101)

	for i := range guesses {
		fmt.Printf("Enter your guess:")
		scanner.Scan()
		guess := scanner.Text()
		guess = strings.TrimSpace(guess)

		guessInt, err := strconv.ParseInt(guess, 10, 64)
		if err != nil {
			fmt.Println("Invalid guess, please enter a valid number")
			continue
		}

		switch {
		case guessInt < x:
			fmt.Println("Your guess is too low")
		case guessInt > x:
			fmt.Println("Your guess is too high")
		case guessInt == x:
			fmt.Println("Congratulations! You guessed the number")
			return
		}

		guesses[i] = guessInt
	}
}
