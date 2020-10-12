package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS(0))

	width, _ = strconv.Atoi(os.Args[1])
	height, _ = strconv.Atoi(os.Args[2])
	mapString := os.Args[3]

	fmt.Println(width, height, mapString)

	GameProc(width, height, mapString, false)

}
