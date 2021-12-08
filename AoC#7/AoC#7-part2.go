package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func findUniqueDigits(values []string) (uniqueDigits map[int]string) {
	uniqueDigits = make(map[int]string)
	for _, val := range values {
		if len(val) == 2 {
			uniqueDigits[1] = val
		}

		if len(val) == 4 {
			uniqueDigits[4] = val
		}

		if len(val) == 3 {
			uniqueDigits[7] = val
		}

		if len(val) == 7 {
			uniqueDigits[8] = val
		}
	}

	return

}

func decodeDigits(keys []int, uniqueDigits map[int]string, n_tries int) (numbers [][]string) {
	var display [7]string
	for _, k := range keys {
		switch k {
		case 1:
			// fmt.Println(k)
			// fmt.Println(uniqueDigits[k])
			if n_tries == 4 || n_tries == 5 || n_tries == 6 || n_tries == 7 {
				display[3] = string(uniqueDigits[k][0])
				display[1] = string(uniqueDigits[k][1])
			} else {
				display[1] = string(uniqueDigits[k][0])
				display[3] = string(uniqueDigits[k][1])
			}
			break
		case 4:
			// fmt.Println(uniqueDigits[k])

			uniqueDigits[k] = strings.ReplaceAll(uniqueDigits[k], display[1], "")
			uniqueDigits[k] = strings.ReplaceAll(uniqueDigits[k], display[3], "")

			if n_tries == 0 || n_tries == 2 || n_tries == 4 || n_tries == 6 {
				display[6] = string(uniqueDigits[k][0])
				display[2] = string(uniqueDigits[k][1])
			} else if n_tries == 1 || n_tries == 3 || n_tries == 5 || n_tries == 7 {
				display[2] = string(uniqueDigits[k][0])
				display[6] = string(uniqueDigits[k][1])
			}
			break
		case 7:
			// fmt.Println(uniqueDigits[k])

			uniqueDigits[k] = strings.ReplaceAll(uniqueDigits[k], display[1], "")
			uniqueDigits[k] = strings.ReplaceAll(uniqueDigits[k], display[3], "")

			display[0] = string(uniqueDigits[k][0])
			break
		case 8:
			// fmt.Println(uniqueDigits[k])

			uniqueDigits[k] = strings.ReplaceAll(uniqueDigits[k], display[1], "")
			uniqueDigits[k] = strings.ReplaceAll(uniqueDigits[k], display[3], "")
			uniqueDigits[k] = strings.ReplaceAll(uniqueDigits[k], display[0], "")
			uniqueDigits[k] = strings.ReplaceAll(uniqueDigits[k], display[6], "")
			uniqueDigits[k] = strings.ReplaceAll(uniqueDigits[k], display[2], "")

			if n_tries == 0 || n_tries == 1 || n_tries == 4 || n_tries == 5 {
				display[4] = string(uniqueDigits[k][0])
				display[5] = string(uniqueDigits[k][1])
			} else if n_tries == 2 || n_tries == 3 || n_tries == 6 || n_tries == 7 {
				display[5] = string(uniqueDigits[k][0])
				display[4] = string(uniqueDigits[k][1])
			}
			break
		}

	}
	// fmt.Println(display, uniqueDigits)

	return [][]string{
		// 0
		{display[0], display[1], display[3], display[4], display[5], display[6]},
		// 1
		{display[1], display[3]},
		// 2
		{display[0], display[1], display[2], display[5], display[4]},
		// 3
		{display[0], display[1], display[2], display[3], display[4]},
		// 4
		{display[6], display[2], display[1], display[3]},
		// 5
		{display[0], display[6], display[2], display[3], display[4]},
		// 6
		{display[0], display[6], display[2], display[5], display[4], display[3]},
		// 7
		{display[0], display[1], display[3]},
		// 8
		{display[0], display[1], display[2], display[3], display[4], display[5], display[6]},
		// 9
		{display[0], display[1], display[2], display[3], display[4], display[6]},
	}
}

func main() {
	scanner := bufio.NewReader(os.Stdin)

	final_res := 0
	for {
		entry, err := scanner.ReadString('\n')

		if err != nil {
			break
		}

		pos_sep := strings.Index(entry, "|")

		if pos_sep != -1 {
			uniqueDigits := findUniqueDigits(strings.Split(entry[:pos_sep-1], " "))
			// fmt.Println(uniqueDigits)
			out_val := strings.Split(entry[pos_sep+2:len(entry)-1], " ")

			// fmt.Println(out_val)

			keys := make([]int, 0, len(uniqueDigits))
			for k := range uniqueDigits {
				keys = append(keys, k)
			}
			sort.Ints(keys)

			is_not_ok := true
			n_tries := 0
			for is_not_ok {
				numbers := decodeDigits(keys, uniqueDigits, n_tries)

				// for i := range numbers {
				// 	fmt.Println(numbers[i], i)
				// }

				for i := range numbers {
					sort.Strings(numbers[i])
				}

				parz_res := ""
				for i := range out_val {
					val := strings.Split(out_val[i], "")
					sort.Strings(val)

					for x, n := range numbers {
						if len(n) == len(val) {
							isEqual := true
							// fmt.Println(val, n)

							for j := 0; j < len(val); j++ {
								if val[j] != n[j] {
									isEqual = false
									break
								}
							}

							if isEqual {
								parz_res += fmt.Sprint(x)
							}
						}
					}
				}

				if len(parz_res) != len(out_val) {
					n_tries++
				} else {
					is_not_ok = false
					// fmt.Println(parz_res)
					temp_res, _ := strconv.Atoi(parz_res)
					final_res += temp_res
				}

			}
		}
	}

	fmt.Println(final_res)
}
