package concurrency

import (
	"fmt"
	"sync"
)

type answer struct {
	userId   string
	question uint
	answer   string
}

type score map[string]int

func roundGen() {

}

func player(name string, ch1 chan answer, wg *sync.WaitGroup) {
	defer wg.Done()
	ansString := ""
	fmt.Scan(&ansString)
	playerAns := answer{
		userId: name,
		answer: ansString}
	ch1 <- playerAns
}

func playerGen(wg *sync.WaitGroup) {
	ch1 := make(chan answer)
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go player(fmt.Sprintf("user%d", i), ch1, wg)
	}

}

func counter(correctAnswers map[int]string, ch1 chan answer, ch2 chan score) {
	newScore := make(score)
	for newAnswer := range ch1 {
		if correctAnswers[int(newAnswer.question)] == newAnswer.answer {
			newScore[newAnswer.userId] += 1
		}
	}
	ch2 <- newScore
}
