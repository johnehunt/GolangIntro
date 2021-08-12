package calc_test

import (
	"calc"
	"fmt"
	"testing"
)

func TestCalcAddOneAndOne(t *testing.T) {
	v, err := calc.Calculator(calc.PLUS, 1, 1)
	if v != 2 || err != nil {
		t.Error("Expected 2, got ", v)
	}
}

func ExampleDivide() {
	result, err := calc.Calculator(calc.PLUS, 5, 2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// Output:
	// 7
}

func TestToCheckExp(t *testing.T) {
	t.Skip("skipping test is not implemented yet.")
}

func TestMain(m *testing.M) {
	// Do any start up
	fmt.Println("Setup behaviour")
	m.Run()
	// Do any shutdown behaviour
	fmt.Println("Shutdown behaviour")
}

func TestCalculatorPlusOperations(t *testing.T) {
	// <setup code>
	t.Run("AddZeroAndZero", func(t *testing.T) {
		v, err := calc.Calculator(calc.PLUS, 0, 0)
		if v != 0 || err != nil {
			t.Error("Expected 0, got ", v)
		}
	})
	t.Run("AddOneAndOne", func(t *testing.T) {
		v, err := calc.Calculator(calc.PLUS, 1, 1)
		if v != 2 || err != nil {
			t.Error("Expected 2, got ", v)
		}
	})
	t.Run("AddOneAndMinusOne", func(t *testing.T) {
		v, err := calc.Calculator(calc.PLUS, 1, -1)
		if v != 0 || err != nil {
			t.Error("Expected 0, got ", v)
		}
	})
	// <tear-down code>
}
