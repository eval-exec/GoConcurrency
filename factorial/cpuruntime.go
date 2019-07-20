package main

import (
	"fmt"
	"runtime"
)

func count() {
	cpuCount := runtime.NumCPU()
	goos := runtime.GOOS
	fmt.Println(cpuCount, goos)

}

