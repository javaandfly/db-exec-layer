package protocol

import (
	"time"

	"github.com/panjf2000/gnet/v2"
)

type ServerHandler func(ctx *HandlerContext)

type HandlerContext struct {
	ProtocolData *ProtocolData

	Conn gnet.Conn

	ExecuteMult bool

	Server *TcpDBServer
}

func InitServer(protocolData *ProtocolData, c gnet.Conn, server *TcpDBServer) *HandlerContext {
	return &HandlerContext{
		ProtocolData: protocolData,
		Conn:         c,
		Server:       server,
	}
}

func (ctx *HandlerContext) Deadline() (deadline time.Time, ok bool) {
	return
}

func (ctx *HandlerContext) Done() <-chan struct{} {
	return nil
}

func (ctx *HandlerContext) Err() error {
	return nil
}

func (ctx *HandlerContext) Value(key interface{}) interface{} {
	return nil
}
