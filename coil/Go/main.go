package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var gIsSolved = false
var width int
var height int
var wCount int //white count
var gDirection = []pos{pos{0, 1}, pos{1, 0}, pos{0, -1}, pos{-1, 0}}

var dpath = []string{"R", "D", "L", "U"}

//맵복사
func cpmap(m [][]int) [][]int {
	var ret [][]int
	for i := 0; i < height; i++ {
		var tmp []int
		for j := 0; j < width; j++ {
			tmp = append(tmp, m[i][j])
		}
		ret = append(ret, tmp)
	}
	return ret
}

type pos struct {
	i int
	j int
}

func isValid(i, j int) bool {
	if 0 <= i && i < height {
		if 0 <= j && j < width {
			return true
		}
	}
	return false
}

//완탐 재귀
func scan(m [][]int, i int, j int, depth int, path string, whiteCount int, num *int) {
	// fmt.Println("Scan pos ", i, j)
	if gIsSolved {
		return
	}

	if whiteCount < 30 {
		fmt.Println("DEBUG ====== ", *num, depth, path, whiteCount)
	}
	// for i := 0; i < height; i++ {
	// 	fmt.Println(m[i])
	// }

	for d := 0; d < 4; d++ {
		var log []pos
		ni, nj := i+gDirection[d].i, j+gDirection[d].j

		for isValid(ni, nj) && m[ni][nj] == 0 {
			m[ni][nj] = depth
			whiteCount--
			if whiteCount == 0 {

				fmt.Println("DEBUG ====== ", depth, "========path :", path)
				for i := 0; i < height; i++ {
					fmt.Println(m[i])
				}

				gIsSolved = true

			}
			log = append(log, pos{ni, nj})
			ni += gDirection[d].i
			nj += gDirection[d].j

		}
		// fmt.Println("-------------------")
		//탐색후 복구
		ni -= gDirection[d].i
		nj -= gDirection[d].j
		if len(log) > 0 {
			scan(m, ni, nj, depth+1, path+dpath[d], whiteCount, num)
			// fmt.Println("Scan end")
			for k := 0; k < len(log); k++ {
				// fmt.Println("remove", log[k].i, " ", log[k].j)
				m[log[k].i][log[k].j] = 0
				whiteCount++
			}
		}
	}
}

var goCount int = 0
var mutex = &sync.Mutex{}

//한 좌표당 한게임 고루틴으로 뺼것
func game(m [][]int, i int, j int) {
	mymap := cpmap(m)

	mymap[i][j] = 2

	num := i*10 + j
	mutex.Lock()
	goCount++
	fmt.Println("go start", i, j, goCount)
	mutex.Unlock()

	scan(mymap, i, j, 3, "", wCount-1, &num)

	mutex.Lock()
	goCount--
	fmt.Println("go end", i, j, goCount)
	mutex.Unlock()

}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS(0))

	width, _ = strconv.Atoi(os.Args[1])
	height, _ = strconv.Atoi(os.Args[2])
	board := os.Args[3]

	var m [][]int
	for i := 0; i < height; i++ {
		var tmp []int
		for j := 0; j < width; j++ {
			if board[i*width+j] == '.' {
				tmp = append(tmp, 0)
				wCount++
			} else {
				tmp = append(tmp, 1)
			}
		}
		m = append(m, tmp)

	}

	for i := 0; i < height; i++ {
		fmt.Println(m[i])
	}

	fmt.Scanln()

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if m[i][j] == 0 && gIsSolved == false {
				fmt.Println(i, j, "start")
				go game(m, i, j)
			}
		}
	}

	for gIsSolved == false {
		time.Sleep(1000)
		fmt.Println("wait", goCount)
	}

	fmt.Println(gIsSolved)

}
