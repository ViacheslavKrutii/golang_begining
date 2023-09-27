package mapgame

import (
	"fmt"
	"strconv"
)

type cell uint8

const osmarc = uint8(1)
const xsmarc = uint8(2)

type board map[int]cell

func Makeboard() board {
	gameboard := make(board)
	for i := 1; i <= 9; i++ {
		gameboard[i] = cell(0)
	}
	return gameboard
}

func clearboard(b board) board {
	for i := range b {
		b[i] = 0
	}
	return b
}

func xsturn(b board) {
	var input string
enter:

	fmt.Scan(&input)

	intInput, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Error:%v\nTry again.\n", err)
		goto enter
	}
	if b[intInput] == 0 {
		b[intInput] = cell(xsmarc)
	} else if b[intInput] == 1 || b[intInput] == 2 {
		fmt.Println("Cell is bussy. Try again")
		goto enter
	} else {
		fmt.Println("Try again!")
		goto enter
	}
}

func osturn(b board) {
	var input string
enter:
	fmt.Scan(&input)

	intInput, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Error:%v\nTry again.", err)
		goto enter
	}
	if b[intInput] == 0 {
		b[intInput] = cell(osmarc)
	} else if b[intInput] == 1 || b[intInput] == 2 {
		fmt.Println("Cell is bussy. Try again")
		goto enter
	} else {
		fmt.Println("Try again!")
		goto enter
	}
}

// does xs win?
func checkXsWin(b board) bool {

	//first raw
	if b[1] == 2 && b[2] == 2 && b[3] == 2 {

		return true
	}

	//second raw
	if b[4] == 2 && b[5] == 2 && b[6] == 2 {

		return true
	}

	//third raw
	if b[7] == 2 && b[8] == 2 && b[9] == 2 {

		return true
	}

	//first column
	if b[1] == 2 && b[4] == 2 && b[7] == 2 {

		return true
	}

	//second column
	if b[2] == 2 && b[5] == 2 && b[8] == 2 {

		return true
	}

	//third column
	if b[3] == 2 && b[6] == 2 && b[9] == 2 {

		return true
	}

	//first diagonal
	if b[1] == 2 && b[5] == 2 && b[9] == 2 {

		return true
	}

	//second diagonal
	if b[3] == 2 && b[5] == 2 && b[7] == 2 {

		return true
	}

	return false
}

// does os win?
func checkOsWin(b board) bool {

	//first raw
	if b[1] == 1 && b[2] == 1 && b[3] == 1 {

		return true
	}

	//second raw
	if b[4] == 1 && b[5] == 1 && b[6] == 1 {

		return true
	}

	//third raw
	if b[7] == 1 && b[8] == 1 && b[9] == 1 {

		return true
	}

	//first column
	if b[1] == 1 && b[4] == 1 && b[7] == 1 {

		return true
	}

	//second column
	if b[2] == 1 && b[5] == 1 && b[8] == 1 {

		return true
	}

	//third column
	if b[3] == 1 && b[6] == 1 && b[9] == 1 {

		return true
	}

	//first diagonal
	if b[1] == 1 && b[5] == 1 && b[9] == 1 {

		return true
	}

	//second diagonal
	if b[3] == 1 && b[5] == 1 && b[7] == 1 {

		return true
	}

	return false
}

func isLastTurn(b board) bool {
	empty := 0
	for i := range b {
		if b[i] == 0 {
			empty++
		}
	}
	return empty == 0
}

// Tie?
func checkTie(b board) bool {
	if isLastTurn(b) && !checkOsWin(b) && !checkXsWin(b) {
		return true
	}
	return false
}

func drawBoard(b board) {
	// formating board
	fb := make(map[int]string)
	for i := range b {
		switch b[i] {
		case 0:
			fb[i] = " "
		case 1:
			fb[i] = "o"

		case 2:
			fb[i] = "x"

		}
	}

	fmt.Printf(" %v | %v | %v\n", fb[1], fb[2], fb[3])
	fmt.Println("-----------")
	fmt.Printf(" %v | %v | %v\n", fb[4], fb[5], fb[6])
	fmt.Println("-----------")
	fmt.Printf(" %v | %v | %v\n", fb[7], fb[8], fb[9])
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func Xsosgame() {
	b := Makeboard()
	b = clearboard(b)
	for s := true; s; {
		drawBoard(b)
		xsturn(b)
		drawBoard(b)
		if checkXsWin(b) {
			fmt.Println("Xs win!")
			break
		}
		if checkTie(b) {
			fmt.Println("Tie")
			break
		}
		osturn(b)
		drawBoard(b)
		if checkOsWin(b) {
			fmt.Println("Os win!")
			break
		}
		if checkTie(b) {
			fmt.Println("Tie")
			break
		}
	}
}
