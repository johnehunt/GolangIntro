package main

import "fmt"

// Worker interface
type Worker interface {
	doWork()
}

// Person struct with an embedded Worker
type Person struct {
	name string
	Worker
}

func newPerson(name string, worker Worker) *Person {
	var p = Person{name, worker}
	return &p
}

// HardWorker struct that implements the
// Worker interface
type HardWorker struct {
	industry string
}

func (h HardWorker) doWork() {
	fmt.Printf("Working in %s", h.industry)
}

func main() {
	var p1 = newPerson("John", HardWorker{"Technology"})
	fmt.Println("p1:", *p1)
	p1.doWork()
}
