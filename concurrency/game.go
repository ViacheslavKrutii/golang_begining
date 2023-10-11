package concurrency

import "fmt"

type answer struct {
	userId string
	answer string
}

func roundGen() {

}

func player(name string, ch1 chan answer) {
	ansString := ""
	fmt.Scan(&ansString)
	playerAns := answer{
		userId: name,
		answer: ansString}
	ch1 <- playerAns
}

func playerGen() {
	ch1 := make(chan answer)
	for i := 1; i <= 5; i++ {
		go player(fmt.Sprintf("user%d", i), ch1)
	}

}

func counter() {

}
