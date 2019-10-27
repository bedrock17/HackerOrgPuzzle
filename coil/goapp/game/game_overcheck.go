package game

import "./gconst"

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

	m[i][j] = gconst.BFSFILLVAL
	whiteCount--
	pq.put(pos{i, j})
	log = append(log, pos{i, j})

	for pq.length() > 0 {
		p := pq.get()
		for d := 0; d < 4; d++ {
			ni, nj := p.i+gconst.Direction[d].I, p.j+gconst.Direction[d].J

			// if ni == i && nj == j && false {
			// 	fmt.Println("========= !\n")
			// 	avaliable = true
			// 	goto OUT
			// }

			if isValid(ni, nj) && m[ni][nj] == gconst.EMPTYVAL {
				m[ni][nj] = gconst.BFSFILLVAL
				whiteCount--
				pq.put(pos{ni, nj})
				log = append(log, pos{ni, nj})
			}
		}

	}

	for k := 0; k < len(log); k++ {
		m[log[k].i][log[k].j] = gconst.EMPTYVAL
	}

	if avaliable {
		return true
	}

	return whiteCount != 0
}
