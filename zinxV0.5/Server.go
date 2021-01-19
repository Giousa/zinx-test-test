package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

//处理conn业务主方法Hook
func (this *PingRouter) Handle(request ziface.IRequest){
	fmt.Println("Call PingRouter Handle")
	//先读取客户端数据，再写回ping..ping..ping..

	id := request.GetMsgID()
	data := request.GetData()

	fmt.Println("server recv id = ",id)
	fmt.Println("server recv data = ",string(data))

	err := request.GetConnection().SendMsg(1,[]byte("ping..ping..ping.."))
	if err != nil{
		fmt.Println(err)
	}

}


func main() {

	s := znet.NewServer("[zinx V0.5]")
	s.AddRouter(&PingRouter{})
	s.Run()
}
