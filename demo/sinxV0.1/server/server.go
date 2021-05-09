package main

import (
	"fmt"
	"github.com/sevico/sinx/siface"
	"github.com/sevico/sinx/snet"
)

type PingRouter struct {

}

func (p *PingRouter) PreHandle(request siface.IRequest) {
	_,err:=request.GetConnetion().GetTCPConnection().Write([]byte("before ping...\n"))
	if err!=nil{
		fmt.Println("call back before ping error!")
	}

}

func (p *PingRouter) Handle(request siface.IRequest) {
	_,err:=request.GetConnetion().GetTCPConnection().Write([]byte("ping ping ping...\n"))
	if err!=nil{
		fmt.Println("call back ping error!")
	}
}

func (p *PingRouter) PostHandle(request siface.IRequest) {
	_,err:=request.GetConnetion().GetTCPConnection().Write([]byte("after ping...\n"))
	if err!=nil{
		fmt.Println("call back after ping error!")
	}
}

func main() {
	s:=snet.NewServer("[sinx V0.3]")


	s.AddRouter(&PingRouter{})

	s.Serve()
}
