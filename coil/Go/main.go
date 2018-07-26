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

//원형큐
type posQueue struct {
	q     []pos
	size  int
	start int
	end   int
}

func (pq *posQueue) create(size int) {
	pq.q = make([]pos, size)
	pq.size = size
	pq.start = 0
	pq.end = 0
}

func (pq *posQueue) put(p pos) {
	// fmt.Println(pq.end+1, pq.start, pq.end<pq.start)
	if pq.end+1 >= pq.start {
		pq.q[pq.end%pq.size] = p
		pq.end++
	} else {
		fmt.Println("PQ OVERFLOW!", p, pq.start, pq.end, pq.size)
		time.Sleep(2000 * time.Millisecond)
	}
}

func (pq *posQueue) get() pos {
	var p = pos{-1, -1}
	if pq.start+1 <= pq.end {
		p = pq.q[pq.start%pq.size]
		pq.start++
	} else {
		fmt.Println("PQ EMPTY")
	}
	return p
}
func (pq *posQueue) length() int {
	return pq.end - pq.start
}

//탐색 불가능한경우
// 가능 false
// 불가능 true
func gameOverCheck(m [][]int, i, j, whiteCount int) bool {

	var pq posQueue
	var log []pos

	if m[i][j] != 0 {
		return true
	}

	pq.create(whiteCount)

	m[i][j] = -1
	whiteCount--
	pq.put(pos{i, j})
	log = append(log, pos{i, j})

	for pq.length() > 0 {
		p := pq.get()
		for d := 0; d < 4; d++ {

			// fmt.Println("check : ", p)

			ni, nj := p.i+gDirection[d].i, p.j+gDirection[d].j

			for isValid(ni, nj) && m[ni][nj] == 0 {
				m[ni][nj] = -1
				whiteCount--
				pq.put(pos{ni, nj})
				log = append(log, pos{ni, nj})
			}
		}

	}
	// fmt.Println("DEBUG 2====== ", whiteCount, pq.length(), len(log))
	// for i := 0; i < height; i++ {
	// 	for j := 0; j < width; j++ {
	// 		if m[i][j] != 0 {
	// 			fmt.Printf("%2X ", m[i][j])
	// 		} else {
	// 			fmt.Printf("   ")
	// 		}
	// 	}
	// 	fmt.Println("")
	// 	// fmt.Println("")
	// }
	// time.Sleep(1000 * time.Millisecond)

	// if len(log) > 0 {
	// scan(m, ni, nj, depth+1, path+dpath[d], whiteCount, num)
	// fmt.Println("Scan end")
	for k := 0; k < len(log); k++ {
		// fmt.Println("remove", log[k].i, " ", log[k].j)
		m[log[k].i][log[k].j] = 0

	}
	// }
	// if whiteCount < 3 && whiteCount != 0 {
	// 	fmt.Println("====*> ", whiteCount)
	// 	fmt.Println(whiteCount)
	// time.Sleep(1 * time.Second)
	// }

	// if i == 1 && j == 1 {
	// fmt.Println("!!!!!!!!!!!!", i, j, whiteCount)
	// }
	return whiteCount != 0
}

// var gCount = 0

// var gi, gj int

func scan(m [][]int, i int, j int, depth int, path string, whiteCount int, num *int) {
	// fmt.Println("Scan pos ", i, j)
	if gIsSolved {
		return
	}

	// if gCount%1 == 0 && gi == 1 && gj == 8 {
	// fmt.Println("DEBUG ====== ", *num, depth, path, whiteCount)
	// for i := 0; i < height; i++ {
	// 	for j := 0; j < width; j++ {
	// 		if m[i][j] != 0 {
	// 			if m[i][j] == 1 {
	// 				fmt.Printf(".  ")
	// 			} else {
	// 				fmt.Printf("%02X ", m[i][j])
	// 			}
	// 		} else {
	// 			fmt.Printf("   ")
	// 		}
	// 	}
	// 	fmt.Println("")
	// fmt.Println("")
	// }
	// time.Sleep(100 * time.Millisecond)
	// gCount = 0
	// }

	for d := 0; d < 4; d++ {
		var log []pos
		ni, nj := i+gDirection[d].i, j+gDirection[d].j

		if !isValid(ni, nj) {
			continue
		}

		if gameOverCheck(m, ni, nj, whiteCount) {
			continue
		}
		for isValid(ni, nj) && m[ni][nj] == 0 {

			m[ni][nj] = depth
			whiteCount--
			if whiteCount == 0 {
				gIsSolved = true
				fmt.Println("DEBUG ====== ", depth, "========path :", path)
				for ii := 0; ii < height; ii++ {
					for jj := 0; jj < width; jj++ {
						if m[ii][jj] != 0 {
							if m[ii][jj] == 1 {
								fmt.Printf(".  ")
							} else {
								fmt.Printf("%02X ", m[ii][jj])
							}
						} else {
							fmt.Printf("   ")
						}
					}
					fmt.Println("")
					// fmt.Println("")
				}

			}
			log = append(log, pos{ni, nj})
			ni += gDirection[d].i
			nj += gDirection[d].j

		}

		//탐색후 복구
		if len(log) > 0 {
			ni -= gDirection[d].i
			nj -= gDirection[d].j
			scan(m, ni, nj, depth+1, path+dpath[d], whiteCount, num)
			for k := 0; k < len(log); k++ {
				m[log[k].i][log[k].j] = 0
				whiteCount++
			}
		}
	}
}

var goCount = 0
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

	if !gIsSolved {
		fmt.Println("go end", i, j, goCount)
	}
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
		time.Sleep(5000 * time.Millisecond)
		fmt.Println("wait", goCount)
	}
	time.Sleep(100 * time.Second)
	fmt.Println(gIsSolved)

}
