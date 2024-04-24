package protocol

import (
	"context"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/panjf2000/ants/v2"
	"github.com/panjf2000/gnet/v2"
	"github.com/sirupsen/logrus"
)

type TcpDBServer struct {
	Port       int
	WorkerPool *ants.Pool

	ConnNum int32

	MethodsMap map[int32][]ServerHandler

	Handler ServerHandler

	ctx context.Context
}

func (tcphs *TcpDBServer) MethodRegist(id int32, fn ...ServerHandler) {
	tcphs.MethodsMap[id] = fn
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

	logrus.Info("on traffic is run")

	protocol := NewTCPProtocol()

	protocolData, err := protocol.DecodeProtocolData(c)

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
		handlerData := InitServer(protocolData, c, s)

		s.Handler(handlerData)

	})
	return
}

func NewTCPServer(port int) *TcpDBServer {
	options := ants.Options{ExpiryDuration: time.Second * 10, Nonblocking: true}
	defaultAntsPool, _ := ants.NewPool(DefaultAntsPoolSize, ants.WithOptions(options))

	server := &TcpDBServer{
		Port:       port,
		WorkerPool: defaultAntsPool,
		Handler:    defaultHandler,
		MethodsMap: make(map[int32][]ServerHandler),
		ctx:        context.Background(),
	}

	{
		server.MethodRegist(ACTION_PING, Ping)
		server.MethodRegist(ACTION_HEART_BEAT, HeartBeat)
	}

	return server
}

// 启动服务
func Run(ctx context.Context, server *TcpDBServer, protoAddr string, opts ...gnet.Option) error {

	go func() {
		// 监听系统信号量
		osSignal := make(chan os.Signal, 1)
		signal.Notify(osSignal, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

		for {
			s, ok := <-osSignal
			if !ok {
				break
			}
			switch s {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT,
				syscall.SIGHUP:
				err := gnet.Stop(ctx, protoAddr)
				if err != nil {
					logrus.Errorf("Failed to stop %s: %v", protoAddr, err)
				}

				return
			default:
			}
		}
	}()

	return gnet.Run(server, protoAddr, opts...)

}
