package main

import (
	"fmt"
)

func tour() {
	c := make(chan int, 2)
	c <- 2
	c <- 1
	fmt.Println(c, c)

}
