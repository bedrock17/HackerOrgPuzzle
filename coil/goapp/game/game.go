package game

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"sync"
	"time"
)

//맵 배열의 값
//0 이하는 예약값
//1 이상은 벽들
const (
	EMPTYVAL            = 0
	BFSFILLVAL          = -1
	DEADPOINTVAL        = -2
	DEADPOINCHECKVAL    = -3
	DEADPOINCHECKOLDVAL = -4 //확정된 DEADPOINT
)

var gIsSolved = false
var width int
var height int
var wCount int //white count
var gDirection = []pos{pos{0, 1}, pos{1, 0}, pos{0, -1}, pos{-1, 0}}

var dpath = []string{"R", "D", "L", "U"}
var gDebugMode = true
var goTestMode = false
var maxGoRoutine int

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

func printMap(m [][]int, depth int) {
	s := ""
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			s += fmt.Sprintf("%x", m[i][j])
			if j < width-1 {
				s += ","
			}
		}
		if i < height-1 {
			s += "|"
		}
	}
}

//모양을 가리지 않고 bfs체크하면서 주변에 비어있는 길이 1개 이면 true 아니면 false
func isDeadPointCheck(m [][]int, i, j, count, nowVal int) bool {
	var pq posQueue
	var log []pos

	pq.create(count)

	m[i][j] = DEADPOINTVAL
	count--
	pq.put(pos{i, j})
	log = append(log, pos{i, j})

	emptyCount := 0
	for pq.length() > 0 {
		p := pq.get()
		for d := 0; d < 4; d++ {

			ni, nj := p.i+gDirection[d].i, p.j+gDirection[d].j

			if isValid(ni, nj) && (m[ni][nj] == EMPTYVAL || m[ni][nj] == nowVal || m[ni][nj] == DEADPOINCHECKOLDVAL) {
				emptyCount++
			}

			if isValid(ni, nj) && m[ni][nj] == DEADPOINCHECKVAL {
				m[ni][nj] = DEADPOINTVAL
				pq.put(pos{ni, nj})
				log = append(log, pos{ni, nj})
			}
		}

	}

	if emptyCount > 1 {
		for k := 0; k < len(log); k++ {
			m[log[k].i][log[k].j] = EMPTYVAL
		}
		return false
	}

	for k := 0; k < len(log); k++ {
		m[log[k].i][log[k].j] = DEADPOINCHECKOLDVAL
	}
	return true
}

//정사각형으로 영역을 확장하면서 사각형 모양이 가능한지 판단한다.
func rectCheck(m [][]int, i, j, width, height int) bool {
	for ii := i; ii < i+height; ii++ {
		for jj := j; jj < j+width; jj++ {
			if isValid(ii, jj) == false || m[ii][jj] != EMPTYVAL {
				return false
			}
		}
	}
	return true
}

//정사각형 모양으로 값을 채워준다
//나중에는 deadpoint 의심지점을 채워주는 기능으로 다시만들어야 함
func rectFill(m [][]int, i, j, width, height int) {
	for ii := i; ii < i+height; ii++ {
		for jj := j; jj < j+width; jj++ {
			m[ii][jj] = DEADPOINCHECKVAL
		}
	}
}

// m: 맵
// nowVal: 현재 진행중인 칸 값 (이값은 벽으로 치지 않는다)
// return: 탐색 불가능 하면 true
func deadPointGameOverCheck(m [][]int, nowVal int) bool {
	ret := false
	deadPointCnt := 0

	// 들어가면 나올 수 없는곳이 2개 이상인 경우
	for ii := 0; ii < height; ii++ {
		for jj := 0; jj < width; jj++ {
			if m[ii][jj] == EMPTYVAL {
				var cnt = 0
				for d := 0; d < 4; d++ {
					ni, nj := ii+gDirection[d].i, jj+gDirection[d].j
					if isValid(ni, nj) && (m[ni][nj] == EMPTYVAL || m[ni][nj] == nowVal) {
						cnt++
						if cnt > 1 {
							break
						}
					}
				}
				if cnt == 1 {
					deadPointCnt++
				}

				if deadPointCnt > 1 {
					return true
				}

			}
		}
	}

	for i := 0; i < height-1; i++ {
		for j := 0; j < width-1; j++ {
			if rectCheck(m, i, j, 2, 2) {
				rectwidth := 2
				rectheight := 2
				wcheck := true
				hcheck := true
				for wcheck || hcheck {
					wcheck = rectCheck(m, i, j, rectwidth+1, rectheight)

					if wcheck {
						rectwidth++
					}

					hcheck = rectCheck(m, i, j, width, rectheight+1)

					if hcheck {
						rectheight++
					}

				}
				rectFill(m, i, j, rectwidth, rectheight)

				if isDeadPointCheck(m, i, j, rectwidth*rectheight, nowVal) {

					deadPointCnt++

					if deadPointCnt > 1 {
						// if true {
						// 	printMap(m, nowVal)
						// 	fmt.Println("gameovercheck=================", deadPointCnt)
						// 	time.Sleep(time.Second * 3)
						// }
						ret = true
						goto END
					}
				}
			}
		}
	}

END:
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if m[i][j] == DEADPOINCHECKOLDVAL {
				m[i][j] = EMPTYVAL
			}
		}
	}
	return ret
}

//탐색 불가능한경우
// 가능 false
// 불가능 true
func gameOverCheck(m [][]int, i, j, whiteCount, nowVal int) bool {

	var pq posQueue
	var log []pos

	// 시작지점으로 돌아온 경우 탐색 가능
	var avaliable bool

	if m[i][j] != 0 {
		return true
	}

	//맵이 반토막 난경우
	pq.create(whiteCount)

	m[i][j] = BFSFILLVAL
	whiteCount--
	pq.put(pos{i, j})
	log = append(log, pos{i, j})

	for pq.length() > 0 {
		p := pq.get()
		for d := 0; d < 4; d++ {
			ni, nj := p.i+gDirection[d].i, p.j+gDirection[d].j

			// if ni == i && nj == j && false {
			// 	fmt.Println("========= !\n")
			// 	avaliable = true
			// 	goto OUT
			// }

			if isValid(ni, nj) && m[ni][nj] == 0 {
				m[ni][nj] = BFSFILLVAL
				whiteCount--
				pq.put(pos{ni, nj})
				log = append(log, pos{ni, nj})
			}
		}

	}

	// OUT:

	for k := 0; k < len(log); k++ {
		m[log[k].i][log[k].j] = EMPTYVAL
	}

	if avaliable {
		return true
	}

	return whiteCount != 0
}

var gCount int //디버깅용 count
var gi, gj int
var gqpath string

func scan(m [][]int, i int, j int, depth int, path string, whiteCount int, num *int) {

	if gIsSolved {
		return
	}

	if depth > 0 {
		if deadPointGameOverCheck(m, depth-1) {
			return
		}
	}

	var noLog = false
	cnt := 0
	for d := 0; d < 4; d++ {

		ni, nj := i+gDirection[d].i, j+gDirection[d].j
		if isValid(ni, nj) && m[ni][nj] == 0 {
			cnt++
		}

	}

	if cnt <= 1 {
		noLog = true
	}

	for d := 0; d < 4; d++ {
		var log []pos
		ni, nj := i+gDirection[d].i, j+gDirection[d].j

		if !isValid(ni, nj) {
			continue
		}

		if gameOverCheck(m, ni, nj, whiteCount, depth-1) {
			continue
		}

		for isValid(ni, nj) && m[ni][nj] == 0 {

			m[ni][nj] = depth
			whiteCount--
			if whiteCount == 0 {
				gIsSolved = true
				gqpath = path

				if !goTestMode {
					fmt.Println("DEBUG ====== ", depth, "========path :", path)
					for ii := 0; ii < height; ii++ {
						for jj := 0; jj < width; jj++ {
							if m[ii][jj] != 0 {

								if m[ii][jj] == 2 {
									gi = ii
									gj = jj
								}

								if m[ii][jj] == 1 {
									fmt.Printf("111 ")
								} else {
									fmt.Printf("%03X ", m[ii][jj])
								}
							} else {
								fmt.Printf("   ")
							}
						}
						fmt.Println("")
					}
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
			if noLog && depth > 3 {
				scan(m, ni, nj, depth+1, path, whiteCount, num)
			} else {
				//디버깅은 이곳으로!
				if gDebugMode && depth >= maxGoRoutine {
					if gCount%1 == 0 {
						// fmt.Println("DEBUG ====== ", *num, depth, path, whiteCount)
						// printMap(m, depth)
						// time.Sleep(3000 * time.Millisecond)
						gCount = 0
						maxGoRoutine = maxRoutineCheck()
					}
					gCount++
				}

				scan(m, ni, nj, depth+1, path+dpath[d], whiteCount, num)
			}
			for k := 0; k < len(log); k++ {
				m[log[k].i][log[k].j] = 0
				whiteCount++
			}
		}
	}
}

var goCount int
var mutex = &sync.Mutex{}
var chkMap [][]int
var chkCount int

//한 좌표당 한게임 고루틴으로 뺼것
func game(m [][]int, i int, j int) {
	mymap := cpmap(m)

	mymap[i][j] = 2
	startTime := time.Now()

	num := i*10 + j
	if gDebugMode {
		gCount = 0
	}
	scan(mymap, i, j, 3, "", wCount-1, &num)

	mutex.Lock()
	goCount--

	if !gIsSolved {
		elapsedTime := time.Since(startTime)
		if !goTestMode {
			fmt.Printf("go end %d %d %d %s\n", i, j, goCount, elapsedTime)
		}
		// fmt.Println(maxGoRoutine)
		chkMap[i][j] = 3
	}
	mutex.Unlock()

}

func maxRoutineCheck() int {
	dat, err := ioutil.ReadFile("./maxGoRoutine")
	if err == nil {
		// fmt.Println("?", string(dat))
		n, err := strconv.Atoi(string(dat))
		if err == nil {
			return n
		}
	}
	return 40
}

func gameInit() {
	gIsSolved = false
	width = 0
	height = 0
	wCount = 0
	maxGoRoutine = 0

	//옵션은 초기화 하지 않는다.
	// gDebugMode = true
	// goTestMode = false
}

//GetSolution game의 정답을 찾는다.
func GetSolution(w, h int, board string) (string, int, int) {

	gameInit()

	var solution string

	width = w
	height = h

	var m [][]int //map

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

	// for i := 0; i < height; i++ {
	// 	fmt.Println(m[i])
	// }

	chkMap = cpmap(m)
	maxGoRoutine = maxRoutineCheck()
	startTime := time.Now()

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			// for i := height - 1; i >= 0; i-- {
			// 	for j := width - 1; j >= 0; j-- {
			if m[i][j] == 0 && gIsSolved == false {
				// fmt.Println(i, j, "start ", i, j)

				for goCount > maxGoRoutine {
					time.Sleep(1 * time.Second)

					chkCount++
					if chkCount%20 == 0 {
						mutex.Lock()
						// printMap(chkMap, 3)
						elapsedTime := time.Since(startTime)
						fmt.Println("Spent", elapsedTime)
						fmt.Println("wait===2> ", i, j, goCount, maxGoRoutine, "[", (i*width + j), "]", "h w", height, width)
						maxGoRoutine = maxRoutineCheck()
						mutex.Unlock()
					}
				}

				if !goTestMode {
					fmt.Println("mutex wait")
				}
				mutex.Lock()
				goCount++

				if !goTestMode {
					fmt.Println("go start", i, j, goCount)
				}
				mutex.Unlock()
				if !goTestMode {
					fmt.Println("mutex end")
				}

				//성능 측정 및 디버깅 시 goRoutine을 사용하지 않으면 됨

				chkMap[i][j] = 2
				if gDebugMode {
					game(m, i, j)
				} else {
					go game(m, i, j)
				}
			}
		}
	}

	var wait time.Duration
	wait = 5000
	if goTestMode {
		wait = 1
	}

	for gIsSolved == false {
		time.Sleep(wait * time.Millisecond)

		if !goTestMode {
			fmt.Println("wait", goCount)
		}
	}

	if !goTestMode {
		time.Sleep(3 * time.Second)
	}

	elapsedTime := time.Since(startTime)
	fmt.Println("elapsed ", elapsedTime)

	solution = gqpath
	return solution, gi, gj
}
