package main

import "fmt"

func getCalculator(amount int) func(int) int {
	return func(value int) int {
		return amount * value
	}
}

func main() {

	doubleIt := getCalculator(2)
	tripleIt := getCalculator(3)

	fmt.Println("doubleIt(2):", doubleIt(2))
	fmt.Println("doubleIt(3):", doubleIt(3))

	fmt.Println("tripleIt(2):", tripleIt(2))
	fmt.Println("tripleIt(3):", tripleIt(3))

}
