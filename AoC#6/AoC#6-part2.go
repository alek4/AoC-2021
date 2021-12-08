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

	avg := int(Avg(pos_crabs))

	res := calcFuel(pos_crabs, avg - 1)

	fmt.Println(res)
}

func calcFuel(slice []int, target int) int {
	tot_fuel := 0

	for _, n := range slice {
		dist := int(math.Abs(float64(n - target)))
		for i := dist; i > 0; i-- {
			tot_fuel += i
		}
	}

	return tot_fuel
}

func Avg(slice []int) float64 {
	sum := 0

	for _, n := range slice {
		sum += n
	}

	return float64(sum) / float64(len(slice))
}
