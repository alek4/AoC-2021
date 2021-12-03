package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	var bitMap []map[string]int

	bitMap = make([]map[string]int, 12)

	for {
		_, err := fmt.Scanf("%s", &input)

		if err != nil {
			break
		}

		for i := 0; i < len(input); i++ {
			if bitMap[i] == nil {
				bitMap[i] = make(map[string]int, 2)
			}

			if input[i] == '0' {
				bitMap[i]["0"]++
			} else {
				bitMap[i]["1"]++
			}
		}
	}

	gamma := ""
	epsilon := ""

	for i := 0; i < len(bitMap); i++ {
		if bitMap[i]["0"] > bitMap[i]["1"] {
			gamma = gamma + "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	// gammaInt, _ := strconv.Atoi(gamma)
	// epsilonInt, _ := strconv.Atoi(epsilon)

	res, _ := strconv.ParseInt(gamma, 2, 64)
	res2, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Println(res * res2)
}
