package calc

import (
	"fmt"
	"testing"
)

// TestCalcAddOneAndOne checks addition
func TestCalcAddOneAndOne(t *testing.T) {
	v, err := Calculator("+", 1, 1)
	if v != 2 || err != nil {
		t.Error("Expected 2, got ", v)
	}
}

// Benchmark test
func BenchmarkCalc1000DividedBy3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v, err := Calculator("/", 1000, 3)
		if v != 333 || err != nil {
			b.Error("Expected 333, got ", v)
		}
	}
}

// Skipping a Test
func TestToCheckExp(t *testing.T) {
	t.Skip("skipping test is not implemented yet.")
}

// Parametric tests
type testdata struct {
	operation string
	xvalue    int
	yvalue    int
	result    int
}

func TestCalcParametricTest(t *testing.T) {

	var tests = []testdata{
		{"+", 1, 1, 2},
		{"-", 1, 1, 0},
		{"*", 1, 1, 1},
		{"/", 1, 1, 1},
	}

	for _, data := range tests {
		v, err := Calculator(data.operation, data.xvalue, data.yvalue)
		if v != data.result || err != nil {
			t.Error("Expected", data.result, ", got ", v)
		}
	}
}

// Example function test
func ExampleCalculator() {
	r, _ := Calculator("+", 1, 1)
	fmt.Println("Calculator('+', 1, 1):", r)
	// Output: Calculator('+', 1, 1): 2
}

// Subtests
func TestCalculatorPlusOperations(t *testing.T) {
	// <setup code>
	t.Run("AddZeroAndZero", func(t *testing.T) {
		v, err := Calculator("+", 0, 0)
		if v != 0 || err != nil {
			t.Error("Expected 0, got ", v)
		}
	})
	t.Run("AddOneAndOne", func(t *testing.T) {
		v, err := Calculator("+", 1, 1)
		if v != 2 || err != nil {
			t.Error("Expected 2, got ", v)
		}
	})
	t.Run("AddOneAndMinusOne", func(t *testing.T) {
		v, err := Calculator("+", 1, -1)
		if v != 0 || err != nil {
			t.Error("Expected 0, got ", v)
		}
	})
	// <tear-down code>
}

// Main test with start up and shut down behaviour
func TestMain(m *testing.M) {
	fmt.Println("Setup behaviour")
	m.Run()
	fmt.Println("Shutdown behaviour")
}
