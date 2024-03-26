package main

import (
	"log"
	"math"
)

func getDigits(number int) []int {
	var digits []int

	for number > 0 {
		d := number % 10
		digits = append(digits, d)
		number /= 10
	}

	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return digits
}

func DigPow(n, p int) int {
	digits := getDigits(n)

	res := 0
	for i, c := 0, p; i < len(digits); i, c = i+1, c+1 {
		res += int(math.Pow(float64(digits[i]), float64(c)))
	}

	out := res / n
	if out > 0 {
		return out
	}
	return -1
}

func main() {
	log.Println(DigPow(695, 2))
	log.Println(DigPow(92, 1))
	log.Println(DigPow(46288, 3))
	log.Println(DigPow(3263, 1))
}
