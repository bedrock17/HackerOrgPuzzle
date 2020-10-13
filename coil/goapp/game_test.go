package main_test

import (
	"testing"

	"goapp/game"
)

// 게임코드 튜닝 시 검증 하기위한 코드

func testOption() {
	game.SetDebugMode(true)
	game.SetGoTestMode(true)
}
func test(w, h int, board, answer string, t *testing.T) {
	testOption()
	testAnswer, i, j := game.GetSolution(w, h, board)

	if testAnswer == answer {
	} else {
		t.Error(answer, "["+testAnswer+"]", i, j)
	}
}
func Test_Game50(t *testing.T) {
	w := 21
	h := 19
	board := "X....X....X...XX...XX..XX...XX...X....X..X.X.......X....X.....X...X.......X...XX.X..X..X..X......X....X..X.XXX...........X....X......X....X.XX........X...XX.....XX.X....X..XX.........X.X....XX......X...X...XXX...X..XXX...X..X..XX....X.........X....X..X........XXX....X.X....X...XX.XX.....X..XX..X...XX....XX...X.X.....X.....X....X.X...X.......X...X..X.X..XX............XX..X..............X....XX...."
	answer := "RDRDDRDDDLDRDRLDDLRDLR"
	test(w, h, board, answer, t)
}
func Test_Game70(t *testing.T) {
	w := 27
	h := 26
	board := "XX......XX.....XXX............X....X..XXX.XXX.X...XXX..X.....XX..........X.X..XX.....XX...XX...X.........XX.XXX..X......X.X....XX......XX...X..X.......XX.X...XXXXX..X.........X.......X.......XX...X.X...X.XX...XXX.............X.X...XX.X.....X..XXX...X.....X..XX......X......XX...X.X....X.....X..........X..X......X......X...X..X.....XX....XX....X........X...X.XXXX.....X..XX...X.....X.....X..XXXXX.XX.............X....X...XX.XX.....XXXXXXX...X....X....XXX.........X......X....XXXXXX.X.X.....X........X.........X.....XXXXXXX...X..X.XXX....X....XX..X...XXX.X..X...XX...X...X....XXXXX.X.......XX....X...XXXX...X.......X....X..X..XX....X.X..XX...........XX....X..X.X...X.XX...........X..X..X...X...XX......."
	answer := "RUDRURRLDDDRDLRRDULDLDLDRRDRDLLDLDLLUURULL"
	test(w, h, board, answer, t)
}
func Test_Game100(t *testing.T) {
	w := 37
	h := 36
	board := "......XXXXXXX..XXXXX...XXX............X.......XXXX......X.X..XX.XX...XX....X..XX.X..XXX...X..X..X....X..X..X....XX....XX...X...X.X...X......XXX..........X......X.XX....XXX..XX...XXXXX.....X....XX....XX.X...XXX....X.........X.XX.X..X.X.....XXX.X....X..X..........X..XX...X.XX..XX..X..X....X..X....X....XXX......X..XX.....X.XX..X.XXX....XX..XX.XX.X...XX....X...XXX...X.....X..........X....X.XX.......XX..X...........XX......X.X..X..XXXX.XX.....XXXXX........X......X.......X....XXX...X.......X....XX.X...XXXXX..XX...XXX....XX..X...X.....X...XXX....XX.X.XXX.X......X.XX..X...........X.....X......XXXXX....X......XX.XXX....XX...X...X...XX...X.XXX.....X...X.XXXX..X.X..XX....X.XXX...X...X.....X...X..X..X..X...X.X.XXX.....XXXX..X...X...X..XX....X.X...X....XXX..XX....X.....X.X...X.....XX....X...X...X...XX..XX...X.....X.X.....X...X...X.X...XXX..X..X...XXXX.XXX.X.X....XXXX.X...X....X....X..X...........XX..XXX...X...XX.X.XXXXX...XX.......XX...XXX........X...XX..XX......XXXX....XXX...XX...X.XX..XX.....XX......X...........X...X..X.....X....X..XXX.XXXXXXX...X...X..X...X.XXX.XX......X..X...XXXX......X.......XXX...X........X.X....XX...X.XX...X...XXX..X......XXX.X..X....X.....XX...X.....XXXX..XXX...X.....X...........X..XX..................XX.......XX...X.....XX....XXX...X.X...X.X...XXX..XX..XXXX............X...X...XXX.....XXX.......XX.."
	answer := "LLURLDRLDLLLLULLUDLULDULLLULULUDUUDULDDLDLURRRDRDDURDRLDDRDUDRRURDDLLDDD"
	test(w, h, board, answer, t)
}
func Test_Game110(t *testing.T) {
	w := 41
	h := 39
	board := "X....XXXXX...........XXXX...XXXXX....X..X..XX.XXX..........XX........X.....XX.X..X.XXX.XX...XX...........X.X..X......X....X...X....X.....XX.X...X.X....XX...X....X..X....XXX....X..X...X...XXXX....X....X....X.X......XX..X.XX..X..XX......XXX...XX...X.X.X....X...X....XX..XX......XX..XXXX.X.....X.XX...X..X...XX...X.....XX..X.....X..X.XX....XX...X...X..X.X...X..X......XX.....X...X..X.......X......XXX..XXX.X..XX.XXXX..X.XX.....XX....X...X..X....X...X...X....XX...X....XX......XXX....X......X.XXX.X.....X...XXXXX.....XXX..XX.X........X.....X.X..X..........XXX...XXX.X.....X.XX.......XX...XX........X..X.......XXX...X.....XXXXXX.....XXX..X.X.X....XX.XXXX.....X.XXX....X....X....XX.X.X.X...X....X..X....X...XX.XXX..X....XX.X.....X....X.XX.....XX..XXX.....XX...XX.....X...X.....X....XXXXX.....XX.....X.X..X.....X...XX..X...........XX....X.XXX...XX..X.......XXXX..XX.X..XX...X.....XXXXXXXX....XXXXX....X.....X.....X...XX.....XXXXX......XXXX.....X.......X.XX...XXXXX...X.....XX.XXXX..X....XX..X......X.......X.X.XXX....XXX...X.X.....XX.XX.XXXX...X....X.XX...XXX...X...X.X...X..X..X..X.X...X....X..X..X....XX....X.XXX.X..XX..........X..X.XXX...X....XX........X.......X........XX...XX............XXX.....X......XXX.XX.X....X..X......X.........X.....XX...X..X...XXX...XXXX...........XXXXXX..........X.......X.XXXXXXX...X.........X....X............XX......X..X...X.....X..X........X.XXXXX.....XXX.X......X..X..XX............XXX....X......X...X..XXX..XXX..X......X.X.....XX..X......X....XXX.XXXX....X.XX...X.X......X...XX..XX...XXX....X.XX.X........X.....X..X....XX....XXX..X...XX..............X...XX..X....X..XXX"
	answer := "RDRRUDLUURRDRUURRRURURDLDRLDDLRDDRRUURRRURURRUURUUDUULLLLUDULLUULUDLDULLDLLURLLUULDLLDDDDLDDDL"
	test(w, h, board, answer, t)
}
func Test_Game120(t *testing.T) {
	w := 43
	h := 43
	board := "XXX....XXXXXX....XXXX..XXX..X.........X....XXX.XX..XXXXX.XX.X.....X....X.....XXX...XX.....X...X........X...X......XXXXX.............XX........X.XXXXX....X...XXXXXXX..X...XXX..X...X...X.X.X...XX.......X........X.X..XX.XX.X...X.....X.X..X.XX.XXXX......X......X..X..X..XX..X.......X.XX..XX...XX..XX...X...XX...X..XX..XXXX.......X....X....X...X.......XX....XXX.....XXXXX......XXX..X..X......XX....XX......XX.X.....XXXX...X....X...XXXXXXX......X..X..X...X..........X.X....X..XXX.........X...X..X..........X..X.....X...XXX...X....XX.X.XX...X...X....X.....X..X....XX.....X...X.....X..X......XXX...XX..X.....XX..X..X.X....X....XX....X.XX..X....XX.X.....X...X..X..X...X.....X.......X..X...X.X.........X.XX..X.....XXX...XXX.XXX.X..X.X..XXX...X..X............XX...XX..X.....X.....X.......X.....X.....X.........XX.XX.X..X.....XX......XXX.....X.......X....X.XX...X.............XXXX..X.....X......XX.X..XX....XXXX..XX..XXXXXXX...X.XXX.....XXX..X....X........XXX.........XXX..XXX..X....X.XX........XXX...X............X......XX.....X.....XXX...XXX.XXXXX....X....XXX...XX..X.XX.X...XXX....X..XXXXX.XX....X....XX.XXX.X..X...X...X..X...X......X.....XX...XX...X.XX.......X.X..XX........XX......X......X.X.XXX........X.....X.X....X.....X......X.X...XXXXX...X........X.X.X....XX..XXXX.XXX...XXXXXX..X...X.X....X.....XX....XXX...XXX.....XXXX........X.X..X..X.....X....X.....X.........XXXX...X.....X.....X..X........X........X.......X...X...X.X...X..XXX...XX...XX...X...X.X.....XXXXXX......X..XXX...XXX.....X...XX..X.XX.........X.X..X..XXX.X.....X........X.XX....X.....XXX.XXXXXX..X......X.......X.X...X...X.X...XXX......X.......X.X.......X.....X.X.X.XXXXXXXXX..X.X...........XXXXX....XX...X...XXXXXXXX...X.......XX.X....XXXX.X....XX...XXXXXXXX..X....X....XX.X.X.....X.X...XX..X..........X...X....XXXX.X.X...X.X..XX.XX...XX..........X....X......X...XXX....XX......XX......."
	answer := "LLURLDRLDLLLLULLUDLULDULLLULULUDUUDULDDLDLURRRDRDDURDRLDDRDUDRRURDDLLDDD"
	test(w, h, board, answer, t)
}
