package protocol

import (
	"db-exec-layer/protocol/pb"
	"sync"
	"time"

	"github.com/panjf2000/gnet/v2"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

// 默认handler echo server
var defaultHandler ServerHandler = func(ctx *HandlerContext) {

	logrus.Info("defaultHandler is doing")

	if ctx.ProtocolData == nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("func is panic: %v  proto id is %v", r, ctx.ProtocolData.ProtoId)
		}
	}()

	fns, ok := ctx.Server.MethodsMap[ctx.ProtocolData.ProtoId]
	if !ok {
		logrus.Errorf("func is not find, proto id is %v", ctx.ProtocolData.ProtoId)
		return
	}

	var wg sync.WaitGroup

	if ctx.ExecuteMult {
		for index := range fns {
			wg.Add(1)
			//确保执行的是当前函数，不能用 _,fn := range fns 的方式
			fn := fns[index]

			go func() {
				defer wg.Done()
				fn(ctx)
			}()
		}
		wg.Wait()
		return

	}

	for _, fn := range fns {
		fn(ctx)
	}

}

func StartMult(ctx *HandlerContext) {
	ctx.ExecuteMult = true
}

func Ping(ctx *HandlerContext) {

	if ctx.ProtocolData == nil {
		return
	}

	protocol := NewTCPProtocol()

	req := &pb.DB_PING_REQ{}
	err := proto.Unmarshal(ctx.ProtocolData.Data, req)
	if err != nil {
		logrus.Errorf("read data is err:%v", err)
		return
	}

	logrus.Infof("req is read from server %v", req)

	rsp := &pb.DB_PING_RESP{
		TimeStamp: time.Now().UnixMilli(),
	}

	protoData, err := proto.Marshal(rsp)
	if err != nil {
		logrus.Errorf("paoro marshal pong , %v, err:%v", protoData, err)
		return
	}

	pongData, err := protocol.EncodeData(ACTION_PONG,
		protoData)
	if err != nil {
		logrus.Errorf("server encode pong , %v, err:%v", pongData, err)
		return
	}

	if ctx.Conn != nil {
		ctx.Conn.AsyncWrite(pongData, func(c gnet.Conn, err error) error { return nil })
	}
}

func HeartBeat(ctx *HandlerContext) {

	if ctx.ProtocolData == nil {
		return
	}

	protocol := NewTCPProtocol()

	req := &pb.DB_HEART_BEAT_REQ{}
	err := proto.Unmarshal(ctx.ProtocolData.Data, req)
	if err != nil {
		logrus.Errorf("read data is err:%v", err)
		return
	}

	logrus.Infof("req is read from server %v", req)

	rsp := &pb.DB_HEART_BEAT_RESP{
		Ticl:      req.Tick + 1,
		Timestamp: time.Now().UnixMilli(),
	}

	protoData, err := proto.Marshal(rsp)
	if err != nil {
		logrus.Errorf("paoro marshal pong , %v, err:%v", protoData, err)
		return
	}

	pongData, err := protocol.EncodeData(ACTION_HEART_BEAT_RESP,
		protoData)
	if err != nil {
		logrus.Errorf("server encode pong , %v, err:%v", pongData, err)
		return
	}

	if ctx.Conn != nil {
		ctx.Conn.AsyncWrite(pongData, func(c gnet.Conn, err error) error { return nil })
	}
}
