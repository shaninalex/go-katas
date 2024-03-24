package main

import "fmt"

func FindOdd(seq []int) int {
	duplicates := map[int]int{}

	// fill duplicates structure
	for _, i := range seq {
		_, ok := duplicates[i]
		if !ok {
			duplicates[i] = 1
		} else {
			duplicates[i] = duplicates[i] + 1
		}
	}

	fmt.Println(duplicates)
	// find odd
	for k, v := range duplicates {
		if v%2 > 0 {
			return k
		}
	}

	return 0
}

func main() {
	sequences := [][]int{
		{7},
		{0},
		{1, 1, 2},
		{0, 1, 0, 1, 0},
		{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1},
	}

	for _, seq := range sequences {
		fmt.Println(FindOdd(seq))
	}

}
