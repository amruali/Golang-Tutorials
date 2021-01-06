package main

import "testing"

func TestCalculateSumFrom1ToN(t *testing.T) {
	tests := []struct{
		input int
		expected int
	}{
		{100, 5050},
		{50, 1275},
		{70, 2485},
		{90, 4095},
	}
	for _, test := range tests{
		output := CalculateSumFrom1ToN(test.input)
		if output != test.expected{
			t.Errorf("Test-Case failed input : %d, Expected : %d, output: %d", test.input, test.expected, output)
		}
	}
}
