package protocol

import "github.com/panjf2000/gnet"

type echoServer struct {
	*gnet.EventServer
}

func (es *echoServer) Ping(c gnet.Conn) (out []byte, action gnet.Action) {
	out = c.Read()
	c.ResetBuffer()
	return
}
