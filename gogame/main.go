package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(20)
	var number int
	var counter int

	for counter = 0; counter < 5; counter++ {
		fmt.Print("Tahmininizi girin: ")
		fmt.Scan(&number)

		if number < randomNumber {
			fmt.Println("Yanlis tuttugunuz sayi daha yuksek.")
		} else if number > randomNumber {
			fmt.Println("Yanlis tuttugunuz sayi daha alcak.")
		} else {
			fmt.Println("Dogru tahmin!!!!")
			break
		}
	}

	if counter == 5 {
		fmt.Println("Hakkin bitti :(")
	}
}
