package main

import "fmt"

func main() {
	var forward, depth, aim int
	var s string
	var n int

	for {
		fmt.Scanf("%s", &s)

		_, err := fmt.Scanf("%d", &n)

		if err != nil {
			break
		}

		switch s {
		case "forward":
			forward += n
			depth += aim * n
		case "down":
			aim += n
		case "up":
			aim -= n
		}
	}
	
	fmt.Println(forward * depth)
}
