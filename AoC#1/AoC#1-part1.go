package main

import "fmt"

func main() {
	var forward, depth int
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
		case "down":
			depth += n
		case "up":
			depth -= n
		}
	}
	
	fmt.Println(forward * depth)
}
