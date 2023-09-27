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

func turn(b board, mark uint8) {
	var input string
enter:
	fmt.Scan(&input)

	intInput, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Error:%v\nTry again.", err)
		goto enter
	}
	if b[intInput] == 0 {
		b[intInput] = cell(mark)
	} else if b[intInput] == 1 || b[intInput] == 2 {
		fmt.Println("Cell is bussy. Try again")
		goto enter
	} else {
		fmt.Println("Try again!")
		goto enter
	}
}

func checkRaw(b board) (player cell) {
	for i := 1; i <= 7; i += 3 {
		if b[i] != b[i+1] {
			continue
		}
		if b[i] != b[i+2] {
			continue
		}
		return b[i]
	}
	return 0
}

func checkColumn(b board) (player cell) {
	for i := 1; i <= 3; i += 1 {
		if b[i] != b[i+3] {
			continue
		}
		if b[i] != b[i+6] {
			continue
		}
		return b[i]
	}
	return 0
}

func checkDigonal(b board) (player cell) {
	for i := 1; i <= 3; i += 2 {
		if b[i] != b[i+4] {
			continue
		}
		if b[i] != b[i+8] {
			continue
		}
		return b[i]
	}
	return 0
}

func whoWin(b board) (player cell) {
	switch {
	case checkColumn(b) != 0:
		return checkColumn(b)
	case checkRaw(b) != 0:
		return checkRaw(b)
	case checkDigonal(b) != 0:
		return checkDigonal(b)
	default:
		return 0
	}

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
	if isLastTurn(b) && whoWin(b) == 0 {
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
}

func Xsosgame() {
	b := Makeboard()
	b = clearboard(b)
	var player uint
	for player == 0 {
		drawBoard(b)
		turn(b, xsmarc)
		drawBoard(b)
		if whoWin(b) == cell(xsmarc) {
			fmt.Println("Xs win!")
			break
		}
		if checkTie(b) {
			fmt.Println("Tie")
			break
		}
		turn(b, osmarc)
		drawBoard(b)
		if whoWin(b) == cell(osmarc) {
			fmt.Println("Os win!")
			break
		}
		if checkTie(b) {
			fmt.Println("Tie")
			break
		}
	}
}
