package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var init_state string
	state := make(map[int]int, 0)

	fmt.Scanf("%s", &init_state)

	state_s := strings.Split(init_state, ",")

	for _, f := range state_s {
		n, _ := strconv.Atoi(f)
		state[n]++
	}

	for x := 0; x < 256; x++ {
		six := 0
		eight := 0

		for i := 0; i < 9; i++ {
			if state[i] == 0 {
				continue
			}
			if i == 0 {
				six += state[i]
				eight += state[i]
			} else {
				state[i - 1] += state[i]
			}
			state[i] = 0
		}

		state[6] += six
		state[8] += eight
	}

	var res int64

	for _, v := range state {
		res += int64(v)
	}

	fmt.Println(res)
}
