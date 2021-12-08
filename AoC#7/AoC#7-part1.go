package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countDigits(values []string) int {
	res := 0

	for _, val := range values {
		switch len(val) {
		case 2, 4, 3, 7:
			res++
		}
	}

	return res
}

func main() {
	scanner := bufio.NewReader(os.Stdin)
	
	res := 0
	for {
		entry, err := scanner.ReadString('\n')

		if err != nil {
			break
		}

		pos_sep := strings.Index(entry, "|")

		if pos_sep != -1 {
			out_val := strings.Split(entry[pos_sep+2:len(entry)-1], " ")

			// fmt.Println(out_val)

			res += countDigits(out_val)
		}
	}

	fmt.Println(res)
}
