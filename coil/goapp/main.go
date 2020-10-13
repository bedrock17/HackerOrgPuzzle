package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"

	"goapp/game"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS(0))

	width, _ := strconv.Atoi(os.Args[1])
	height, _ := strconv.Atoi(os.Args[2])
	board := os.Args[3]

	fmt.Println(width, height, board)

	game.SetDebugMode(false)
	path, i, j := game.GetSolution(width, height, board)

	file1, _ := os.Create("outurl") // outurl 파일 생성 (정답이 적혀있는 파일)
	defer file1.Close()             // main 함수가 끝나기 직전에 파일을 닫음
	fmt.Fprintf(file1, "http://www.hacker.org/coil/index.php?x=%d&y=%d&qpath=%s\n", j, i, path)

}
