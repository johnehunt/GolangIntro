package calc

import "testing"

// TestCalcAddOneAndOne checks addition
func TestCalcAddOneAndOne(t *testing.T) {
	v, err := calculator("+", 1, 1)
	if v != 2 || err != nil {
		t.Error("Expected 2, got ", v)
	}
}

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
		v, err := calculator(data.operation, data.xvalue, data.yvalue)
		if v != data.result || err != nil {
			t.Error("Expected", data.result, ", got ", v)
		}
	}
}
