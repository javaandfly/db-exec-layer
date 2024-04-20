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
		if err := recover(); err != nil {
			logrus.Errorf("func is panic: %v  proto id is %v", err, ctx.ProtocolData.ProtoId)
		}
	}()

	fns := ctx.MethodsMap[ctx.ProtocolData.ProtoId]

	var wg sync.WaitGroup

	if ctx.ExecuteMult {
		for index := range fns {
			wg.Add(1)

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
		logrus.Infof("read data is err:%v", err)
		return
	}

	logrus.Infof("req is read from server %v", req)

	if err != nil {
		logrus.Infof("read data is err:%v", err)
		return
	}

	rsp := &pb.DB_PING_RESP{
		TimeStamp: time.Now().UnixMilli(),
	}

	protoData, err := proto.Marshal(rsp)
	if err != nil {
		logrus.Infof("paoro marshal pong , %v, err:%v", protoData, err)
		return
	}

	pongData, err := protocol.EncodeData(ACTION_PONG,
		protoData)
	if err != nil {
		logrus.Infof("server encode pong , %v, err:%v", pongData, err)
		return
	}

	data, err := protocol.DecodeFrame(pongData)
	if err != nil {
		logrus.Infof("Decode pong , %v, err:%v", pongData, err)
		return
	}

	logrus.Infof("data is %#v", data)

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
		logrus.Infof("read data is err:%v", err)
		return
	}

	logrus.Infof("req is read from server %v", req)

	rsp := &pb.DB_HEART_BEAT_RESP{
		Ticl:      req.Tick + 1,
		Timestamp: time.Now().UnixMilli(),
	}

	protoData, err := proto.Marshal(rsp)
	if err != nil {
		logrus.Infof("paoro marshal pong , %v, err:%v", protoData, err)
		return
	}

	pongData, err := protocol.EncodeData(ACTION_HEART_BEAT_RESP,
		protoData)
	if err != nil {
		logrus.Infof("server encode pong , %v, err:%v", pongData, err)
		return
	}
}
