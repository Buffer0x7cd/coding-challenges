package math

import "testing"

type addTest struct {
	arg1, arg2, want int
}

var addTests = []addTest{
	{2, 3, 5},
	{4, 6, 10},
	{0, 0, 0},
	{0, 1, 1},
	{3, 10, 12},
}

func TestAdd(t *testing.T) {
	for _, test := range addTests {
		got := Add(test.arg1, test.arg2)
		if got != test.want {
			t.Errorf("Add(%d, %d) = %d; want %d", test.arg1, test.arg2, got, test.want)
		}
	}
}