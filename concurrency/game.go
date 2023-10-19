package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

type Answer struct {
	Name         string
	QuestionText string
	AnswerText   string
}

type Question struct {
	Answer chan Answer
	Text   string
}

type score map[string]int

type Player struct {
	Name  string
	Inbox chan Question
}

type Lobby []Player

type PlayerCH chan Player

const ROUND_COUNT = 2
const ROUND_QUESTION_COUNT = 2
const MAX_LOBBY_PLAYERS = 3

var questions = []string{
	"Is 2 + 2 = 4?",
	"Is the sky blue?",
}

var answersToQuestions = []string{
	"Yes",
	"Yes",
}

// recive player, question and chan for answering questions
func AskRoundQuestion(player Player, answers chan Answer, question string) {
	player.Inbox <- Question{answers, question}

}

// cycle wich recive question from player.inbox (chanal) and give answer
func AnswerRoundQuestion(player Player) {
	for i := 0; i < ROUND_QUESTION_COUNT; i++ {
		question := <-player.Inbox
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)+300))
		question.Answer <- Answer{player.Name, question.Text, "Yes"}
	}
}

// create and return lobby, recive players from Player chan
func AddPlayersToLobby(playerCH PlayerCH) Lobby {
	lobby := make(Lobby, 0)

	for j := 0; j < MAX_LOBBY_PLAYERS; j++ {
		lobby = append(lobby, <-playerCH)
	}

	return lobby
}

func SendQuestionsToPlayers(lobby Lobby) chan Answer {

	answers := make(chan Answer)
	// range of all questions
	for _, question := range questions {
		// each player ask question from questions

		for _, player := range lobby {
			go AskRoundQuestion(player, answers, question)

		}
		time.Sleep(10 * time.Second)
	}

	return answers
}

func PrintScore(newScore score) {
	for key, value := range newScore {
		fmt.Printf("Player: %s, Points: %d\n", key, value)
	}

}

func GatherAnswers(lobby Lobby, answers chan Answer, newScore score) {

	for i := range questions {
		for range lobby {
			playerAnswer := <-answers

			if playerAnswer.AnswerText == answersToQuestions[i] {
				newScore[playerAnswer.Name] += 1
			} else {
				newScore[playerAnswer.Name] += 0
			}

		}
		PrintScore(newScore)
	}
}

func StartRounds(quit chan bool) PlayerCH {
	rounds := make(PlayerCH)

	go func() {
		defer close(rounds)
		defer close(quit)

		newScore := make(score)
		for i := 0; i < ROUND_COUNT; i++ {
			lobby := AddPlayersToLobby(rounds)

			roundAnswers := SendQuestionsToPlayers(lobby)

			GatherAnswers(lobby, roundAnswers, newScore)

			fmt.Println(i+1, "round ends!")
		}

	}()

	return rounds
}

func Program() {
	quit := make(chan bool)
	rounds := StartRounds(quit)

	players := Lobby{
		{"Greg", make(chan Question)},
		{"Bob", make(chan Question)},
		{"Jack", make(chan Question)},
	}

	for round := 0; round < ROUND_COUNT; round++ {
		for _, player := range players {
			rounds <- player
			go AnswerRoundQuestion(player)
		}
	}

	<-quit
}
