package main

import "fmt"

// Worker interface
type Worker interface {
	doWork()
}

type Animal interface {
	doPlay()
}

// WorkHorse has an embedded Animal and Worker
type WorkHorse struct {
	name string
	Animal
	Worker
}

type Donkey struct {
	colour string
}

func (d Donkey) doPlay() {
	fmt.Println("DoPlay")
}

type HardWorker struct {
	industry string
}

func (h HardWorker) doWork() {
	fmt.Printf("Working in %s", h.industry)
}

func main() {
	var wh = WorkHorse{"dobbin", Donkey{"gray"}, HardWorker{"tourism"}}
	fmt.Println("wh:", wh)
}
