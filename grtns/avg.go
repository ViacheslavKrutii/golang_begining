package grtns

import (
	"fmt"
	"math/rand"
)

func createRandNumToChan(n float32, ch1 chan float32) {
	var i float32
	for i = 0.0; i < n; i++ {
		num := rand.Intn(10)

		ch1 <- float32(num)

	}
	close(ch1)
}

func avarageNum(ch1, ch2 chan float32) {
	var sum float32
	sum = 0.0
	var i float32
	for num := range ch1 {
		i++
		sum += num

		avg := sum / i

		ch2 <- avg

	}
	close(ch2)

}

func printRandNum(ch2 chan float32) {
	for avg := range ch2 {
		fmt.Println("avg", avg)

	}

}

func Grtns() {
	ch1 := make(chan float32)
	ch2 := make(chan float32)

	go createRandNumToChan(6, ch1)
	go avarageNum(ch1, ch2)
	printRandNum(ch2)

}
