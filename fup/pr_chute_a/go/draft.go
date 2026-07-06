package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	inferior := 0
	superior := 100
	secreto := rand.Intn(99) + 1

	var chute int

	for {
		fmt.Printf("Diga um numero entre ]%d, %d[: ", inferior, superior)
		fmt.Scan(&chute)

		if chute <= inferior || chute >= superior {
			continue
		}

		if chute == secreto {
			fmt.Printf("Era %d, você ganhou!\n", secreto)
			break
		}

		if chute < secreto {
			inferior = chute
		} else {
			superior = chute
		}

		if superior-inferior == 2 {
			fmt.Printf("Era %d, você perdeu!\n", secreto)
			break
		}
	}
}