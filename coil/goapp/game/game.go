package game

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"sync"
	"time"

	"./gconst"
)

var gIsSolved = false
var width int
var height int
var wCount int //white count

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

	m[i][j] = gconst.DEADPOINTVAL
	count--
	pq.put(pos{i, j})

	log = make([]pos, count*2)
	logLen := 0
	// log = append(log, pos{i, j})
	log[logLen] = pos{i, j}
	logLen++

	emptyCount := 0
	for pq.length() > 0 {
		p := pq.get()
		for d := 0; d < 4; d++ {

			ni, nj := p.i+gconst.Direction[d].I, p.j+gconst.Direction[d].J

			if isValid(ni, nj) && (m[ni][nj] == gconst.EMPTYVAL || m[ni][nj] == nowVal || m[ni][nj] == gconst.DEADPOINCHECKOLDVAL) {
				emptyCount++
			}

			if isValid(ni, nj) && m[ni][nj] == gconst.DEADPOINCHECKVAL {
				m[ni][nj] = gconst.DEADPOINTVAL
				pq.put(pos{ni, nj})
				// log = append(log, pos{ni, nj})
				log[logLen] = pos{ni, nj}
				logLen++
			}
		}
	}

	if emptyCount > 1 {
		for k := 0; k < logLen; k++ {
			m[log[k].i][log[k].j] = gconst.EMPTYVAL
		}
		return false
	}

	for k := 0; k < logLen; k++ {
		m[log[k].i][log[k].j] = gconst.DEADPOINCHECKOLDVAL
	}
	return true
}

//정사각형으로 영역을 확장하면서 사각형 모양이 가능한지 판단한다.
func rectCheck(m [][]int, i, j, width, height int) bool {
	for ii := i; ii < i+height; ii++ {
		for jj := j; jj < j+width; jj++ {
			if isValid(ii, jj) == false || m[ii][jj] != gconst.EMPTYVAL {
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
			m[ii][jj] = gconst.DEADPOINCHECKVAL
		}
	}
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

		ni, nj := i+gconst.Direction[d].I, j+gconst.Direction[d].J
		if isValid(ni, nj) && m[ni][nj] == 0 {
			cnt++
		}

	}

	if cnt <= 1 {
		noLog = true
	}

	disconnectionCheck := false //맵이 서로 연결되지 않았는지
	for d := 0; d < 4; d++ {
		max := func(x, y int) int {
			if x > y {
				return x
			} else {
				return y
			}
		}
		log := make([]pos, max(width, height))
		logLen := 0

		ni, nj := i+gconst.Direction[d].I, j+gconst.Direction[d].J

		if !isValid(ni, nj) {
			continue
		}

		if !disconnectionCheck {

			if whiteCount > (int)(float64(width*height)*0.6) {
				tileLog := makeCheckTile(m, i, j, whiteCount)
				if checkBFS(m, ni, nj, tileLog.length, tileLog) {
					continue
				}
			} else {
				if gameOverCheck(m, ni, nj, whiteCount, depth-1) {
					continue
				}
			}
			// if gameOverCheck(m, ni, nj, whiteCount, depth-1) {
			// 	continue
			// }
			disconnectionCheck = true //한쪽만 체크하면 된다.
		}

		for isValid(ni, nj) && m[ni][nj] == 0 {

			m[ni][nj] = depth
			whiteCount--

			log[logLen] = pos{ni, nj}
			logLen++

			ni += gconst.Direction[d].I
			nj += gconst.Direction[d].J

		}

		if whiteCount == 0 { // 정답!
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

		//탐색후 복구
		if logLen > 0 {
			ni -= gconst.Direction[d].I
			nj -= gconst.Direction[d].J
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

				scan(m, ni, nj, depth+1, path+gconst.Dpath[d], whiteCount, num)
			}
			for k := 0; k < logLen; k++ {
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
	gqpath = ""

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

	if !goTestMode {
		for gIsSolved == false {
			time.Sleep(wait * time.Millisecond)

			if !goTestMode {
				fmt.Println("wait", goCount)
			}
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
