package util

// Celsius Represents a temperature
type Celsius float64

// MAX represents the maximum value
const MAX = 100

// Freezing used to represent the value for water to freeze
const Freezing Celsius = 0.0

func Printer(value interface{}) {
	fmt.Println(value)
}

type Model struct {
	ID   string
	Name string
}
