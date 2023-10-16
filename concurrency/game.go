package concurrency

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
)

type question string

type answer struct {
	userId       string
	numbQuestion uint
	answer       string
}

type score map[string]int

func roundGen(questions []string) chan question {
	questionCH := make(chan question)
	go func() {
		for i := 0; i < len(questions); i++ {
			for j := 0; j < 10; j++ {
				questionCH <- question(questions[i])
			}
			// <- block1(10 times)
			// <- unblock2

		}
	}()

	return questionCH
}

func player(ctx context.Context, name string, answerCH chan answer, questionCH chan question) {
	for {
		select {
		case <-questionCH:
			ansString := strconv.Itoa(rand.Intn(9))
			playerAns := answer{
				userId: name,
				answer: ansString}
			answerCH <- playerAns
			// <- unblock1
			// <- block2

		case <-ctx.Done():
			return
		}
	}
}

func playerGen(ctx context.Context, questionCH chan question) chan answer {
	answerCH := make(chan answer)
	for i := 1; i <= 10; i++ {

		go player(ctx, fmt.Sprintf("user%d", i), answerCH, questionCH)
	}
	return answerCH
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
	questions := []string{"1+1", "1+2"}
	correctAnswers := map[int]string{1: "2", 2: "3"}
	ctx, cancel := context.WithCancel(context.Background())

	questionCH := roundGen(questions)
	answerCH := playerGen(ctx, questionCH)
	scoreCH := make(chan score)

	go counter(correctAnswers, answerCH, scoreCH)

	for newScore := range scoreCH {
		fmt.Println(newScore)
	}
	cancel()

}
