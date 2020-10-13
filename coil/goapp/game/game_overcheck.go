package game

// i, j위치부터 비어있지 않은 부분을 채워준다.
// 그리고 비어있지 않은 부분 근처 비어있는 타일을 체크 타일로 바꿔준다
func makeCheckTile(m [][]int, i, j, whiteCount int) (checkTileLogs recoverLogArray) {
	var pq posQueue
	pq.create(width*height - whiteCount)
	pq.put(pos{i, j})

	recoverLogs := makeLogAraay(width*height - whiteCount)
	checkTileLogs = makeLogAraay(whiteCount)

	recoverLogs.append(recoverLog{pos{i, j}, m[i][j]}) //원래대로 돌리기위해 이전값도 넣어준다.
	m[i][j] = DISCONNECTCHECKUSED                      //visit
	for pq.length() > 0 {
		p := pq.get()

		// fmt.Println(len(recoverLogs.logs), recoverLogs.length)

		for d := 0; d < 8; d++ {
			ni, nj := p.i+Direction8[d].i, p.j+Direction8[d].j

			if isValid(ni, nj) && m[ni][nj] > 0 {
				recoverLogs.append(recoverLog{pos{ni, nj}, m[ni][nj]}) //원래대로 돌리기위해 이전값도 넣어준다.
				// fmt.Println(ni, nj, m[ni][nj], DISCONNECTCHECKUSED)
				m[ni][nj] = DISCONNECTCHECKUSED //visit
				pq.put(pos{ni, nj})

			} else if isValid(ni, nj) && m[ni][nj] == EMPTYVAL {
				checkTileLogs.append(recoverLog{pos{ni, nj}, m[ni][nj]})
				m[ni][nj] = DISCONNECTCHECKSTART
			}
		}
	}

	for i := 0; i < recoverLogs.length; i++ {
		m[recoverLogs.logs[i].pos.i][recoverLogs.logs[i].pos.j] = recoverLogs.logs[i].value
	}

	return checkTileLogs
}

// 탐색 불가능한경우
// 가능 false
// 불가능 true
func checkBFS(m [][]int, i, j, tileCount int, tileLog recoverLogArray) bool {

	var pq posQueue
	pq.create(tileCount)
	pq.put(pos{i, j})

	for pq.length() > 0 {
		p := pq.get()
		for d := 0; d < 4; d++ {
			ni, nj := p.i+Direction[d].i, p.j+Direction[d].j

			if isValid(ni, nj) && m[ni][nj] == DISCONNECTCHECKSTART {
				m[ni][nj] = DISCONNECTCHECKEND
				tileCount--
				pq.put(pos{ni, nj})

			}
		}
	}

	for i := 0; i < tileLog.length; i++ {
		m[tileLog.logs[i].pos.i][tileLog.logs[i].pos.j] = tileLog.logs[i].value
	}

	return tileCount != 0
}

// 탐색 불가능한경우
// 가능 false
// 불가능 true
func gameOverCheck(m [][]int, i, j, whiteCount, nowVal int) bool {

	var pq posQueue
	var log []pos

	if m[i][j] != 0 {
		return true
	}

	//맵이 반토막 난경우
	pq.create(whiteCount)

	m[i][j] = BFSFILLVAL
	whiteCount--
	pq.put(pos{i, j})

	log = make([]pos, whiteCount+1)
	logLen := 0
	// log = append(log, pos{i, j})
	log[logLen] = pos{i, j}
	logLen++

	for pq.length() > 0 {
		p := pq.get()
		for d := 0; d < 4; d++ {
			ni, nj := p.i+Direction[d].i, p.j+Direction[d].j

			if isValid(ni, nj) && m[ni][nj] == EMPTYVAL {
				m[ni][nj] = BFSFILLVAL
				whiteCount--
				pq.put(pos{ni, nj})
				// log = append(log, pos{ni, nj})
				log[logLen] = pos{ni, nj}
				logLen++

			}
		}

	}

	for k := 0; k < logLen; k++ {
		m[log[k].i][log[k].j] = EMPTYVAL
	}

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
					ni, nj := ii+Direction[d].i, jj+Direction[d].j
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
