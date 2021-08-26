package main

import "fmt"

type Foo int

func (f Foo) Bar() {
	fmt.Printf("My receiver is %v\n", f)
}

func main() {
	a := Foo(46)
	a.Bar()
	b := Foo(51)
	Foo.Bar(b)
}
