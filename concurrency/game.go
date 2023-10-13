package concurrency

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

type question string

type answer struct {
	userId       string
	numbQuestion uint
	answer       string
}

type score map[string]int

func roundGen(wg *sync.WaitGroup, questions []string, questionCH chan question) {
	for i := 0; i < len(questions); i++ {
		for j := 0; j < 10; j++ {
			questionCH <- question(questions[i])
		}
		wg.Wait()
	}
}

func player(ctx context.Context, wg *sync.WaitGroup, name string, answerCH chan answer, questionCH chan question) {
	select {
	case <-questionCH:
		ansString := strconv.Itoa(rand.Intn(9))
		playerAns := answer{
			userId: name,
			answer: ansString}
		answerCH <- playerAns
		wg.Done()
	case <-ctx.Done():
		return
	}
}

func playerGen(ctx context.Context, wg *sync.WaitGroup, answerCH chan answer, questionCH chan question) {
	for i := 1; i <= 10; i++ {

		go player(ctx, wg, fmt.Sprintf("user%d", i), answerCH, questionCH)
	}
	wg.Wait()
}

func counter(correctAnswers map[int]string, answerCH chan answer, scoreCH chan score) {
	newScore := make(score)
	for newAnswer := range answerCH {
		if correctAnswers[int(newAnswer.numbQuestion)] == newAnswer.answer {
			newScore[newAnswer.userId] += 1
		}
	}
	scoreCH <- newScore
}

func Program() {
	correctAnswers := map[int]string{1: "2", 2: "3"}
	answerCH := make(chan answer)
	questionCH := make(chan question)
	scoreCH := make(chan score)
	wg := &sync.WaitGroup{}
	wg.Add(10)
	ctx, cancel := context.WithCancel(context.Background())
	questions := []string{"1+1", "1+2"}
	go roundGen(wg, questions, questionCH)
	go playerGen(ctx, wg, answerCH, questionCH)
	go counter(correctAnswers, answerCH, scoreCH)

	for newScore := range scoreCH {
		fmt.Println(newScore)
	}
	cancel()

}
