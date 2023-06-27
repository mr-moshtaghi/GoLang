package sqrt

import (
	"fmt"
	"testing"
)

func almostEqual(a, b float64) bool {
	if (a - b) < 0.001{
		return true
	}
	return false
}

func TestSimple(t *testing.T) {
	val, err := Sqrt(2)
	if err != nil {
		t.Fatalf("error in calculatin %s", err)
	}
	if !almostEqual(val, 1.414214) {
		t.Fatalf("bad value - %f", val)
	}

}

type testCase struct {
	value    float64
	expected float64
}

func TestMany(t *testing.T) {
	testCase := []testCase{
		{0.0, 0.0},
		{2.0, 1.414214},
		{100, 4},
	}
	for _, tc := range testCase {
		t.Run(fmt.Sprintf("%f", tc.value), func(t *testing.T) {
			out, err := Sqrt(tc.value)
			if err != nil {
				t.Fatal("error")
			}
			if !almostEqual(out, tc.expected) {
				t.Fatalf("%f != %f", out, tc.expected)
			}
		})
	}

}
