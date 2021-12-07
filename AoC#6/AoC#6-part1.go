package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var input string
	pos_crabs := make([]int, 0)

	fmt.Scanf("%s", &input)

	parse_input := strings.Split(input, ",")

	for _, c := range parse_input {
		crab, _ := strconv.Atoi(c)
		pos_crabs = append(pos_crabs, crab)
	}

	sort.Ints(pos_crabs)

	fmt.Println(pos_crabs)

	avg := Avg(pos_crabs)

	occMap := make(map[int]int, 0)

	for _, n := range pos_crabs {
		if float64(n) < avg {
			occMap[n]++
		} else {
			break
		}
	}

	fmt.Println(occMap)

	target := findMostOcc(occMap)

	fmt.Println(target)

	res := calcFuel(pos_crabs, target)

	min := res

	flag := true

	for flag {
		target += 1
		res_temp := calcFuel(pos_crabs, target)

		if min > res_temp {
			min = res_temp
		}

		if min < res_temp {
			flag = false
		}
	}

	fmt.Println(min)
}

func calcFuel(slice []int, target int) int {
	tot_fuel := 0

	for _, n := range slice {
		tot_fuel += int(math.Abs(float64(n - target)))
	}

	return tot_fuel
}

func findMostOcc(map_ map[int]int) int {
	var pos int
	max := -1

	for i, v := range map_ {
		if v > max {
			max = v
			pos = i
		}
	}

	return pos
}

func Avg(slice []int) float64 {
	sum := 0

	for _, n := range slice {
		sum += n
	}

	return float64(sum) / float64(len(slice))
}