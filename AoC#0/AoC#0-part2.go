package main

import (
	"fmt"
)

func main() {
  var depth0, depth1, depth2 int
  
  prevSum := 99999
  count := 0
  fmt.Scanf("%d", &depth0)
  fmt.Scanf("%d", &depth1)
	for {
    _, err := fmt.Scanf("%d", &depth2)

    if err != nil {
      break
    }

    currSum := depth0 + depth1 + depth2

    if prevSum < currSum {
      count++
    }
    
    depth0 = depth1
    depth1 = depth2
    prevSum = currSum
  }

  fmt.Println("\ncounter", count)
}
