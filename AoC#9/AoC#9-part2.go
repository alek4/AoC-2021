package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func calcPoints(char rune) int {
	if char == ')' {
		return 1
	}
	if char == ']' {
		return 2
	}
	if char == '}' {
		return 3
	}
	if char == '>' {
		return 4
	}

	return 0
}

func main() {
	var line string
	var opening_char_index int
	CLOSING_CHAR := [4]rune{')', ']', '}', '>'}
	OPENING_CHAR := [4]rune{'(', '[', '{', '<'}

	corruptedChunks := []string{}

	for {
		_, err := fmt.Scanf("%s", &line)

		if err != nil {
			break
		}

		corruptedChunks = append(corruptedChunks, line)
		// fmt.Println(line)
	}

	

	file, err := os.Open("input")
	if err != nil {
		fmt.Println("file not found")
		return
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	res := 0
	res_slice := []int{}
	for {
		fileScanner.Scan()
		lines := fileScanner.Text()

		if lines == "" {
			break
		}

		is_ok := true
		for _, c := range corruptedChunks {
			if lines == c {
				// fmt.Println(string(lines[i]), c)
				is_ok = false
				break
			}
		}
		
		if is_ok {
			currChunks := []rune{}
			res_string := ""
			for i := len(lines) - 1; i >= 0; i-- {
				is_v_opening := false

				for j, c := range OPENING_CHAR {
					if rune(lines[i]) == c {
						// fmt.Println(string(lines[i]), c)
						opening_char_index = j
						is_v_opening = true
						break
					}
				}

				if is_v_opening {
					if len(currChunks) != 0 {
						if CLOSING_CHAR[opening_char_index] == currChunks[len(currChunks)-1] {
							currChunks = currChunks[:len(currChunks)-1]
						}
					} else {
						res_string += string(CLOSING_CHAR[opening_char_index])
						res = res * 5 + calcPoints(CLOSING_CHAR[opening_char_index])
					}

				} else {
					currChunks = append(currChunks, rune(lines[i]))
				}

			}
			// fmt.Println(res)
			res_slice = append(res_slice, res)
			res = 0
		}
	}

	sort.Ints(res_slice)
	fmt.Println(res_slice[len(res_slice)/2])
}

// [
// 	[
// 		<
// 			[
// 				(
// 					[

// 					]
// 				)
// 			)<([[{}[[()]]]
