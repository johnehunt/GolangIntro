package main

import "fmt"

type Worker interface {
	DoWork() string
}

type HomeWorker interface {
	StayHome() string
	Worker
}

type Person struct {
	Name string
}

func (p Person) DoWork() string {
	return fmt.Sprintf("%s working", p.Name)
}
func (p Person) StayHome() string {
	return fmt.Sprintf("%s is staying home", p.Name)
}

type StringableWorker interface {
	Worker
	fmt.Stringer
}

type Employee struct {
	Name string
	Task string
}

func (e Employee) DoWork() string {
	return fmt.Sprintf("working on %s", e.Task)
}
func (e Employee) String() string {
	return fmt.Sprintf("Employee(%s)", e.Name)
}

func main() {

	p1 := Person{"Denise"}
	fmt.Println("p1.DoWork():", p1.DoWork())
	fmt.Println("p1.StayHome():", p1.StayHome())

	e1 := Employee{"John", "Coding"}
	fmt.Println("e1:", e1)
	fmt.Println(e1.DoWork())
}
