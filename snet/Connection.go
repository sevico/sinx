package snet

import (
	"fmt"
	"github.com/sevico/sinx/siface"
	"net"
)

type Connection struct{
	Conn *net.TCPConn

	ConnID uint32

	isClosed bool

	//HandleAPI siface.HandleFunc
	ExitChan chan bool

	Router siface.IRouter
}

func (c *Connection) Reader(){
	fmt.Println("Reader goroutine is running...")
	defer fmt.Println("connID = ",c.ConnID," Reader is exit, remote addr is ",c.RemoteAddr().String())
	defer c.Stop()
	for  {
		buf:=make([]byte,512)
		_,err:=c.Conn.Read(buf)
		if err!=nil{
			fmt.Println("recv buf err ",err)
			continue
		}

		req:=&Request{
			conn: c,
			data: buf,
		}
		go func(request siface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(req)
		//if err:= c.HandleAPI(c.Conn,buf,cnt);err!=nil{
		//	fmt.Println("ConnID",c.ConnID," handle is error ",err)
		//	break
		//}

	}

}

func (c *Connection) Start() {
	fmt.Println("Conn Start()... ConnID= ",c.ConnID)
	go c.Reader()
}

func (c *Connection) Stop() {

	fmt.Println("Conn Stop()... ConnID= ",c.ConnID)
	if c.isClosed==true{
		return
	}
	c.isClosed=true
	c.Conn.Close()
	close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	return nil
}

func NewConnection(conn *net.TCPConn,connID uint32,router siface.IRouter) *Connection{
	c:=&Connection{
		Conn: conn,
		ConnID: connID,
		//HandleAPI: callbackAPI,
		isClosed: false,
		ExitChan: make(chan bool,1),
		Router:router,
	}
	return c

}