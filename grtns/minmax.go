package grtns

import (
	"fmt"
	"math/rand"
)

// Випадкове значення bool
func randomBool() bool {
	return rand.Intn(2) == 0
}

// Генератор рандомних цілих чисел
func generateRandInt(rangeint int) int {
	num := rand.Intn(rangeint)
	if randomBool() {
		return num
	}
	return -num
}

// Перша горутина
func randIntToChan(rangeint, n int, ch1 chan int, ch2 chan []int) {
	for i := 0; i <= n; i++ {
		num := generateRandInt(rangeint)
		ch1 <- num
	}
	close(ch1)
	minmaxSlice := <-ch2
	close(ch2)
	fmt.Printf("min:%d\nmax:%d\n", minmaxSlice[0], minmaxSlice[1])
}

// Друга горутина
func minmax(ch1 chan int, ch2 chan []int) {
	// мін макс
	var minmaxSlice []int
	//Зчитування з каналу від першої горутини
	for num := range ch1 {
		if minmaxSlice == nil {
			minmaxSlice = []int{num, num}
		}
		switch {
		case num < minmaxSlice[0]:
			minmaxSlice[0] = num
		case num > minmaxSlice[1]:
			minmaxSlice[1] = num
		}

	}
	// Повернення мін макс в першу горутину через канал
	ch2 <- minmaxSlice

}

func MinMaxMain() {
	ch1 := make(chan int)
	ch2 := make(chan []int)
	done := make(chan bool) // Канал для сигналу про завершення

	// Запуск горутини randIntToChan
	go func() {
		randIntToChan(1000, 10, ch1, ch2)
		done <- false // Сигнал про завершення горутини randIntToChan
	}()

	// Запуск горутини minmax
	go minmax(ch1, ch2)

	<-done

}
