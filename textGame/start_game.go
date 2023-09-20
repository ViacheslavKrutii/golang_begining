package textgame

import "fmt"

//Castle actions
var talkToKing = action{"Try to talk with King", []string{"Ask for lands", "Tell about myself"}}
var talkGuardians = action{"Try to talk with Guardiavns", []string{"Information about King", "Ask for job"}}

//Central Square actions
var cadge = action{"Try to cadge some money", []string{"Shout to the whole square", "Sit with outstretched arm"}}
var arrangeDuel = action{"Try arrange a duel", []string{"Arrange a duel with knight", "Arrange a duel with chicken"}}

//Forest actions
var chopOaks = action{"Try to chop oaks", []string{"Chop huge oak", "Chop small oak"}}
var hunt = action{"Try to hunt", []string{"Hunt a hare", "Hunt a deer"}}

var castle = location{"Castle", []action{talkToKing, talkGuardians}}
var centralSquare = location{"Central Square", []action{cadge, arrangeDuel}}
var forest = location{"Forest", []action{chopOaks, hunt}}

//character actions
func whatToDo(c *character) {
	var input1 string
	var input2 string
	curentEvents := make([]string, 0)
	for _, v := range c.curentLocation.actions {
		curentEvents = append(curentEvents, v.discroption)
	}
loop:
	fmt.Printf("Your location is %s\nWhat do you want?\n", c.curentLocation.name)
	for i, item := range curentEvents {
		fmt.Printf("%d) %s\n", i+1, item)
	}
	fmt.Scan(&input1)
	switch input1 {
	case "1":
		curentEvents = c.curentLocation.actions[0].choise
		fmt.Println("What do you want?")
		for i, item := range curentEvents {
			fmt.Printf("%d) %s\n", i+1, item)
		}
		fmt.Scan(&input2)

	case "2":
		curentEvents = c.curentLocation.actions[1].choise
		fmt.Println("What do you want?")
		for i, item := range curentEvents {
			fmt.Printf("%d) %s\n", i+1, item)
		}
		fmt.Scan(&input2)

	default:
		fmt.Println("Try again!")
		goto loop
	}

	if c.curentLocation.name == "Castle" && input1 == "1" && input2 == "1" {
		c.takeDamage()
		fmt.Println("King angry.\nKing hit you.")
	} else if c.curentLocation.name == "Castle" && input1 == "1" && input2 == "2" {
		fmt.Println("King The king is delighted with the story of your life.")
		fmt.Println("The king transfers his powers to you!")
		c.becameTheKing()
	} else if c.curentLocation.name == "Castle" && input1 == "2" && input2 == "1" {
		fmt.Println("*Story about king*")
	} else if c.curentLocation.name == "Castle" && input1 == "2" && input2 == "2" {
		fmt.Println("We don't have a job for you.")
	} else if c.curentLocation.name == "Central Square" && input1 == "1" && input2 == "1" {
		fmt.Println("The merchant throws an egg at you.")
		c.takeDamage()
	} else if c.curentLocation.name == "Central Square" && input1 == "1" && input2 == "2" {
		fmt.Println("Passer-by gives you bread.")
		c.eat()
	} else if c.curentLocation.name == "Central Square" && input1 == "2" && input2 == "1" {
		fmt.Println("Knight accepts your challenge.")
		c.takeDamage()
	} else if c.curentLocation.name == "Central Square" && input1 == "2" && input2 == "2" {
		fmt.Println("The owner of the chicken hits you with a broom.")
		c.takeDamage()
	} else if c.curentLocation.name == "Forest" && input1 == "1" && input2 == "1" {
		fmt.Println("You became the king ...\nBut the king of lumberjacks.")
	} else if c.curentLocation.name == "Forest" && input1 == "1" && input2 == "2" {
		fmt.Println("You chop small oak.")
	} else if c.curentLocation.name == "Forest" && input1 == "2" && input2 == "1" {
		fmt.Println("You hunted a hare.")
		c.eat()
	} else if c.curentLocation.name == "Forest" && input1 == "2" && input2 == "2" {
		fmt.Println("The deer injured you.")
		c.takeDamage()
	} else {
		fmt.Println("Try again!")
		goto loop
	}
}

func whereToGo(c *character) {
	var input string
	fmt.Printf("Your location is %s\n", c.curentLocation.name)
loop:
	fmt.Println("Where to go?\n1)Castle\n2)Forest\n3)CentralSquare")
	fmt.Scan(&input)
	switch input {
	case "1":
		c.curentLocation = castle
		fmt.Printf("You arive to %s\n", c.curentLocation.name)
	case "2":
		c.curentLocation = forest
		fmt.Printf("You arive to %s\n", c.curentLocation.name)
	case "3":
		c.curentLocation = centralSquare
		fmt.Printf("You arive to %s\n", c.curentLocation.name)
	default:
		fmt.Println("Try again")
		goto loop
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
