package protocol

import (
	"time"

	"github.com/panjf2000/gnet/v2"
)

type ServerHandler func(ctx *HandlerContext)

type HandlerContext struct {
	ProtocolData *ProtocolData

	MethodsMap map[int32][]ServerHandler

	Conn gnet.Conn

	Server *TcpDBServer
}

func InitServer(protocolData *ProtocolData, c gnet.Conn, server *TcpDBServer) *HandlerContext {
	return &HandlerContext{
		ProtocolData: protocolData,
		Conn:         c,
		Server:       server,
	}
}

func (ctx *HandlerContext) MethodRegist(id int32, fn ...ServerHandler) {
	ctx.MethodsMap[id] = fn
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
