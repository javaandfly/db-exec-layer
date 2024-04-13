package protocol

import (
	"bytes"
	"db-exec-layer/protocol/pb"
	"encoding/binary"
	"time"

	"github.com/panjf2000/gnet/v2"
	"github.com/sirupsen/logrus"
)

type ServerHandler func(ctx *HandlerContext)

// 默认handler echo server
var defaultHandler ServerHandler = func(ctx *HandlerContext) {
	if ctx.ProtocolData == nil {
		return
	}

	protocol := NewTCPProtocol()

	switch ctx.ProtocolData.ProtoId {
	case ACTION_PING:

		buf := &bytes.Buffer{}
		_, err := buf.Read(ctx.ProtocolData.Data)
		if err != nil {
			logrus.Infof("read data is err:%v", err)
			break
		}

		req := &pb.DB_PING_REQ{}
		err = binary.Read(buf, binary.BigEndian, &req)

		logrus.Infof("req is read from server %v", req)

		if err != nil {
			logrus.Infof("read data is err:%v", err)
			break
		}

		rsp := &pb.DB_PING_RESP{
			TimeStamp: time.Now().UnixMilli(),
		}

		pongData, err := protocol.EncodeData(ACTION_PONG,
			[]byte((rsp.String())))

		if err != nil {
			logrus.Infof("server encode pong , %v, err:%v", pongData, err)
			break
		}

		if ctx.Conn != nil {
			ctx.Conn.AsyncWrite(pongData, func(c gnet.Conn, err error) error { return nil })
		}
	}

	logrus.Infof("服务端收到数据, data:%s", string(ctx.ProtocolData.ProtoId))
}
