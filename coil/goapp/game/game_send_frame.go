package game

import (
	// "bytes"
	"fmt"
	"net"
)

// draw 서버에 데이터를 보내기 위한 구조체
type client struct {
	Conn *net.Conn // client Connection
}

// Connect 서버와 연결
func (cli *client) connect(serverInfo string) {
	conn, err := net.Dial("tcp", serverInfo)
	if nil != err {
		fmt.Println("failed to connect to server")
	}
	cli.Conn = &conn
}

// SendData 서버에 맵 데이터 전송
func (cli *client) SendData(sendMapData [][]int, width int, height int) {

	if cli.Conn == nil {
		cli.connect("localhost:9999")
	}

	data := "[START]" + fmt.Sprintf("%d", width) + ":" + fmt.Sprintf("%d", height) + ":"

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			data += fmt.Sprintf("%d", sendMapData[i][j]) + ","
		}
	}

	data += ":" + "LEVEL 1" + ":LLDDUURR"
	data += "[END]"

	// fmt.Println(data)
	dataBuf := []byte(data)

	if _, err := (*cli.Conn).Write(dataBuf); nil != err {
		fmt.Println("SendData Error!!!")
	}

}
