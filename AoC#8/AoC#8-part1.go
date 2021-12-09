package main

import (
	"fmt"
	"strconv"
	"strings"
)

func checkAdjacent(grid [][]int) int {
	res := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if j == 0 {
				if i == 0 {
					if grid[i][j] < grid[i][j + 1] && grid[i][j] < grid[i + 1][j] {
						res += grid[i][j] + 1
					}
				} else if i == len(grid) - 1 {
					if grid[i][j] < grid[i][j + 1] && grid[i][j] < grid[i - 1][j] {
						res += grid[i][j] + 1
					}
				} else {
					if grid[i][j] < grid[i][j + 1] && grid[i][j] < grid[i - 1][j] && grid[i][j] < grid[i + 1][j] {
						res += grid[i][j] + 1
					}
				}
			} else if i == 0 {
				if j == len(grid[i]) - 1 {
					if grid[i][j] < grid[i][j - 1] && grid[i][j] < grid[i + 1][j] {
						res += grid[i][j] + 1
					}
				} else {
					if grid[i][j] < grid[i][j + 1] && grid[i][j] < grid[i][j - 1] && grid[i][j] < grid[i + 1][j] {
						res += grid[i][j] + 1
					}
				}
			} else if j == len(grid[i]) - 1 {
				if i == len(grid) - 1 {
					if grid[i][j] < grid[i][j - 1] && grid[i][j] < grid[i - 1][j] {
						res += grid[i][j] + 1
					}
				} else {
					if grid[i][j] < grid[i][j - 1] && grid[i][j] < grid[i + 1][j] && grid[i][j] < grid[i - 1][j] {
						res += grid[i][j] + 1
					}
				}
			} else if i == len(grid) - 1 {
				if grid[i][j] < grid[i][j - 1] && grid[i][j] < grid[i][j + 1] && grid[i][j] < grid[i - 1][j] {
					res += grid[i][j] + 1
				}
			} else {
				if grid[i][j] < grid[i][j - 1] && grid[i][j] < grid[i][j + 1] && grid[i][j] < grid[i - 1][j] && grid[i][j] < grid[i + 1][j] {
					res += grid[i][j] + 1
				}
			}
		}
	}

	return res
}

func main() {
	var row string
	var grid [][]int

	grid = make([][]int, 0)

	for {
		_, err := fmt.Scanf("%s", &row)

		if err != nil {
			break
		}

		row_split := strings.Split(row, "")

		row_slice := []int{}
		for _, v := range row_split {
			v_int, _ := strconv.Atoi(v)
			row_slice = append(row_slice, v_int)
		}

		grid = append(grid, row_slice)
	}

	fmt.Println(checkAdjacent(grid))
}
