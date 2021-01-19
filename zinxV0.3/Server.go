package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

//处理conn业务前的钩子方法Hook
func (this *PingRouter) PreHandle(request ziface.IRequest){
	fmt.Println("Call PingRouter PreHandle")
	_,err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
	if err != nil{
		fmt.Println("call back before ping err : ",err)
	}
}
//处理conn业务主方法Hook
func (this *PingRouter) Handle(request ziface.IRequest){
	fmt.Println("Call PingRouter Handle")
	_,err := request.GetConnection().GetTCPConnection().Write([]byte("handle ping...\n"))
	if err != nil{
		fmt.Println("call back handle ping err : ",err)
	}
}
//处理conn业务后钩子方法Hook
func (this *PingRouter) PostHandle(request ziface.IRequest){
	fmt.Println("Call PingRouter PostHandle")
	_,err := request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
	if err != nil{
		fmt.Println("call back after ping err : ",err)
	}
}

func main() {

	s := znet.NewServer("[zinx V0.3]")
	s.AddRouter(&PingRouter{})

	fmt.Println(s)
	s.Run()
}
