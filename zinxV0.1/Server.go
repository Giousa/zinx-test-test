package main

import (
	"fmt"
	"zinx/znet"
)

func main() {

	s := znet.NewServer("[zinx V0.1]")

	fmt.Println(s)
	s.Run()
}
