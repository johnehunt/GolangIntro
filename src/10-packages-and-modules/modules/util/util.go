package util

// Package names match the folders that hold them
// Can get around this - but don;t

// Optional leading article comments added
// have the format Type followed by explanation

// Go can generate documentation for packages in a similar way to the
// standard package documentation. In a terminal run this command:
// godoc golang-book/chapter11/math Average

// Distance Represents a distance such as 100 feet
type Distance int

// Celsius Represents a temperature
type Celsius float64

// MAX represents the maximum value
const MAX = 100

// Freezing used to represent the value for water to freeze
const Freezing Celsius = 0.0
