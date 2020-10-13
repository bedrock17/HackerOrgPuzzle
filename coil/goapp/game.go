package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
	"time"
)

//맵 배열의 값
//0 이하는 예약값
//1 이상은 벽들
const (
	EMPTYVAL            = 0
	BFSFILLVAL          = -1 + 1000000
	DEADPOINTVAL        = -2 + 1000000
	DEADPOINCHECKVAL    = -3 + 1000000
	DEADPOINCHECKOLDVAL = -4 + 1000000 //확정된 DEADPOINT
	FILLWALLVAL         = -5 + 1000000
	FILLVAL             = -6 + 1000250 //새로운 단절 판별을 위한 값
	FILLVALCHECK        = -7 + 1000600 //새로운 단절 판별을 위한 값
)

var gIsSolved = false
var width int
var height int
var wCount int //white count
var gDirection = []pos{{0, 1, 0}, {1, 0, 0}, {0, -1, 0}, {-1, 0, 0}}

var g8PointDirection = []pos{{0, 1, 0}, {1, 0, 0}, {0, -1, 0}, {-1, 0, 0}, {-1, -1, 0}, {-1, 1, 0}, {1, -1, 0}, {1, 1, 0}}

var dpath = []string{"R", "D", "L", "U"}
var gDebugMode = false
var gTestMode = false
var maxGoRutine int

var gMapStatus mapStatus

const statusPath string = "gameStatus.json"

//한쪽방향으로 주변영역을 채우는 bfs를 했을 때 주변에 빈칸이 있는지 검사한다.
//현재 무조건 O(n^2)
//TODO : 채워진영역주변만 검사 할 수 있도록 수정
// func fillCheck(m [][]int) bool {
// 	for i := 0; i < height; i++ {
// 		for j := 0; j < width; j++ {
// 			if m[i][j] > 1 {
// 				for k := 0; k < 4; j++ {
// 					ni, nj := i+gDirection[k].i, j+gDirection[k].j
// 					if isValid(ni, nj) && m[i][j] == EMPTYVAL {
// 						//채워지지 않은 타일을 찾은 것으로 맵이 두동강 난 상태
// 						return false
// 					}
// 				}
// 			}
// 		}
// 	}
// 	//빈자리를 찾을 수 없음
// 	return true
// }

//채워야 하는 영역인지 검사
func isFilleTile(m [][]int, i, j int) bool {
	for _, v := range g8PointDirection {
		ni, nj := i+v.i, j+v.j
		if m[ni][nj] > 1 {
			return true
		} else if m[ni][nj] == 1 {
			return isFilleTile(m, ni, nj)
		}
	}
	return false
}

// 주변영역을 bfs하고 주변에 빈칸이 있는지 검사
// 고립영역 체크를 위해 사용
// 탐색 가능 false
// 탐색 불가능 true
func fillAndCheck(m [][]int, i, j, whiteCount int) bool {
	var pq posQueue
	var log []pos
	pq.create(whiteCount)
	pq.put(pos{i, j, m[i][j]})
	log = append(log, pos{i, j, m[i][j]})

	m[i][j] = FILLWALLVAL

	lastNcount := 0
	outlineCount := 0

	// time.Sleep(3 * time.Second)
	for pq.length() > 0 {
		p := pq.get()

		for d := 0; d < 8; d++ {
			ni, nj := p.i+g8PointDirection[d].i, p.j+g8PointDirection[d].j
			if isValid(ni, nj) && m[ni][nj] == EMPTYVAL {
				log = append(log, pos{ni, nj, EMPTYVAL})
				m[ni][nj] = FILLVAL
				outlineCount++
			}
		}

		for d := 0; d < 8; d++ {
			ni, nj, nCount := p.i+g8PointDirection[d].i, p.j+g8PointDirection[d].j, p.count

			if nCount > lastNcount {
				// printMapPyserver(m, width, height)
				// fmt.Println(nCount)
				lastNcount = nCount
			}

			if isValid(ni, nj) && m[ni][nj] > EMPTYVAL && m[ni][nj] != FILLWALLVAL && m[ni][nj] != FILLVAL {

				pq.put(pos{ni, nj, nCount + 1})

				log = append(log, pos{ni, nj, m[ni][nj]})
				m[ni][nj] = FILLWALLVAL //어떤꺼가 태두리야?!?!?!?!?!?!?!?!?!?
			}
		}
	}

	var fpq posQueue
	startFillValueIndex := 0
	for k := 0; k < len(log); k++ {
		if m[log[k].i][log[k].j] == FILLVAL {
			startFillValueIndex = k
			break
		}
	}

	if len(log) > 0 {
		fpq.create(whiteCount)
		fpq.put(log[startFillValueIndex])
		m[log[startFillValueIndex].i][log[startFillValueIndex].j] = FILLVALCHECK
		outlineCount--

		for fpq.length() > 0 {
			p := fpq.get()
			for d := 0; d < 4; d++ {
				ni, nj, nCount := p.i+gDirection[d].i, p.j+gDirection[d].j, p.count

				if isValid(ni, nj) && m[ni][nj] == FILLVAL {
					fpq.put(pos{ni, nj, nCount + 1})
					m[ni][nj] = FILLVALCHECK
					outlineCount--
				}
			}
		}
	}

	// 맵 현재상태 디버깅용
	// if outlineCount != 0 {
	// 	printMapPyserver(m, width, height)
	// }

	for k := 0; k < len(log); k++ {
		m[log[k].i][log[k].j] = log[k].count
	}

	if outlineCount != 0 { //태두리를 모두 채우지 못했으므로 더이상 탐색 불가능
		return true
	}

	return false

}

//탐색 불가능한경우
// 가능 false
// 불가능 true
func gameOverCheck(m [][]int, i, j, si, sj, whiteCount, nowVal int) bool {

	var pq posQueue
	var log []pos
	var avaliable bool

	// if m[i][j] != 0 {
	// 	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!", i, j, m[i][j])
	// 	return true
	// }

	pq.create(whiteCount)

	m[i][j] = BFSFILLVAL
	whiteCount--
	pq.put(pos{i, j, 0})
	log = append(log, pos{i, j, 0})

	// fmt.Println("gameover check start", i, j, si, sj)
	// lastNcount := 0
	for pq.length() > 0 {
		p := pq.get()
		for d := 0; d < 4; d++ {
			ni, nj, nCount := p.i+gDirection[d].i, p.j+gDirection[d].j, p.count

			// if ni == si && nj == sj {
			// 	fmt.Println("=========> ", nCount, p.i, p.j, ni, nj, i, j)
			// 	printMap(m, nowVal)
			// 	time.Sleep(20 * time.Second)

			// }

			// if nCount > lastNcount {
			// printMapPyserver(m, width, height)
			// fmt.Println(nCount)
			// lastNcount = nCount
			// }

			// if false && nCount > 2 {
			// 	if nCount%10 != 0 && ni == si && nj == sj {
			// 		// fmt.Println(nCount, p.i, p.j, ni, nj)
			// 		avaliable = true
			// 		// printMap(m, nowVal)
			// 		// time.Sleep(10 * time.Second)
			// 		goto OUT
			// 	}
			// }

			if isValid(ni, nj) && m[ni][nj] == EMPTYVAL {
				m[ni][nj] = BFSFILLVAL
				whiteCount--

				pq.put(pos{ni, nj, nCount + 1})
				log = append(log, pos{ni, nj, nCount})
			}
		}
	}

	// OUT:
	for k := 0; k < len(log); k++ {
		m[log[k].i][log[k].j] = EMPTYVAL
	}

	if avaliable {
		//탐색이 가능하니 false
		return false
	}
	// fmt.Println("===============>", whiteCount)
	return whiteCount != 0
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

	// 성능 최적화중으로 사용하지 않음
	if true {
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

						hcheck = rectCheck(m, i, j, rectwidth, rectheight+1)

						if hcheck {
							rectheight++
						}

					}
					rectFill(m, i, j, rectwidth, rectheight)

					if isDeadPointCheck(m, i, j, rectwidth*rectheight, nowVal) {

						deadPointCnt++

						if deadPointCnt > 1 {

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
	}

	return ret
}

//맵복사
func cpmap(m [][]int, width, height int) [][]int {
	ret := make([][]int, height)
	var tmp []int
	for i := 0; i < height; i++ {
		tmp = make([]int, width)
		for j := 0; j < width; j++ {
			tmp[j] = m[i][j]
		}
		ret[i] = tmp
	}
	return ret
}

type pos struct {
	i     int
	j     int
	count int
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

func printMapPyserver(m [][]int, width, heith int) {
	// return
	if gDebugMode {
		if gIsSolved == false {
			gClient.SendData(m, width, height)
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func printMap(m [][]int, depth int) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if m[i][j] != 0 {
				if m[i][j] == 1 {
					fmt.Printf("■")
				} else if depth == m[i][j] {
					fmt.Printf("★")
				} else {
					fmt.Printf("%02X", m[i][j]%0x100)
				}
			} else {
				fmt.Printf("○")
			}
		}
		fmt.Printf(" %d\n", i)
	}
}

//모양을 가리지 않고 bfs체크하면서 주변에 비어있는 길이 1개 이면 true 아니면 false
func isDeadPointCheck(m [][]int, i, j, count, nowVal int) bool {
	var pq posQueue
	var log []pos

	pq.create(count)

	m[i][j] = DEADPOINTVAL
	count--
	pq.put(pos{i, j, 0})
	log = append(log, pos{i, j, 0})

	emptyCount := 0
	for pq.length() > 0 {
		p := pq.get()
		for d := 0; d < 4; d++ {

			ni, nj, nCount := p.i+gDirection[d].i, p.j+gDirection[d].j, p.count+1

			if isValid(ni, nj) && (m[ni][nj] == EMPTYVAL || m[ni][nj] == nowVal || m[ni][nj] == DEADPOINCHECKOLDVAL) {
				emptyCount++
			}

			if isValid(ni, nj) && m[ni][nj] == DEADPOINCHECKVAL {
				m[ni][nj] = DEADPOINTVAL
				pq.put(pos{ni, nj, nCount})
				log = append(log, pos{ni, nj, nCount})
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

//직사각형으로 영역을 확장하면서 사각형 모양이 가능한지 판단한다.
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

//직사각형 모양으로 채워준다.
//나중에는 deadpoint 의심지점을 채워주는 기능으로 다시만들어야 함
func rectFill(m [][]int, i, j, width, height int) {
	for ii := i; ii < i+height; ii++ {
		for jj := j; jj < j+width; jj++ {
			m[ii][jj] = DEADPOINCHECKVAL
		}
	}
}

var gCount int //디버깅용 count
var gClient client

var gi, gj int
var gqpath string

func scan(ms *mapStatus, i int, j int, depth int, path string, whiteCount int, num *int) {

	deadPointCnt := 0

	if gIsSolved {
		return
	}

	if depth > 0 {
		if deadPointGameOverCheck(ms.GameMap, depth-1) {
			return
		}
	}

	var noLog = false
	cnt := 0
	for d := 0; d < 4; d++ {

		ni, nj := i+gDirection[d].i, j+gDirection[d].j
		if isValid(ni, nj) && ms.GameMap[ni][nj] == 0 {
			cnt++
		}

	}

	if cnt <= 1 {
		noLog = true
	}

	// 더이상 사용되지 않기 떄문에 일단 무조건 사용한다고 체크
	BFScheckPass := true

	for d := 0; d < 4; d++ {
		var log []pos
		ni, nj := i+gDirection[d].i, j+gDirection[d].j

		if !isValid(ni, nj) {
			continue
		}

		if false { // BFS를 사용하여 맵이 단절되었는지 체크하는 부분
			if !noLog && BFScheckPass == false {
				if ms.GameMap[ni][nj] == 0 {
					if gameOverCheck(ms.GameMap, ni, nj, i, j, whiteCount, depth-1) {
						// fmt.Println("=============", depth)
						break
					}
					BFScheckPass = true
				}
			}
		}

		// BFScheckPass = true

		// JustGo
		for isValid(ni, nj) && ms.GameMap[ni][nj] == 0 {

			ms.GameMap[ni][nj] = depth
			whiteCount--
			if whiteCount == 0 {
				printMapPyserver(ms.GameMap, width, height)
				gIsSolved = true
				fmt.Println("DEBUG ====== ", depth, "========path :", path)
				gqpath = path
				for ii := 0; ii < height; ii++ {
					for jj := 0; jj < width; jj++ {
						if ms.GameMap[ii][jj] != 0 {

							if ms.GameMap[ii][jj] == 2 {
								gi = ii
								gj = jj
							}

							if ms.GameMap[ii][jj] == 1 {
								fmt.Printf("111 ")
							} else {
								fmt.Printf("%03X ", ms.GameMap[ii][jj])
							}
						} else {
							fmt.Printf("   ")
						}
					}
					fmt.Println("")
					// fmt.Println("")
				}

			}
			log = append(log, pos{ni, nj, 0})
			ni += gDirection[d].i
			nj += gDirection[d].j

		}

		// time.Sleep(1 * time.Second)
		// fmt.Println("DELAY")

		bCheck := true
		if len(log) > 0 {
			ni -= gDirection[d].i
			nj -= gDirection[d].j

			if noLog && depth > 3 {
				scan(ms, ni, nj, depth+1, path, whiteCount, num)
			} else {

				//디버깅은 이곳으로!
				// if gDebugMode && depth >= 500 {
				// 	if gCount%100 == 0 {
				// 		fmt.Println("DEBUG ====== ", *num, depth, path, whiteCount)
				// 		printMap(m, depth)
				// 		time.Sleep(3000 * time.Millisecond)
				// 		gCount = 0
				// 		maxGoRutine = maxRutineCheck()
				// 	}
				// 	gCount++
				// }
				printMapPyserver(ms.GameMap, width, height)

				if !noLog && BFScheckPass {
					if fillAndCheck(ms.GameMap, ni, nj, whiteCount) {
						// fmt.Println("FAIL!")
						bCheck = false
					}
				}

				if bCheck {
					scan(ms, ni, nj, depth+1, path+dpath[d], whiteCount, num)
				}
			}

			//탐색후 복구
			for k := 0; k < len(log); k++ {
				ms.GameMap[log[k].i][log[k].j] = 0
				whiteCount++
			}
		}
	}

	ms.deadPoint -= deadPointCnt
}

var goCount int
var mutex = &sync.Mutex{}
var chkMap [][]int
var chkCount int

//한 좌표당 한게임 고루틴으로 뺼것
func game(m [][]int, i int, j int) {

	mymap := cpmap(m, width, height)

	ms := mapStatus{width, height, mymap, 0}

	ms.GameMap[i][j] = 2
	startTime := time.Now()

	num := i*10 + j
	if gDebugMode {
		gCount = 0
	}
	scan(&ms, i, j, 3, "", wCount-1, &num)

	mutex.Lock()
	goCount--

	if !gIsSolved {
		elapsedTime := time.Since(startTime)
		if gTestMode == false {
			fmt.Printf("go end %d %d %d %s\n", i, j, goCount, elapsedTime)
		}
		chkMap[i][j] = 3
	}

	gMapStatus.GameMap[i][j] = 1
	if gTestMode == false {
		gMapStatus.SaveGmaeStatus(statusPath)
	}
	mutex.Unlock()

}

func maxRutineCheck() int {
	dat, err := ioutil.ReadFile("./maxGoRutine")
	if err == nil {
		fmt.Println(string(dat))
		n, err := strconv.Atoi(string(dat))
		if err == nil {
			return n
		}
	}
	return 40
}

func initGame() {
	gIsSolved = false
	gqpath = ""
	gMapStatus = mapStatus{0, 0, [][]int{}, 0}
	wCount = 0
}

// GameProc : 하나의 레벨을 처리하는 함수
// 성능테스트시 이 함수를 호출하여 확인한다.
func GameProc(pwidth, pheight int, mapString string, testMode bool) bool {

	width = pwidth
	height = pheight
	gTestMode = testMode
	initGame()

	var m [][]int
	for i := 0; i < height; i++ {
		var tmp []int
		for j := 0; j < width; j++ {
			if mapString[i*width+j] == '.' {
				tmp = append(tmp, 0)
				wCount++
			} else {
				tmp = append(tmp, 1)
			}
		}
		m = append(m, tmp)
	}
	gMapStatus.GameMap = cpmap(m, width, height)

	if fileExists(statusPath) {
		// 현재 맵이 저장되어 있는 맵과 다른경우만
		gMapStatus.LoadGameStatus(statusPath)

	}

	if gMapStatus.Width != width || gMapStatus.Heigtt != height || gDebugMode == true {
		gMapStatus.GameMap = cpmap(m, width, height)
		gMapStatus.Width = width
		gMapStatus.Heigtt = height
	}

	fmt.Println(len(gMapStatus.GameMap), len(gMapStatus.GameMap[0]))
	chkMap = cpmap(gMapStatus.GameMap, width, height)
	maxGoRutine = maxRutineCheck()
	startTime := time.Now()

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if gMapStatus.GameMap[i][j] == 0 && gIsSolved == false {

				for goCount > maxGoRutine {
					time.Sleep(1 * time.Second)

					chkCount++
					if true || gTestMode == false {
						if chkCount%20 == 0 {
							mutex.Lock()
							printMap(chkMap, 3)
							elapsedTime := time.Since(startTime)
							fmt.Println("Spent", elapsedTime)
							fmt.Println("wait===2> ", i, j, goCount, maxGoRutine, "[", (i*width + j), "]", "h w", height, width)
							maxGoRutine = maxRutineCheck()
							mutex.Unlock()
						}
					}
				}

				mutex.Lock()
				goCount++
				if gTestMode == false {
					fmt.Println("go start", i, j, goCount)
				}
				mutex.Unlock()

				//성능 측정 및 디버깅 시 gorutine을 사용하지 않으면 됨

				chkMap[i][j] = 2
				if gDebugMode {
					game(m, i, j)
				} else {
					go game(m, i, j)
				}
			}
		}
	}

	for gIsSolved == false {
		if gTestMode == false {
			time.Sleep(5000 * time.Millisecond)
			fmt.Println("wait", goCount)
		}
	}

	if gTestMode == false {
		os.Remove(statusPath)
		time.Sleep(10 * time.Second)
	}

	// fmt.Println(gj, gi, gqpath)
	// fmt.Printf("http://www.hacker.org/coil/index.php?x=%d&y=%d&qpath=%s\n", gj, gi, gqpath)

	file1, _ := os.Create("outurl") // hello1.txt 파일 생성
	defer file1.Close()             // main 함수가 끝나기 직전에 파일을 닫음
	fmt.Fprintf(file1, "http://www.hacker.org/coil/index.php?x=%d&y=%d&qpath=%s\n", gj, gi, gqpath)

	elapsedTime := time.Since(startTime)
	fmt.Println(gIsSolved)
	fmt.Println("elapsed ", elapsedTime)

	return gIsSolved
}
