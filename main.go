package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	num := rand.IntN(greeting())
	var a int

	for {

		_, err := fmt.Scan(&a)
		if err != nil {
			fmt.Println("Input a value")
			return
		}
		result := compareNumber(num, a)

		fmt.Println(result)
		if result == "You have guessed!" {
			break
		}
	}
}

func compareNumber(r, input int) string {
	switch {
	case input == r:
		return "You have guessed!"
	case input > r:
		return "Too high!"
	case input < r:
		return "Too low!"
	}

	return "Something happened..."
}

func greeting() int {
	fmt.Println("Guess the number from 0 to ...")
	fmt.Println("Type your value here")
	var b int
	_, err := fmt.Scan(&b)
	if err != nil {
		fmt.Println("Input a value")
		return 0
	}

	fmt.Println("Start!")
	return b
}
