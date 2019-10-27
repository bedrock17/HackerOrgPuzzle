package main_test

import (
	"testing"

	"./game"
)

// 게임코드 튜닝 시 검증 하기위한 코드

func testOption() {
	game.SetGoTestMode(true)
}
func test(w, h int, board, answer string, t *testing.T) {
	testOption()
	testAnswer, i, j := game.GetSolution(w, h, board)

	if testAnswer == answer {

	} else {
		t.Error(testAnswer, i, j)
	}
}

func Test_Game50(t *testing.T) {
	w := 21
	h := 19
	board := "X....X....X...XX...XX..XX...XX...X....X..X.X.......X....X.....X...X.......X...XX.X..X..X..X......X....X..X.XXX...........X....X......X....X.XX........X...XX.....XX.X....X..XX.........X.X....XX......X...X...XXX...X..XXX...X..X..XX....X.........X....X..X........XXX....X.X....X...XX.XX.....X..XX..X...XX....XX...X.X.....X.....X....X.X...X.......X...X..X.X..XX............XX..X..............X....XX...."
	answer := "RDRDDRDDDLDRDRLDDLRDLR"
	test(w, h, board, answer, t)
}
func Test_Game100(t *testing.T) {
	w := 37
	h := 36
	board := "......XXXXXXX..XXXXX...XXX............X.......XXXX......X.X..XX.XX...XX....X..XX.X..XXX...X..X..X....X..X..X....XX....XX...X...X.X...X......XXX..........X......X.XX....XXX..XX...XXXXX.....X....XX....XX.X...XXX....X.........X.XX.X..X.X.....XXX.X....X..X..........X..XX...X.XX..XX..X..X....X..X....X....XXX......X..XX.....X.XX..X.XXX....XX..XX.XX.X...XX....X...XXX...X.....X..........X....X.XX.......XX..X...........XX......X.X..X..XXXX.XX.....XXXXX........X......X.......X....XXX...X.......X....XX.X...XXXXX..XX...XXX....XX..X...X.....X...XXX....XX.X.XXX.X......X.XX..X...........X.....X......XXXXX....X......XX.XXX....XX...X...X...XX...X.XXX.....X...X.XXXX..X.X..XX....X.XXX...X...X.....X...X..X..X..X...X.X.XXX.....XXXX..X...X...X..XX....X.X...X....XXX..XX....X.....X.X...X.....XX....X...X...X...XX..XX...X.....X.X.....X...X...X.X...XXX..X..X...XXXX.XXX.X.X....XXXX.X...X....X....X..X...........XX..XXX...X...XX.X.XXXXX...XX.......XX...XXX........X...XX..XX......XXXX....XXX...XX...X.XX..XX.....XX......X...........X...X..X.....X....X..XXX.XXXXXXX...X...X..X...X.XXX.XX......X..X...XXXX......X.......XXX...X........X.X....XX...X.XX...X...XXX..X......XXX.X..X....X.....XX...X.....XXXX..XXX...X.....X...........X..XX..................XX.......XX...X.....XX....XXX...X.X...X.X...XXX..XX..XXXX............X...X...XXX.....XXX.......XX.."
	answer := "LLURLDRLDLLLLULLUDLULDULLLULULUDUUDULDDLDLURRRDRDDURDRLDDRDUDRRURDDLLDDD"
	test(w, h, board, answer, t)
}
