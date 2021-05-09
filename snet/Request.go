package snet

import "github.com/sevico/sinx/siface"

type Request struct {
	conn siface.IConnection
	data []byte
}

func (r *Request) GetConnetion() siface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}



