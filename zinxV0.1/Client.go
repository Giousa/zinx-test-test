package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("client start connecting...")
	time.Sleep(time.Second)
	conn,err := net.Dial("tcp","127.0.0.1:999")

	if err != nil{
		fmt.Println("Dial err ",err)
		return
	}

	buf := make([]byte,1024)
	for{
		conn.Write([]byte("Hello Zinx V1.0..."))

		n,err := conn.Read(buf)
		if err != nil{
			fmt.Println("Read err ",err)
			return
		}

		fmt.Println("Client Read: ",string(buf[:n]))

		time.Sleep(time.Second*2)

	}

}
