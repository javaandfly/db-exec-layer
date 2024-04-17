package protocol

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants/v2"
	"github.com/panjf2000/gnet/v2"
	"github.com/sirupsen/logrus"
)

type TcpDBServer struct {
	Port       int
	WorkerPool *ants.Pool

	ConnNum int32

	Handler ServerHandler

	ctx context.Context
}

func (tcphs *TcpDBServer) OnBoot(eng gnet.Engine) (action gnet.Action) {
	logrus.Infof("server started.....")
	return
}

func (tcphs *TcpDBServer) OnShutdown(eng gnet.Engine) {
	logrus.Infof("server shutdown...... ")
}

func (tcphs *TcpDBServer) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	totalNum := atomic.AddInt32(&tcphs.ConnNum, 1)
	logrus.Infof(
		"total connection:%d, new connection: %s", totalNum, c.RemoteAddr())

	return
}

func (tcphs *TcpDBServer) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	atomic.AddInt32(&tcphs.ConnNum, -1)
	logrus.Infof("close connection: %s", c.RemoteAddr())

	return
}

func (tcphs *TcpDBServer) OnTick() (delay time.Duration, action gnet.Action) {
	return
}

func (s *TcpDBServer) OnTraffic(c gnet.Conn) (action gnet.Action) {

	protocol := NewTCPProtocol()

	protocolData, err := protocol.Decode(c)

	if err != nil {
		if err == ErrIncompletePacket {
			logrus.Errorf("ErrIncompletePacket Decode error :%+v\n", err)
			return gnet.None
		}
		logrus.Errorf("Protocol Decode error :%+v\n", err)

		return gnet.Close // 关闭连接
	}

	if protocolData == nil {

		logrus.Errorf("Protocol data is nil :%+v\n", err)

		return gnet.None
	}

	// 具体业务在 worker pool中处理
	s.WorkerPool.Submit(func() {
		handlerData := &HandlerContext{}
		handlerData.ProtocolData = protocolData
		handlerData.Conn = c
		handlerData.Server = s

		s.Handler(handlerData)
	})
	return
}

func NewTCPServer(port int) *TcpDBServer {
	options := ants.Options{ExpiryDuration: time.Second * 10, Nonblocking: true}
	defaultAntsPool, _ := ants.NewPool(DefaultAntsPoolSize, ants.WithOptions(options))

	server := &TcpDBServer{}

	server.Port = port
	server.WorkerPool = defaultAntsPool
	server.Handler = defaultHandler
	server.ctx = context.Background()

	return server
}

// 启动服务
func Run(server *TcpDBServer, protoAddr string, opts ...gnet.Option) error {

	return gnet.Run(server, protoAddr, opts...)
}
