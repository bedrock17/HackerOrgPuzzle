package game

import "./gconst"

// i, j위치부터 비어있지 않은 부분을 채워준다.
// 그리고 비어있지 않은 부분 근처 비어있는 타일을 체크 타일로 바꿔준다
func makeCheckTile(m [][]int, whiteCount, i, j int) (checkTileLogs recoverLogArray) {
	var pq posQueue
	pq.create(width*height - whiteCount)
	pq.put(pos{i, j})

	logs := makeLogAraay(width*height - whiteCount)
	checkTileLogs = makeLogAraay(whiteCount)

	for pq.length() > 0 {
		p := pq.get()

		logs.append(recoverLog{p, m[p.i][p.j]})  //원래대로 돌리기위해 이전값도 넣어준다.
		m[p.i][p.j] = gconst.DISCONNECTCHECKUSED //visit

		for d := 0; d < 8; d++ {
			ni, nj := p.i+gconst.Direction8[d].I, p.j+gconst.Direction8[d].J

			if isValid(ni, nj) && m[ni][nj] != gconst.DISCONNECTCHECKUSED && m[ni][nj] != gconst.EMPTYVAL {

				pq.put(pos{ni, nj})

			} else if m[ni][nj] == gconst.EMPTYVAL {
				checkTileLogs.append(recoverLog{pos{ni, nj}, m[ni][nj]})
			}
		}
	}

	return checkTileLogs
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

	m[i][j] = gconst.BFSFILLVAL
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
			ni, nj := p.i+gconst.Direction[d].I, p.j+gconst.Direction[d].J

			if isValid(ni, nj) && m[ni][nj] == gconst.EMPTYVAL {
				m[ni][nj] = gconst.BFSFILLVAL
				whiteCount--
				pq.put(pos{ni, nj})
				// log = append(log, pos{ni, nj})
				log[logLen] = pos{ni, nj}
				logLen++

			}
		}

	}

	for k := 0; k < logLen; k++ {
		m[log[k].i][log[k].j] = gconst.EMPTYVAL
	}

	return whiteCount != 0
}
