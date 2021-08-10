package main

import "fmt"

type Talker interface {
	Talk()
}

type Speaker struct {
	Name string
}

func (s Speaker) Talk() {
	fmt.Printf("%s is talking", s.Name)
}

func main() {
	fmt.Println("Starting")

	speaker := Speaker{Name: "Adam"}

	var talker Talker = speaker
	talker.Talk()

	talker = &speaker
	talker.Talk()

	// var tpr *Talker = &speaker // Won't compile

	fmt.Println("Done")
}
