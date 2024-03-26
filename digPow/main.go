package main

import (
	"log"
	"math"
)

func getDigits(number int) []float64 {
	var digits []float64

	for number > 0 {
		d := number % 10
		digits = append(digits, float64(d))
		number /= 10
	}

	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return digits
}

func DigPow(n, p int) int {
	total := 0
	e := float64(p)

	for _, digit := range getDigits(n) {
		total += int(math.Pow(digit, e))
		e++
	}

	if total%n == 0 {
		return total / n
	} else {
		return -1
	}
}

func main() {
	log.Println(DigPow(695, 2))
	log.Println(DigPow(92, 1))
	log.Println(DigPow(46288, 3))
	log.Println(DigPow(3263, 1))
}
