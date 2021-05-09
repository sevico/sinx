package siface

type IRequest interface {
	GetConnetion() IConnection
	GetData() []byte
}