package main

import (
	"fmt"
	"zinx/znet"
)

func main() {

	s := znet.NewServer("[zinx V0.2]")

	fmt.Println(s)
	s.Run()
}
