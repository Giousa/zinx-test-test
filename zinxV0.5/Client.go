package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zinx/znet"
)

func main() {

	fmt.Println("client start connecting...")
	time.Sleep(time.Second)
	conn,err := net.Dial("tcp","127.0.0.1:999")

	if err != nil{
		fmt.Println("Dial err ",err)
		return
	}

	for{



		dp := znet.NewDataPack()
		//发送数据

		sendMsg := znet.NewMsgPackage(0,[]byte("ZinxV0.5 send test msg"))

		bys,err := dp.Pack(sendMsg)
		if err != nil{
			fmt.Println("Client Pack msg err ",err)
			break
		}

		_,errW := conn.Write(bys)
		if errW != nil{
			fmt.Println("Client Write msg err ",err)
			break
		}


		//接收数据
		headData := make([]byte,dp.GetHeadLen())

		if _,err := io.ReadFull(conn,headData);err != nil{
			fmt.Println("client read head err ",err)
			break
		}

		headBody,err  := dp.Unpack(headData)
		if err != nil{
			fmt.Println("client Unpack head err ",err)
			break
		}

		dataLen := headBody.GetDataLen()

		if dataLen > 0{
			data := make([]byte,dataLen)

			if _,err := io.ReadFull(conn,data);err != nil{
				fmt.Println("client read data err ",err)
				break
			}

			fmt.Println("客户端读取数据：",string(data))
		}

		time.Sleep(time.Second)
	}

}
