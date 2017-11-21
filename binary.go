package gostructures

import "strings"

func DoubleToBinary(double float64) string {
	bits := []string{"0."}

	bit := 0.5
	for len(bits) < 32 {
		if double == 0 {
			break
		}
		if double >= bit {
			double -= bit
			bits = append(bits, "1")
		} else {
			bits = append(bits, "0")
		}
		bit = bit * 0.5
	}
	return strings.Join(bits, "")
}
