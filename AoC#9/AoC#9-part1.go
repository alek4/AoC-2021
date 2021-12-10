package main

import "fmt"

func calcPoints(char rune) int {
	if char == ')' {
		return 3
	}
	if char == ']' {
		return 57
	}
	if char == '}' {
		return 1197
	}
	if char == '>' {
		return 25137
	}

	return 0
}

func main() {
	var line string
	var closing_char_index int
	CLOSING_CHAR := [4]rune{')', ']', '}', '>'}
	OPENING_CHAR := [4]rune{'(', '[', '{', '<'}

	// res := 0
	for {
		_, err := fmt.Scanf("%s", &line)

		if err != nil {
			break
		}

		currChunks := []rune{}
		for _, v := range line {
			is_v_closing := false

			for i, c := range CLOSING_CHAR {
				if v == c {
					closing_char_index = i
					is_v_closing = true
					break
				}
			}

			if is_v_closing {
				if OPENING_CHAR[closing_char_index] == currChunks[len(currChunks)-1] {
					currChunks = currChunks[:len(currChunks)-1]
				} else {
					// fmt.Println(line)
					// res += calcPoints(CLOSING_CHAR[closing_char_index])
					break
				}
			} else {
				currChunks = append(currChunks, v)
			}
		}
		fmt.Println(line)
		
	}

	// fmt.Println(res)
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
