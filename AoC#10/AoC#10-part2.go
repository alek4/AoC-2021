package main

import (
	"fmt"
	"strconv"
	"strings"
)

type octopus struct {
	val   int
	flash bool
}

func increment_adjacent(en_levels [][]octopus, count_flash *int, row, col int) {
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i < 0 || i >= len(en_levels) || j < 0 || j >= len(en_levels[0]) {
				continue
			} else {
				if en_levels[i][j] != en_levels[row][col] {
					en_levels[i][j].val++
					if en_levels[i][j].val > 9 && en_levels[i][j].flash == false {
						en_levels[i][j].val = 0
						en_levels[i][j].flash = true
						*count_flash++
						increment_adjacent(en_levels, count_flash, i, j)
					}

					// if en_levels[i][j].val == 9 {
					// 	en_levels[i][j].flash = true
					// 	en_levels[i][j].val = 0
					// 	*count_flash++
					// 	increment_adjacent(en_levels, count_flash, i, j)
					// } else if !en_levels[i][j].flash {
					// 	en_levels[i][j].val++
					// }
				}
			}
		}
	}

}

func increment(en_levels [][]octopus, count_flash *int) [][]octopus {
	for i, row := range en_levels {
		for j := range row {
			if en_levels[i][j].flash == false {
				en_levels[i][j].val++
			}
			if en_levels[i][j].val > 9 && en_levels[i][j].flash == false {
				en_levels[i][j].val = 0
				en_levels[i][j].flash = true
				*count_flash++
				increment_adjacent(en_levels, count_flash, i, j)
			}
			// if en_levels[i][j].val == 9 {
			// 	en_levels[i][j].val = 0
			// 	*count_flash++
			// 	increment_adjacent(en_levels, count_flash, i, j)
			// } else if !en_levels[i][j].flash {
			// 	en_levels[i][j].val++
			// }
		}
	}

	return en_levels
}

func printMatrix(matrix [][]octopus) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Printf("%d ", matrix[i][j].val)
		}
		fmt.Println()
	}
}

func refresh_oct(en_levels [][]octopus) {
	for i, row := range en_levels {
		for j := range row {
			en_levels[i][j].flash = false
		}
	}
}

func is_synchronized(en_levels [][]octopus) bool {
	for i, row := range en_levels {
		for j := range row {
			if en_levels[i][j].val != 0 {
				return false
			}
		}
	}

	return true
}

func main() {
	var row string
	en_levels := [][]octopus{}

	count_flash := 0
	for {
		_, err := fmt.Scanf("%s", &row)
		row_arr := []octopus{}

		if err != nil {
			break
		}

		for _, v := range strings.Split(row, "") {
			n, _ := strconv.Atoi(v)
			row_arr = append(row_arr, octopus{n, false})
		}

		en_levels = append(en_levels, row_arr)
	}

	var i int
	for {
		i++
		increment(en_levels, &count_flash)
		refresh_oct(en_levels)
		if is_synchronized(en_levels) {
			break
		}
		// fmt.Println("---")
		// printMatrix(en_levels)
	}

	fmt.Println(i)
}
