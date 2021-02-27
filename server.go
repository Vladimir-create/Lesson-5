package main

import (
	"fmt"
	"net"
	"encoding/binary"
	"bytes"
	"github.com/nsf/termbox-go"
	"time"
)

func main() {
	err:= termbox.Init()
	if err != nil {panic(err)}
	defer termbox.Close()
	adr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10236")
	if err != nil {
		fmt.Println(err)
		return
	}

	listener, err := net.ListenUDP("udp", adr)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		handleConnection(listener)
	}

}

func handleConnection(con *net.UDPConn) {
	buf := make([]byte, 2000)
	n, err := con.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	buff := bytes.NewReader(buf[0:n])

	var data struct {
		X int32
		Y int32
	}
	err = binary.Read(buff, binary.LittleEndian, &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	termbox.SetCell(int(data.X), int(data.Y), ' ', termbox.ColorDefault, termbox.ColorDefault)
	time.Sleep(1000*time.Millisecond)
}
