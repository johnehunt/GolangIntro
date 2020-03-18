package main

import (
	"fmt"
	"os"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

func main() {
	fmt.Println("Starting")

	sweaters := Inventory{"wool", 17}
	tmpl, err :=
		template.
			New("test").
			Parse("{{.Count}} items are made of {{.Material}}\n")

	if err != nil {
		fmt.Println("Something went wrong with the template", err)
	}

	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		fmt.Println("Something went wrong with executing the template", err)
	}

	fmt.Println("Done")
}
