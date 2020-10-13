package main

import "testing"

func TestGame1(t *testing.T) {
	width := 4
	height := 3
	mapString := `X..........X`

	GameProc(width, height, mapString, true)
}
func TestGame10(t *testing.T) {
	width := 7
	height := 6
	mapString := `.................XX..X..XX.....X......X...`

	GameProc(width, height, mapString, true)
}
func TestGame15(t *testing.T) {
	width := 8
	height := 8
	mapString := `...........XX......X......XX......XX......XX......X.......X.....`

	GameProc(width, height, mapString, true)
}
func TestGame50(t *testing.T) {
	width := 21
	height := 19
	mapString := `X....X....X...XX...XX..XX...XX...X....X..X.X.......X....X.....X...X.......X...XX.X..X..X..X......X....X..X.XXX...........X....X......X....X.XX........X...XX.....XX.X....X..XX.........X.X....XX......X...X...XXX...X..XXX...X..X..XX....X.........X....X..X........XXX....X.X....X...XX.XX.....X..XX..X...XX....XX...X.X.....X.....X....X.X...X.......X...X..X.X..XX............XX..X..............X....XX....`

	GameProc(width, height, mapString, true)
}
func TestGame100(t *testing.T) {
	width := 37
	height := 36
	mapString := `......XXXXXXX..XXXXX...XXX............X.......XXXX......X.X..XX.XX...XX....X..XX.X..XXX...X..X..X....X..X..X....XX....XX...X...X.X...X......XXX..........X......X.XX....XXX..XX...XXXXX.....X....XX....XX.X...XXX....X.........X.XX.X..X.X.....XXX.X....X..X..........X..XX...X.XX..XX..X..X....X..X....X....XXX......X..XX.....X.XX..X.XXX....XX..XX.XX.X...XX....X...XXX...X.....X..........X....X.XX.......XX..X...........XX......X.X..X..XXXX.XX.....XXXXX........X......X.......X....XXX...X.......X....XX.X...XXXXX..XX...XXX....XX..X...X.....X...XXX....XX.X.XXX.X......X.XX..X...........X.....X......XXXXX....X......XX.XXX....XX...X...X...XX...X.XXX.....X...X.XXXX..X.X..XX....X.XXX...X...X.....X...X..X..X..X...X.X.XXX.....XXXX..X...X...X..XX....X.X...X....XXX..XX....X.....X.X...X.....XX....X...X...X...XX..XX...X.....X.X.....X...X...X.X...XXX..X..X...XXXX.XXX.X.X....XXXX.X...X....X....X..X...........XX..XXX...X...XX.X.XXXXX...XX.......XX...XXX........X...XX..XX......XXXX....XXX...XX...X.XX..XX.....XX......X...........X...X..X.....X....X..XXX.XXXXXXX...X...X..X...X.XXX.XX......X..X...XXXX......X.......XXX...X........X.X....XX...X.XX...X...XXX..X......XXX.X..X....X.....XX...X.....XXXX..XXX...X.....X...........X..XX..................XX.......XX...X.....XX....XXX...X.X...X.X...XXX..XX..XXXX............X...X...XXX.....XXX.......XX..`

	GameProc(width, height, mapString, true)
}
func TestGame120(t *testing.T) {
	width := 43
	height := 43
	mapString := `XXX....XXXXXX....XXXX..XXX..X.........X....XXX.XX..XXXXX.XX.X.....X....X.....XXX...XX.....X...X........X...X......XXXXX.............XX........X.XXXXX....X...XXXXXXX..X...XXX..X...X...X.X.X...XX.......X........X.X..XX.XX.X...X.....X.X..X.XX.XXXX......X......X..X..X..XX..X.......X.XX..XX...XX..XX...X...XX...X..XX..XXXX.......X....X....X...X.......XX....XXX.....XXXXX......XXX..X..X......XX....XX......XX.X.....XXXX...X....X...XXXXXXX......X..X..X...X..........X.X....X..XXX.........X...X..X..........X..X.....X...XXX...X....XX.X.XX...X...X....X.....X..X....XX.....X...X.....X..X......XXX...XX..X.....XX..X..X.X....X....XX....X.XX..X....XX.X.....X...X..X..X...X.....X.......X..X...X.X.........X.XX..X.....XXX...XXX.XXX.X..X.X..XXX...X..X............XX...XX..X.....X.....X.......X.....X.....X.........XX.XX.X..X.....XX......XXX.....X.......X....X.XX...X.............XXXX..X.....X......XX.X..XX....XXXX..XX..XXXXXXX...X.XXX.....XXX..X....X........XXX.........XXX..XXX..X....X.XX........XXX...X............X......XX.....X.....XXX...XXX.XXXXX....X....XXX...XX..X.XX.X...XXX....X..XXXXX.XX....X....XX.XXX.X..X...X...X..X...X......X.....XX...XX...X.XX.......X.X..XX........XX......X......X.X.XXX........X.....X.X....X.....X......X.X...XXXXX...X........X.X.X....XX..XXXX.XXX...XXXXXX..X...X.X....X.....XX....XXX...XXX.....XXXX........X.X..X..X.....X....X.....X.........XXXX...X.....X.....X..X........X........X.......X...X...X.X...X..XXX...XX...XX...X...X.X.....XXXXXX......X..XXX...XXX.....X...XX..X.XX.........X.X..X..XXX.X.....X........X.XX....X.....XXX.XXXXXX..X......X.......X.X...X...X.X...XXX......X.......X.X.......X.....X.X.X.XXXXXXXXX..X.X...........XXXXX....XX...X...XXXXXXXX...X.......XX.X....XXXX.X....XX...XXXXXXXX..X....X....XX.X.X.....X.X...XX..X..........X...X....XXXX.X.X...X.X..XX.XX...XX..........X....X......X...XXX....XX......XX.......`

	GameProc(width, height, mapString, true)
}

// func TestGame150(t *testing.T) {
// 	width := 53
// 	height := 53
// 	mapString := `..X.....XXXX..........XXXXX..XX..X.......X..........X....X.....XX.XX...XX.....XX...X....X...X...XXXX.....X...X.......X....X.X..X.X....X...XXXX.X.X..X..XX..XXXX...X.X...X.X...XX...X..X..X....X.....X....X..XXX......X....XX.X...X...X..X........XXX....XXXXX....XXXX..X..X..X....XX...XX...X.........X....X.......XX...X...X.....X..X..X.X......X.X.........XX....X.........X.X...XXX..XX...X...X......X..XXXXX......X...X...XX.......XX....XX.X..X..X...XXXXX...X.....X......X......X...X..X..........X.XXXX..XXXXX....XXX.....X.............XX...XXXX...X.....X...XXXXX.X.......XX.....XXXXXX..XX.......XXXX.........XXXX.......X..X.....XXX...XX...X........X.....X...XXXXX.....X...X....X...XXX.X....X...X........X.....X..XXX...X......X......X.......X..XX..X...........X.X...X...X.....X..X...XX.X....X.....XX....X..XXXX.X...X.X.....XX.....XXX.X....X.X..X.....XX.XX.....X...X..X..XXXX..X....X...X.X...X..X.X.....XXX.XXX..X...XXX....XXX....X.X........X.....X....XXX...X...X..X..............X..X.....X...XX..XXXX.X............X.....XXX.X..X....X.X...XX...XX.........X...X...XX.....X.X.XXX....X.XX.....XXXX....X..XXX...XX.X...XX.........X......X.....X.....XXX....XX.XX..X....X.......XX......X...............X....X....X....XXXXXXX...X...XX...XX.....XX.......XX......X.XX.XX..X...XX.......X.X....X...X.......X....XX.....X..X.....X.X....XX.X....X.X..X.....X.XX.....X.....XX..X....XX........X..X.X....X........X....X.X.....X.....X..X.....X.XXX.X....X...XXXX......X...XX....X...X....XXXXXXX..X..X....X....X.....XXXXXX..X.X...X...X....X.........XXX...X....X...X........X..XX...XX....XX.X........X.....X....XX...X...X...XX.X....X...X.X.....XX.........XX....X......XX..XX...XX......X.X.....X.......XXXX......X.....X.X..X.....X..X..XXXX............X.X...X..XXX....X...X..X...X...XX.........X.......X....XXX..........XX......XXX....XX.....XXX......X...X.X....X..X......XX...XX...X...X............X......XXX..XXX....X......X..X....X...X..X......X......XXX.....X......X....X...X.XXXX.....XXX...X........X..XX....X.XX........X...X........X.....XXXXX...XXX......X.XX...XX....X...X......X.XX....................X....X.XXX...XX...XXX....X.....XX.X.....X.....XXX...XX...X..X...X....XX.....X....X.....X..........XXX..X....X.X.....XXXX..XX.XXX....X...XX.........X...XXX.............X..XXX..X..XXX...........X....X....XXXX.....XXX...X...X...X..X...XX...X.....X...XX...X....X..X.....X.....X..........X...XX.X...XXXX....X....X.......XX..XXXX..XX...XXXXX.....XX...X..XXX.X....XX...X....XXX...XXX..XXXX..XX....X...X......X.....X...XX.....X.....X...XXXXXX...XX.XX...X......X....X....X....X...X..........XX....X........XXX...XX.....X....XXX..X.......XX...X....XX...X.XX...XXXXX....X...XX...XXX..X..X.XX....XXXXX...X...X.X....X..X.......X........X..X....XX..XXX.....X.....X.X.X..X....XX..........X.....X............XX..XXXX..X...X..X.......X..........X...XXX`

// 	GameProc(width, height, mapString, true)
// }
