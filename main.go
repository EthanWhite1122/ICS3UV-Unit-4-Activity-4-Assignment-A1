//Author Ethan White
//Version 1.0.0
//Date 2025-12-10
//Fileoverview This program will come up with a random number and make the user guess it
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	target := rand.Intn(10) + 1
	guess := -1

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Guess a number between 1 and 10. Enter 0 to quit.")

	for guess != target && guess != 0 {
		fmt.Print("Enter a guess: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		guess = 0
		valid := true
		for index := 0; index < len(input); index++ {
			if input[index] < '0' || input[index] > '9' {
				valid = false
				break
			}
			guess = guess*10 + int(input[index]-'0')
		}

		if !valid {
			fmt.Println("Invalid input. Enter numbers only.")
			continue
		}

		if guess == 0 {
			fmt.Println("Program terminated.")
		} else if guess == target {
			fmt.Println("You win!")
		} else {
			fmt.Println("Wrong guess. Try again.")
		}
	}

	fmt.Println("The target number was:", target)

	fmt.Println("\nDone.")
}
