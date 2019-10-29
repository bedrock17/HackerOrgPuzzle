package game

import "./gconst"

//탐색 불가능한경우
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
