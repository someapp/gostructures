package gostructures

import "testing"

func TestDoubleToBinary(t *testing.T) {
	tests := []struct {
		name   string
		input  float64
		output string
	}{
		{
			name:   "1/2 -> 0.1",
			input:  0.5,
			output: "0.1",
		},

		{
			name:   "1/4 -> 0.01",
			input:  0.25,
			output: "0.01",
		},

		{
			name:   "1/8 -> 0.001",
			input:  0.125,
			output: "0.001",
		},

		{
			name:   "1/4+1/16 -> 0.0101",
			input:  0.3125,
			output: "0.0101",
		},

		{
			name:   "0.9999 -> 0.11111",
			input:  0.9999999999,
			output: "0.1111111111111111111111111111111",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := DoubleToBinary(tt.input)
			if actual != tt.output {
				t.Errorf("expected %v to give: %v, actual: %v", tt.input, tt.output, actual)
			}
		})
	}
}
