package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	//change seed everytime to make it truly random
	rand.Seed(time.Now().UnixNano())
	secretNum := rand.Intn(maxNum)
	fmt.Println("The secret number is", secretNum) //use , instead of + to concat in Go

	fmt.Println("Please enter your guess")
	var guess int
	for {
		_, err := fmt.Scanf("%v\n", &guess)
		// if n != 1 {
		// 	fmt.Println("Please only enter one guess")
		// 	return
		// }
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			continue
		}
		fmt.Println("Your guess is", guess)

		if guess > secretNum {
			fmt.Println("Your guess is greater than the secret number. Try again!")
		} else if guess == secretNum {
			fmt.Println("Your guess is corret. Legend!")
			break
		} else {
			fmt.Println("Your guess is smaller than the secret number. Try again!")
		}
	}

}
