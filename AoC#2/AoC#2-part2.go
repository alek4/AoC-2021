package main

import (
	"fmt"
	"strconv"
)

func filterM(arr []string, index int) []string {
	arrFilter := []string{}
	c0 := 0
	c1 := 0

	if len(arr) == 1 {
		return arr
	}

	for i := range arr {
		if arr[i][index] == '0' {
			c0++
		} else {
			c1++
		}
	} 

	if c0 > c1 {
		for i := range arr {
			if arr[i][index] == '0' {
				arrFilter = append(arrFilter, arr[i])
			}
		}
	} else {
		for i := range arr {
			if arr[i][index] == '1' {
				arrFilter = append(arrFilter, arr[i])
			}
		}
	}

	index += 1

	return filterM(arrFilter, index)
}

func filterL(arr []string, index int) []string {
	arrFilter := []string{}
	c0 := 0
	c1 := 0

	if len(arr) == 1 {
		return arr
	}

	for i := range arr {
		if arr[i][index] == '0' {
			c0++
		} else {
			c1++
		}
	} 

	if c0 <= c1 {
		for i := range arr {
			if arr[i][index] == '0' {
				arrFilter = append(arrFilter, arr[i])
			}
		}
	} else {
		for i := range arr {
			if arr[i][index] == '1' {
				arrFilter = append(arrFilter, arr[i])
			}
		}
	}

	index += 1

	return filterL(arrFilter, index)
}

func main() {
	var input string
	var inputArr []string
	inputArr = make([]string, 0)

	for {
		_, err := fmt.Scanf("%s", &input)

		if err != nil {
			break
		}

		inputArr = append(inputArr, input)
	}

	mostCommon := filterM(inputArr, 0)
	leastCommon := filterL(inputArr, 0)

	fmt.Println(mostCommon, leastCommon)

	res, _ := strconv.ParseInt(mostCommon[0], 2, 64)
	res2, _ := strconv.ParseInt(leastCommon[0], 2, 64)

	fmt.Println(res * res2)
}
