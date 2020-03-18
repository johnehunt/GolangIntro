package main

import "fmt"

// Person struct has fields name and age
type Person struct {
	name string
	age  int
}

func newPerson(name string, age int) *Person {
	p := Person{name, age}
	return &p
}

func (p *Person) birthday() {
	fmt.Printf("Happy Birthday you were %d ", p.age)
	p.age++
	fmt.Printf("You are now %d\n", p.age)
}

func (p Person) calculatePay(rate float64, hours float64) float64 {
	fmt.Printf("%s you worked %0.2f hours at %0.2f\n", p.name, hours, rate)
	return rate * hours
}

func main() {
	var p1 = newPerson("John", 21)
	p1.birthday()
	var pay = p1.calculatePay(8.55, 12.5)
	fmt.Printf("You were paid %0.2f\n", pay)

	// Use a pointer receiver type to avoid copying on method calls
	refp1 := &Person{name: "Paul", age: 35}
	pay2 := refp1.calculatePay(8.55, 12.5)
	fmt.Printf("You were paid %0.2f\n", pay2)

}
