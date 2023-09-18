package textgame

import "fmt"

var castle = location{"Castle"}
var centralSquare = location{"Central Square"}
var forest = location{"Forest"}

func whatToDo() {
	fmt.Println("What to do?")
}

func whereToGo(c *character) {
	var location location
loop:
	fmt.Println("Where to go?\n1)Castle\n2)CentralSquare\n3)Forest")
	fmt.Scanln(&location)
	switch location {
	case castle:
		c.curentLocation = castle
	case centralSquare:
		c.curentLocation = centralSquare
	case forest:
		c.curentLocation = forest
	default:
		fmt.Println("Try again.")
		goto loop
	}
}

func choose(c *character) {
	fmt.Println("What you want:\n1)Do something\n2)Go somewhere")
	var choise uint
	fmt.Scan(&choise)
	if choise == 1 {
		whatToDo()
	}
	if choise == 2 {
		whereToGo(c)
	}
}

func Start_game() {
	newCharacter := character{hp: 3, curentLocation: centralSquare}
	fmt.Println("Named your character:")
	fmt.Scanln(&newCharacter.name)
	fmt.Printf("%s awake and see %s", newCharacter.name, newCharacter.curentLocation.name)
	choose(&newCharacter)
}
