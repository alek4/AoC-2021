package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	// "strings"
)

func checkWin(board [5][5]int) bool {
	count := 0
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if board[row][col] == -1 {
				count++

				if count == 5 {
					return true
				}
			}
		}

		count = 0
	}

	count = 0
	for col := 0; col < 5; col++ {
		for row := 0; row < 5; row++ {
			if board[row][col] == -1 {
				count++

				if count == 5 {
					return true
				}
			}
		}

		count = 0
	}

	return false
}

func checkBoard(board [5][5]int, p_inputSlice []string) ([5][5]int, int) {
	for i := 0; i < len(p_inputSlice); i++ {
		for row := 0; row < 5; row++ {
			for col := 0; col < 5; col++ {
				n, _ := strconv.Atoi(p_inputSlice[i])
				// fmt.Println(board[row][col], "==", n)
				if board[row][col] == n {
					board[row][col] = -1
					if checkWin(board) {
						return board, i
					}
				}
			}
		}
	}

	return board, -1
}

func calcRes(board [5][5]int, last_num string) int {
	sum := 0

	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if board[row][col] != -1 {
				sum += board[row][col]
			}
		}
	}

	num, _ := strconv.Atoi(last_num)

	return sum * num
}

func main() {
	var num int
	var p_input string
	var res int
	board := [5][5]int{}

	fmt.Scanf("%s", &p_input)
	p_inputSlice := strings.Split(p_input, ",")

	row := 0
	col := 0

	max := 0

loop:
	for {
		for {
			_, err := fmt.Scanf("%d", &num)
	
			if err == io.EOF {
				break loop
			}

			if err != nil {
				break
			}

	
			board[row][col] = num
	
			col += 1
	
			if col > 4 {
				col = 0
				row += 1
			}
		}

		row = 0
		col = 0

		board, curr := checkBoard(board, p_inputSlice)

		if max == 0 {
			max = curr
		}

		if curr != -1 && max < curr {
			max = curr
			res = calcRes(board, p_inputSlice[max])
		}
	}
	
	println(res)

}
