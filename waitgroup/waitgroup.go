package main

import (
	"fmt"
)

func main() {
	var s = make(chan int,3)
	s <- 1
	s <- 2
	s <- 3
	fmt.Println(<-s)
	fmt.Println(<-s)
	fmt.Println(<-s)
	select  {
	case <-s : fmt.Println(<-s)
	default:
		fmt.Println("we")
	}
	fmt.Println(<-s)
}