package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type fold_type struct {
	axis string
	val  int
}

func fold(grid [][]rune, fold_line fold_type) [][]rune {
	if fold_line.axis == "y" {
		top := grid[:fold_line.val]
		bottom := grid[fold_line.val+1:]

		for i, row := range bottom {
			for j, v := range row {
				if bottom[i][j] == '#' {
					top[len(top)-1-i][j] = v
				}
			}
		}

		return top
	} else {
    left := [][]rune{}
    right := [][]rune{}
		for i, row := range grid {
			left = append(left, row[:fold_line.val])
			right = append(right, row[fold_line.val + 1:])

      // fmt.Println(left)
      // fmt.Println(right)

      if len(left[i]) >= len(right[i]) {
        for j, v := range right[i] {
          if right[i][j] == '#' {
            // fmt.Println(len(left[i]),len(right[i]),i,j)
            left[i][len(left[i])-j-1] = v
          }
        }
      }
    }
    return left
  }
}

func count_dots(grid [][]rune) (counter int) {
	for i, row := range grid {
		for j := range row {
			if grid[i][j] == '#' {
				counter++
			}
		}
	}

	return
}

func main() {
	var row string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	input := [][]int{}
	folds := []fold_type{}
	max_x := -1
	max_y := -1
	for scanner.Scan() {
		row = scanner.Text()

		if !strings.Contains(row, "=") {
			row_slice := strings.Split(row, ",")

			x, _ := strconv.Atoi(row_slice[0])
			if max_x < x {
				max_x = x
			}

			y, _ := strconv.Atoi(row_slice[1])
			if max_y < y {
				max_y = y
			}

			input = append(input, []int{x, y})
		} else {
			pos_equal := strings.Index(row, "=")
			temp_v, _ := strconv.Atoi(string(row[pos_equal+1:]))
			folds = append(folds, fold_type{string(row[pos_equal-1]), temp_v})
		}
	}

	grid := make([][]rune, max_y+1)

	for i := range grid {
		grid[i] = make([]rune, max_x+1)
	}

	for _, v := range input {
		grid[v[1]][v[0]] = '#'
	}

	// grid = fold(grid, folds[0])
	// grid = fold(grid, folds[1])

  for _, f := range folds {
    fmt.Println(f)
	  grid = fold(grid, f)
    fmt.Println("---")
  }

	for _, row := range grid {
		for _, v := range row {
	    if v != '#' {
	      fmt.Print(".")
	    } else {
	      fmt.Print(string(v))
	    }
	  }
	  fmt.Println()
	}

	// fmt.Println(count_dots(grid))
}
