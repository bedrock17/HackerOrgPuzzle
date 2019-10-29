package gconst

//맵 배열의 값
//0 이하는 예약값
//1 이상은 벽들
const (
	EMPTYVAL             = 0
	BFSFILLVAL           = -1
	DEADPOINTVAL         = -2
	DEADPOINCHECKVAL     = -3
	DEADPOINCHECKOLDVAL  = -4 // 확정된 DEADPOINT
	DISCONNECTCHECKUSED  = -5 // 탐색이 끊겨버린 상황인지 확인하기 위한 타일
	DISCONNECTCHECKSTART = -6 // 탐색이 끊겨버린 상황인지 확인하기 위한 타일
	DISCONNECTCHECKEND   = -7 // 탐색이 끊겨버린 상황인지 확인하기 위한 타일
)

type pos struct {
	I int
	J int
}

//Direction 방향
var Direction = []pos{pos{0, 1}, pos{1, 0}, pos{0, -1}, pos{-1, 0}}

//Direction8 8방향
var Direction8 = []pos{
	pos{0, 1}, pos{1, 0}, pos{0, -1}, pos{-1, 0},
	pos{1, 1}, pos{-1, -1}, pos{1, -1}, pos{-1, 1}}

//Dpath 경로
var Dpath = []string{"R", "D", "L", "U"}
