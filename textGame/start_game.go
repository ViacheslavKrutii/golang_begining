package textgame

import "fmt"

//inititalisation locations
var castle = location{"Castle"}
var centralSquare = location{"Central Square"}
var forest = location{"Forest"}

//character actions
func whatToDo(c *character) {
	var action1 string
	var action2 string
	var answer uint

	switch c.curentLocation {
	case castle:
		action1 = "Kill The King"
		action2 = "Ask for a job"
	case forest:
		action1 = "Chop oaks"
		action2 = "Hunt"
	//Central Square
	default:
		action1 = "Cadge"
		action2 = "Rob"

	}
loop:
	fmt.Printf("What to do?\n1)%s\n2)%s\n", action1, action2)

	fmt.Scanln(&answer)

	if answer == 1 {
		switch action1 {
		case "Kill The King":
			fmt.Println("Guardians smash you. You are died.")
			c.hp = 0
		case "Chop oaks":
			fmt.Println("You are The King ... But King of lumberjacks")
		//"Cadge"
		default:
			c.takeDamage()

		}
	} else if answer == 2 {
		switch action2 {
		case "Ask for a job":
			fmt.Println("You are new king")
			c.becameTheKing()
		case "Hunt":
			fmt.Println("You are hunt the rabbit but he bite you")
			c.takeDamage()
			c.eat()

		//"Rob"
		default:
			c.takeDamage()

		}
	} else {
		fmt.Println("Try again.1")
		goto loop
	}

}

func whereToGo(c *character) {
	var location uint
	var answer string
	back := c.curentLocation
loop:
	fmt.Println("Where to go?\n1)Castle\n2)CentralSquare\n3)Forest")
	fmt.Scanln(&location)
	switch location {
	case 1:
		c.curentLocation = castle
	case 2:
		c.curentLocation = centralSquare
	case 3:
		c.curentLocation = forest
	default:
		fmt.Println("Try again.")
		goto loop
	}
	fmt.Println("Are you shure?")
	fmt.Scanln(&answer)
	if answer == "no" {
		c.curentLocation = back
	}
}

//choose to do or to go
func choose(c *character) {
	var choise uint
loop:
	fmt.Println("What you want:\n1)Do something\n2)Go somewhere")
	fmt.Scan(&choise)
	if choise == 1 {
		whatToDo(c)
	} else if choise == 2 {
		whereToGo(c)
	} else {
		fmt.Println("Try again.")
		goto loop
	}
}

func Start_game() {
	newCharacter := character{hp: 3, curentLocation: centralSquare, areYouKing: false}
	fmt.Println("Named your character:")
	fmt.Scanln(&newCharacter.name)
	fmt.Printf("%s awake and see %s", newCharacter.name, newCharacter.curentLocation.name)
	for newCharacter.hp != 0 {
		choose(&newCharacter)
	}
}
