package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX_CHICKEN_PRICE int = 5

const MAX_TOFU_PRICE int = 3

func main() {
	var chickenChannel chan string = make(chan string)
	var tofuChannel chan string = make(chan string)
	var websites = []string{"website-1", "website-2", "website-3", "website-4", "website-5"}

	for i := range websites {
		go checkChickenPrice(websites[i], chickenChannel)
		go checkTofuPrice(websites[i], tofuChannel)
	}

	sendMessage(chickenChannel)
}

func checkChickenPrice(website string, chickenChannel chan string) {
	for {
		time.Sleep(1 * time.Second)
		var chickenPrice int = int(rand.Float32()) * 20

		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}

	close(chickenChannel)

	fmt.Println("The checkChickenPrice function is done.")
}

func checkTofuPrice(website string, tofuChannel chan string) {
	for {
		time.Sleep(1 * time.Second)
		var tofuPrice int = int(rand.Float32()) * 20

		if tofuPrice <= MAX_TOFU_PRICE {
			tofuChannel <- website
			break
		}
	}

	close(tofuChannel)

	fmt.Println("The checkTofuPrice function is done.")
}

func sendMessage(chickenChannel chan string) {
	fmt.Println("The sendMessage function is started.")

	for website := range chickenChannel {
		fmt.Println("The chicken price is good in", website)
	}

	fmt.Println("The sendMessage function is done.")
}
