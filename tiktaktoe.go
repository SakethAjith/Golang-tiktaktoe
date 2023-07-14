package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type game struct {
	board   [][]string
	turn    bool
	winner  string
	lim     int
	choices map[int]struct{}
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
	num := 1
	tiktaktoe.choices = map[int]struct{}{}
	for i := 0; i < tiktaktoe.lim; i++ {
		res := []string{}
		for j := 0; j < tiktaktoe.lim; j++ {
			res = append(res, fmt.Sprintf("%d", num))
			tiktaktoe.choices[num] = struct{}{}
			num++
		}
		tiktaktoe.board = append(tiktaktoe.board, res)
	}
	return *tiktaktoe
}

func winner(checkStr []string, tiktaktoe *game) (bool, string) {
	x_string := strings.Repeat("X", tiktaktoe.lim)
	o_string := strings.Repeat("O", tiktaktoe.lim)
	board_string := strings.Join(checkStr, "")
	fmt.Println(board_string)
	if board_string == x_string {
		return true, "X"
	}
	if board_string == o_string {
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

	var RowCheckStr []string
	var ColCheckStr []string
	var DiagCheckStr []string
	var RevDiagCheckStr []string

	//Check Rows
	for i := 0; i < tiktaktoe.lim; i++ {
		RowCheckStr = board[i] //Check Rows

		if res, ans := winner(RowCheckStr, tiktaktoe); res {
			tiktaktoe.winner = ans
			return check(tiktaktoe)
		}

		DiagCheckStr = append(DiagCheckStr, board[i][i])
		RevDiagCheckStr = append(RevDiagCheckStr, board[i][tiktaktoe.lim-i-1])
	}

	//Check Diagonal
	if res, ans := winner(DiagCheckStr, tiktaktoe); res {
		tiktaktoe.winner = ans
		check(tiktaktoe)
	}

	if res, ans := winner(RevDiagCheckStr, tiktaktoe); res {
		tiktaktoe.winner = ans
		check(tiktaktoe)
	}

	//Check Columns
	for i := 0; i < tiktaktoe.lim; i++ {
		for j := 0; j < tiktaktoe.lim; j++ {
			ColCheckStr = append(ColCheckStr, board[j][i])
		}
		if res, ans := winner(ColCheckStr, tiktaktoe); res {
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
			if board[i][j] != "X" || board[i][j] != "O" {
				fmt.Println(board[i][j])
				holder = append(holder, i)
				holder = append(holder, j)
				res = append(res, holder)
			}
		}
	}
	return res
}

func mark(tiktaktoe *game, x, y int) {
	if tiktaktoe.turn {
		tiktaktoe.board[x][y] = "X"
		tiktaktoe.turn = false
		return
	}

	if !tiktaktoe.turn {
		tiktaktoe.board[x][y] = "O"
		tiktaktoe.turn = true
		return
	}
}

func findPos(tiktaktoe *game, x int) (i, j int) {
	i = x / tiktaktoe.lim
	if (x % tiktaktoe.lim) == 0 {
		i--
	}
	j = x - (i * tiktaktoe.lim)
	j--
	return
}

func randomMove(tiktaktoe *game) int {
	fmt.Println(tiktaktoe.choices)
	for i, _ := range tiktaktoe.choices {
		return i
	}
	return -1
}

func Move(tiktaktoe *game) int {
	n := len(tiktaktoe.choices)
	if n > 0 {
		var choice int
		if tiktaktoe.turn {
			choice = randomMove(tiktaktoe)
		}

		if !tiktaktoe.turn {
			choice = UserInput(tiktaktoe)
		}

		if choice > 0 {
			delete(tiktaktoe.choices, choice)
			res_i, res_j := findPos(tiktaktoe, choice)
			fmt.Println(res_i, res_j)
			mark(tiktaktoe, res_i, res_j)
		}
		return n
	}

	return 0
}

func UserInput(tiktaktoe *game) int {
	fmt.Println(tiktaktoe.board)
	fmt.Println("Enter move position:")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	user_pos, err := strconv.Atoi(input.Text())
	if err != nil {
		fmt.Println("Invalid Input please choose an available spot!")
		return -1
	}

	if _, ok := tiktaktoe.choices[user_pos]; !ok {
		fmt.Println("That Spot is taken! please choose an available spot!")
		return -1
	}

	return user_pos

}

func main() {
	var tiktaktoe game
	tiktaktoe.lim = 5
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
	// fmt.Println(findPos(&tiktaktoe, 5))
}
