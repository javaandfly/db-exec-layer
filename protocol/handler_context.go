package protocol

import (
	"db-exec-layer/protocol/pb"
	"time"

	"github.com/panjf2000/gnet/v2"
)

type HandlerContext struct {
	ProtocolData *ProtocolData

	PingProtocolData *pb.DB_PING_REQ

	Conn gnet.Conn

	Server *TcpDBServer
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
