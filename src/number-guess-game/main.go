package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
)

const MAX_GUESS_RANGE = 10
const MAX_NUMBER_OF_GUESSES = 4
const YES = "y"
const NO = "n"

func printInstructions() {
	fmt.Println("The aim of this game is to guess a number generated by")
	fmt.Printf("the computer. The number will be in the range 0 to %d.\n", MAX_GUESS_RANGE)
	fmt.Printf("You will get %d goes to guess the number\n", MAX_NUMBER_OF_GUESSES)
}

func getIntInput(prompt string) int {
	var i int
	for {
		fmt.Print(prompt)
		_, err := fmt.Scan(&i)
		if err != nil {
			fmt.Println("Input must be a valid integer, please try again")
			log.Print("Scan for i failed, due to ", err)
		} else {
			break
		}
	}
	return i
}

func getUserName() string {
	var name string
	fmt.Print("Please enter your name: ")
	fmt.Scanf("%s", &name)
	return name
}

func getYesOrNoInput(prompt string) string {
	var input string
	for {
		fmt.Print(prompt)
		fmt.Scanf("%s", &input)
		input = strings.ToLower(input)
		if input == YES || input == NO {
			break
		} else {
			fmt.Printf("Error - Input must be %s or %s, please try again")
		}
	}
	return input
}

func main() {
	fmt.Println("Welcome to the Number Guess Game")
	log.Println("Program starting")

	var numberToGuess int
	var guess int
	var name string
	var guessCount int

	name = getUserName()

	userInstructionPrompt := fmt.Sprintf("%s do you want to see the instructions? ", name)
	if getYesOrNoInput(userInstructionPrompt) == YES {
		printInstructions()
	}

	for {
		guessCount = 0
		numberToGuess = rand.Intn(MAX_GUESS_RANGE)
		for {
			guess = getIntInput("Please enter your guess: ")

			if guess == -1 {
				// Cheat mode
				fmt.Printf("Cheat Mode - The number to guess is %d\n", numberToGuess)
			} else {
				// Check the reuslt of the guess
				guessCount++
				if guess == numberToGuess {
					fmt.Printf("Well done %s you guessed the number\n", name)
					fmt.Printf("You took %d attempts\n", guessCount)
					break
				} else if guess < numberToGuess {
					fmt.Println("Your guess is less than the number")
				} else {
					fmt.Println("Your guess is greater than the number")
				}
				// check for max number of guesses
				if guessCount == MAX_NUMBER_OF_GUESSES {
					fmt.Println("You Loose!!")
					fmt.Println("You have reached the maximum number of guesses")
					fmt.Printf("The number to guess was %d\n", numberToGuess)
					break
				}
			}
		}

		playAgain := getYesOrNoInput("Do you want to play again? ")
		if playAgain == NO {
			break
		}
	}

	fmt.Println("Done")
	log.Println("Program Completed")
}
