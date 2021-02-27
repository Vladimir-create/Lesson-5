package main

import (
	"fmt"
	"net"
	"encoding/binary"
	"bytes"
	"math/rand"
)
const width = 120 //ширина
const height = 30 //высота

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:12345")
	if err != nil {
		fmt.Println(err)
		return
	}

	var data struct {
		NumOfClient int32
		X int32
		Y int32
	}
	var Vx, Vy int32
	Vx, Vy = 1, 1
	data.X = int32(rand.Intn(width))
	data.Y = int32(rand.Intn(height))
	for i:=0; i<10000; i++{
		var buf bytes.Buffer
		err = binary.Write(&buf, binary.LittleEndian, data)
		_, err = conn.Write(buf.Bytes())
		if err != nil {
			fmt.Println(err)
			return
		}
		data.X, data.Y, Vx, Vy = change(data.X, data.Y, Vx, Vy)
	}
	conn.Close()
}

func change(x, y, Vx, Vy int32)(x1, y1, Vx1, Vy1 int32){
	if x + Vx > width -1{
		Vx = -1
	}
	if x + Vx < 0 {
		Vx = 1
	}	
	if y + Vy +1 > height {
		Vy = -1
	}
	if y + Vy < 0 {
		Vy = 1
	}
	return 	x + Vx, y+Vy, Vx, Vy			 
}
