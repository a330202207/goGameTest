package main

import "gFrame/gnet"

func main() {
	//创建server句柄,使用
	s := gnet.NewServer("gInx V0.1")

	//启动server
	s.Server()
}
