package main

import "fmt"

// Worker interface
type Worker interface {
	doWork()
}

// HardWorker struct that implements the
// Worker interface
type HardWorker struct {
	industry string
}

func (h HardWorker) doWork() {
	fmt.Printf("Working in %s", h.industry)
}

// Person struct with an embedded Worker
type Person struct {
	name string
	Worker
}

func newPerson(name string, industry string) *Person {
	var hw = HardWorker{industry}
	var p = Person{name, hw}
	return &p
}

func main() {
	var p1 = newPerson("John", "Technology")
	fmt.Println("p1:", *p1)
	p1.doWork()
}
