package main

import (
	"fmt"
	"net"
	"encoding/binary"
	"bytes"
	"math/rand"
)
const width = 800
const height = 1200

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:10236")
	if err != nil {
		fmt.Println(err)
		return
	}

	var data struct {
		//NumOfClient int32
		X int32
		Y int32
	}
	data.X = int32(rand.Intn(width))
	data.Y = int32(rand.Intn(height))
	for i:=0; i<100; i++{
		var buf bytes.Buffer
		err = binary.Write(&buf, binary.LittleEndian, data)
		_, err = conn.Write(buf.Bytes())
		if err != nil {
			fmt.Println(err)
			return
		}
		data.X = int32(rand.Intn(width))
		data.Y = int32(rand.Intn(height))
	}	
	conn.Close()
}
