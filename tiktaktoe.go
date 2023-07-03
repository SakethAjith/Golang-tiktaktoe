package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type game struct {
	board  [][]string
	turn   bool
	winner string
	lim    int
}

func PrintBoard(tiktaktoe *game) {

	for indx, i := range tiktaktoe.board {
		row := ""
		for _, j := range i {
			if len(j) > 0 {
				row += fmt.Sprintf("|%s|", j)
			} else {
				row += fmt.Sprintf("| |")
			}
		}
		fmt.Println(row)
		if indx < tiktaktoe.lim-1 {
			fmt.Println(strings.Repeat("=", len(row)))
		}
	}

}

func makeBoard(tiktaktoe *game) game {
	for i := 0; i < tiktaktoe.lim; i++ {
		res := []string{}
		for j := 0; j < tiktaktoe.lim; j++ {
			res = append(res, "")
		}
		tiktaktoe.board = append(tiktaktoe.board, res)
	}
	return *tiktaktoe
}

func winner(checkStr []string, tiktaktoe *game) (bool, string) {
	x_string := strings.Repeat("X", tiktaktoe.lim)
	o_string := strings.Repeat("O", tiktaktoe.lim)
	if strings.Join(checkStr, "") == x_string {
		return true, "X"
	}
	if strings.Join(checkStr, "") == o_string {
		return true, "O"
	}
	return false, ""
}

func check(tiktaktoe *game) bool {
	if len(tiktaktoe.winner) > 0 {
		fmt.Printf("Winner is %v!\n", tiktaktoe.winner)
		return true
	}

	board := tiktaktoe.board

	var checkStr []string

	//Check Rows
	for i := 0; i < tiktaktoe.lim; i++ {
		checkStr = board[i]
		if res, ans := winner(checkStr, tiktaktoe); res {
			tiktaktoe.winner = ans
			return check(tiktaktoe)
		}
	}

	//Check Columns
	for i := 0; i < tiktaktoe.lim; i++ {
		checkStr = []string{}
		for j := 0; j < tiktaktoe.lim; j++ {
			checkStr = append(checkStr, board[i][j])
		}
		if res, ans := winner(checkStr, tiktaktoe); res {
			tiktaktoe.winner = ans
			check(tiktaktoe)
		}
	}
	checkStr = []string{}

	//Check Diagonal
	for i := 0; i < tiktaktoe.lim; i++ {
		checkStr = append(checkStr, board[i][i])
		if res, ans := winner(checkStr, tiktaktoe); res {
			tiktaktoe.winner = ans
			check(tiktaktoe)
		}

	}

	checkStr = []string{}
	//Check Diagonal
	for i := 0; i < tiktaktoe.lim; i++ {
		checkStr = append(checkStr, board[i][tiktaktoe.lim-i-1])
		if res, ans := winner(checkStr, tiktaktoe); res {
			tiktaktoe.winner = ans
			check(tiktaktoe)
		}

	}

	return false
}

func freeSpaces(tiktaktoe *game) [][]int {
	board := tiktaktoe.board
	res := make([][]int, 0, tiktaktoe.lim*tiktaktoe.lim)

	for i := 0; i < tiktaktoe.lim; i++ {
		for j := 0; j < tiktaktoe.lim; j++ {
			holder := []int{}
			if len(board[i][j]) == 0 {
				holder = append(holder, i)
				holder = append(holder, j)
				res = append(res, holder)
			}
		}
	}
	return res
}

func Move(tiktaktoe *game) int {
	choices := freeSpaces(tiktaktoe)
	numberOfChoices := len(choices)
	if numberOfChoices > 0 {
		res := choices[rand.Intn(len(choices))]

		if tiktaktoe.turn == true {
			tiktaktoe.board[res[0]][res[1]] = "X"
			tiktaktoe.turn = false
			return len(choices)
		}

		if tiktaktoe.turn == false {
			tiktaktoe.board[res[0]][res[1]] = "O"
			tiktaktoe.turn = true
			return len(choices)
		}
		return len(choices)
	}
	return 0
}

func main() {
	var tiktaktoe game
	tiktaktoe.lim = 3
	tiktaktoe = makeBoard(&tiktaktoe)

	for {
		moves_left := Move(&tiktaktoe)
		PrintBoard(&tiktaktoe)
		fmt.Println()
		res := check(&tiktaktoe)
		if len(tiktaktoe.winner) > 0 {
			break
		}

		if (res == false) && (moves_left == 0) {
			fmt.Println("Draw!")
			break
		}
	}

	fmt.Println(tiktaktoe.board)
}
