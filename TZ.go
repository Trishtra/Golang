package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var guess int
	rand.Seed(time.Now().UnixNano())
	answer := rand.Intn(10) + 1

	for count := 0; count < 5; count++ {

		fmt.Println("Угадайте число от 1 до 10!")

		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println("Вы ввели не число.")
			return
		}

		if guess < answer {
			fmt.Println("Загаданное число больше!")

		} else if guess > answer {
			fmt.Println("Загаданное число меньше!")

		} else {
			fmt.Printf("Вы угадали! Искомое число : %d", answer)
			return
		}

		fmt.Printf("Осталось %d попыток\n", 4-count)
	}

	fmt.Printf("Вы проиграли! Искомое число : %d\n", answer)
}
