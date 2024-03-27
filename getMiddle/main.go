package main

import "fmt"

func GetMiddle(s string) string {
	if len(s)%2 > 0 {
		return string(s[len(s)/2])
	} else {
		return string(s[(len(s)/2)-1]) + string(s[(len(s)/2)])
	}
}

func main() {
	fmt.Println("test:", GetMiddle("test"))
	fmt.Println("testing", GetMiddle("testing"))
	fmt.Println("middle", GetMiddle("middle"))
	fmt.Println("A", GetMiddle("A"))
	fmt.Println("ta", GetMiddle("ta"))
}
