package protocol

import (
	"db-exec-layer/protocol/pb"
	"time"

	"github.com/panjf2000/gnet/v2"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

type ServerHandler func(ctx *HandlerContext)

// 默认handler echo server
var defaultHandler ServerHandler = func(ctx *HandlerContext) {

	logrus.Info("defaultHandler is doing")

	if ctx.ProtocolData == nil {
		return
	}

	protocol := NewTCPProtocol()

	switch ctx.ProtocolData.ProtoId {
	case ACTION_PING:

		req := &pb.DB_PING_REQ{}
		err := proto.Unmarshal(ctx.ProtocolData.Data, req)
		if err != nil {
			logrus.Infof("read data is err:%v", err)
			break
		}

		logrus.Infof("req is read from server %v", req)

		if err != nil {
			logrus.Infof("read data is err:%v", err)
			break
		}

		rsp := &pb.DB_PING_RESP{
			TimeStamp: time.Now().UnixMilli(),
		}

		protoData, err := proto.Marshal(rsp)
		if err != nil {
			logrus.Infof("paoro marshal pong , %v, err:%v", protoData, err)
			break
		}

		pongData, err := protocol.EncodeData(ACTION_PONG,
			protoData)
		if err != nil {
			logrus.Infof("server encode pong , %v, err:%v", pongData, err)
			break
		}

		data, err := protocol.DecodeFrame(pongData)
		if err != nil {
			logrus.Infof("Decode pong , %v, err:%v", pongData, err)
			break
		}

		logrus.Infof("data is %#v", data)

		if ctx.Conn != nil {
			ctx.Conn.AsyncWrite(pongData, func(c gnet.Conn, err error) error { return nil })
		}
	}

	logrus.Infof("服务端收到数据, data:%s", string(ctx.ProtocolData.ProtoId))
}
