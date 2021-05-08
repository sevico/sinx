package snet

import (
	"errors"
	"fmt"
	"github.com/sevico/sinx/siface"
	"net"
)

type Server struct {
	Name string
	IPVersion string
	IP string
	Port int
}

func CallBackToClient(conn *net.TCPConn,data []byte,cnt int) error{
	fmt.Println("[Conn Handle] CallbackToClient...")
	if _,err:=conn.Write(data[:cnt]);err!=nil{
		fmt.Println("write back buf err",err)
		return errors.New("CallBackToClient error")

	}
	return nil
}
//go func() {
//	for {
//		buf:=make([]byte,512)
//		cnt,err:=conn.Read(buf)
//		if err!=nil{
//			fmt.Println("Read err, ",err)
//			continue
//		}
//		fmt.Printf("recv client buf %s cnt %d\n",buf,cnt)
//		if _,err:= conn.Write(buf[:cnt]);err!=nil{
//			fmt.Println("Write err, ",err)
//			continue
//		}
//
//	}
//}()
func (s *Server) Start() {
	fmt.Printf("[Start] Server Listener at IP: %s, Port %d, is starting\n",s.IP,s.Port)

	go func() {
		addr,err:=net.ResolveTCPAddr(s.IPVersion,fmt.Sprintf("%s:%d",s.IP,s.Port))
		if err!=nil{
			fmt.Println("resolve tcp addr error: ",err)
			return
		}
		listener,err:=net.ListenTCP(s.IPVersion,addr)
		if err!=nil{
			fmt.Println("listen  ",s.IPVersion," err ",err)
			return
		}
		fmt.Println("start Sinx server succ, ",s.Name,", listening")
		var cid uint32
		cid=0

		for {

			conn,err:= listener.AcceptTCP()
			if err!=nil{
				fmt.Println("Accept err, ",err)
				continue
			}

			dealConn:=NewConnection(conn,cid,CallBackToClient)

			cid++
			dealConn.Start()
		}
	}()



}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()


	select {

	}
}


func NewServer(name string) siface.IServer{
	s:=&Server{
		Name: name,
		IPVersion: "tcp4",
		IP: "0.0.0.0",
		Port: 8999,
	}
	return s

}