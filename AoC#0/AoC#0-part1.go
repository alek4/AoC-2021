package main

import (
	"fmt"
)

func main() {
  var depthPrev, depthCurr int
  
  count := 0
  fmt.Scanf("%d", &depthPrev)
	for {
    _, err := fmt.Scanf("%d", &depthCurr)

    if err != nil {
      break
    }

    if depthCurr > depthPrev {
      count++
    }
    
    depthPrev = depthCurr
  }

  fmt.Println("\ncounter", count)
}
