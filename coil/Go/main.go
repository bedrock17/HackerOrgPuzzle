package main

import "fmt"
import "strconv"
import "os"

var gIsSolved bool = false
var width int
var height int
var wCount int //white count
var gDirection []pos

// func IsEnd() bool {
// 	return true
// }

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
var dpath string = "RDLU"

func scan(m [][]int, i int, j int, depth int, path string, whiteCount int) {
	// fmt.Println("Scan pos ", i, j)
	if gIsSolved {
		return
	}

	fmt.Println("DEBUG ====== ", depth)
	for i := 0; i < height; i++ {
		fmt.Println(m[i])
	}

	for d := 0; d < 4; d++ {
		var log []pos
		ni, nj := i+gDirection[d].i, j+gDirection[d].j

		for isValid(ni, nj) && m[ni][nj] == 0 {
			m[ni][nj] = depth
			whiteCount--
			if whiteCount == 0 {
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
			scan(m, ni, nj, depth+1, "", whiteCount)
			// fmt.Println("Scan end")
			for k := 0; k < len(log); k++ {
				// fmt.Println("remove", log[k].i, " ", log[k].j)
				m[log[k].i][log[k].j] = 0
				whiteCount++
			}
		}
	}
}

//한 좌표당 한게임 고루틴으로 뺼것
func game(m [][]int, i int, j int) {
	mymap := cpmap(m)

	mymap[i][j] = 2

	scan(mymap, i, j, 3, "", wCount-1)

}

func main() {

	width, _ = strconv.Atoi(os.Args[1])
	height, _ = strconv.Atoi(os.Args[2])
	board := os.Args[3]
	// = {{1, 0},{0, 1},{-1, 0},{0, -1}}
	gDirection = append(gDirection, pos{0, 1})
	gDirection = append(gDirection, pos{1, 0})
	gDirection = append(gDirection, pos{0, -1})
	gDirection = append(gDirection, pos{-1, 0})

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
	// fmt.Println("Hello" + "WOrld" + 'a')
	game(m, 1, 8)
	// game(m, 0, 0)
	// for i := 0; i < height; i++ {
	// 	fmt.Println(m[i])
	// }

	// fmt.Println("Hello WOrld", a, b)
	// fmt.Println(width, height, board)
	fmt.Println(gIsSolved)

}
