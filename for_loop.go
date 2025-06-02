package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())

	random_n := rand.New(source)

	//Generate a randome number between 1 and 100
	target := random_n.Intn(100)

	fmt.Println("Welcome to the game")
	fmt.Println("I have chosen a number between 1 and 100")
	fmt.Println("Can you guess what it is?")

	var guess int

	for {
		fmt.Println("Enter your guess")
		fmt.Scanln(&guess)

		if guess > target {
			fmt.Printf("Your guess is %d: It is greater than the target value\n", guess)
		} else if guess < target {
			fmt.Printf("Your guess is %d: It is lesser than the target value\n", guess)
		} else {
			fmt.Printf("Congratulations. You have guessed the correct value: %d\n", guess)
			break
		}
	}

	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// // iterate over collections
	// numbers := []int{1, 2, 3}
	// for index, value := range numbers {
	// 	fmt.Printf("Index: %d, Value: %d\n", index, value)
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println("Odd Number:", i)
	// 	if i == 4 {
	// 		break
	// 	}
	// }

	// rows := 5

	// for i := 1; i <= rows; i++ {
	// 	for j := 1; j <= rows-i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for k := 1; k <= 2*i-1; k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	// for i := range 10 {
	// 	fmt.Println(i)
	// }
	// fmt.Println("We have liftoff!")

	// i := 1
	// for i < 5 {
	// 	fmt.Println("Iteration:", i)
	// 	i++
	// }

}
