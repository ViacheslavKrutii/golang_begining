package textgame

import "fmt"

//Castle actions
var talkToKing = event{
	discroption: "Try to talk with King",
	action:      "",
	nextEvent:   &AskForLands}

var AskForLands = event{
	discroption: "Ask for lands",
	action:      takeDamage,
	nextEvent:   nil,
}
var tellAboutMyself = event{
	discroption: "Tell about myself",
	action:      becameTheKing,
	nextEvent:   nil,
}

var talkGuardians = event{
	discroption: "Try to talk with Guardiavns",
	action:      takeDamage,
	nextEvent:   nil}

//Central Square actions
// var cadge = event{"Try to cadge some money", []string{"Shout to the whole square", "Sit with outstretched arm"}}
// var arrangeDuel = event{"Try arrange a duel", []string{"Arrange a duel with knight", "Arrange a duel with chicken"}}

//Forest actions
// var chopOaks = event{"Try to chop oaks", []string{"Chop huge oak", "Chop small oak"}}
// var hunt = event{"Try to hunt", []string{"Hunt a hare", "Hunt a deer"}}

var castle = location{"Castle", []event{talkToKing, talkGuardians}}

// var centralSquare = location{"Central Square", []event{cadge, arrangeDuel}}
// var forest = location{"Forest", []event{chopOaks, hunt}}

//character actions
func whatToDo(c *character) {
	for i, item := range c.curentLocation.events {
		fmt.Printf("%d. %s\n", i+1, item.discroption)
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
	// case "2":
	// 	c.curentLocation = forest
	// 	fmt.Printf("You arive to %s\n", c.curentLocation.name)
	// case "3":
	// 	c.curentLocation = centralSquare
	// 	fmt.Printf("You arive to %s\n", c.curentLocation.name)
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
	newCharacter := character{hp: 3, curentLocation: castle, areYouKing: false}
	fmt.Println("Named your character:")
	fmt.Scanln(&newCharacter.name)
	fmt.Printf("%s awake and see %s", newCharacter.name, newCharacter.curentLocation.name)
	for newCharacter.hp != 0 {
		choose(&newCharacter)
	}
}
