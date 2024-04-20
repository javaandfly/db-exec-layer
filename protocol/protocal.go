package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/panjf2000/gnet/v2"
)

type ProtocolData struct {
	DataLength int32
	ProtoId    int32
	Data       []byte

	//headDecode bool
	//Lock       sync.Mutex
}

// 协议头长度
func (p *ProtocolData) HeadLength() int {
	return DefaultHeadLength
}

type TCPProtocol struct{}

// new protocal
func NewTCPProtocol() *TCPProtocol {
	return &TCPProtocol{}
}

func (tcpfhp *TCPProtocol) getHeadLength() int {
	return DefaultHeadLength
}

// server端 gnet input 数据 decode
func (tcpfhp *TCPProtocol) DecodeProtocolData(c gnet.Conn) (*ProtocolData, error) {
	curConContext := c.Context()

	if curConContext == nil {
		//解析协议 header
		tempBufferLength := c.InboundBuffered()        // 当前已有多少数据
		if tempBufferLength < tcpfhp.getHeadLength() { // 不够头长度
			return nil, ErrIncompletePacket
		}

		headData, _ := c.Next(tcpfhp.getHeadLength())

		newConContext := &ProtocolData{}

		//数据长度
		bytesBuffer := bytes.NewBuffer(headData)
		binary.Read(bytesBuffer, binary.LittleEndian, &newConContext.DataLength)
		binary.Read(bytesBuffer, binary.LittleEndian, &newConContext.ProtoId)

		c.SetContext(newConContext)
	}

	protocolData := &ProtocolData{}

	//解析协议数据
	protocolData, ok := c.Context().(*ProtocolData)
	if !ok {
		c.SetContext(nil)

		return nil, ErrContext
	}
	tempBufferLength := c.InboundBuffered() // 当前已有多少数据
	frameDataLength := int(protocolData.DataLength)

	if tempBufferLength > frameDataLength {
		return nil, ErrIncompletePacket
	}

	// 数据够了
	data, _ := c.Next(tempBufferLength)

	copyData := make([]byte, tempBufferLength) // 复制
	copy(copyData, data)

	protocolData.Data = copyData

	c.SetContext(nil)

	return protocolData, nil

}

// 数据反解
func (tcpfhp *TCPProtocol) DecodeFrame(frame []byte) (*ProtocolData, error) {
	data := &ProtocolData{}
	//数据长度
	bytesBuffer := bytes.NewBuffer(frame)

	if err := binary.Read(bytesBuffer, binary.LittleEndian, &data.DataLength); err != nil {
		return nil, err
	}

	if err := binary.Read(bytesBuffer, binary.LittleEndian, &data.ProtoId); err != nil {
		return nil, err
	}

	data.Data = frame[tcpfhp.getHeadLength():]

	return data, nil
}

// client 端获取解包后的数据
func (tcpfhp *TCPProtocol) ClientDecode(rawConn net.Conn) (*ProtocolData, error) {
	newPackage := ProtocolData{}

	headData := make([]byte, tcpfhp.getHeadLength())

	n, err := io.ReadFull(rawConn, headData)
	if n != tcpfhp.getHeadLength() {
		return nil, err
	}

	//数据长度
	bytesBuffer := bytes.NewBuffer(headData)
	binary.Read(bytesBuffer, binary.LittleEndian, &newPackage.ProtoId)
	binary.Read(bytesBuffer, binary.LittleEndian, &newPackage.DataLength)

	if newPackage.DataLength < 1 {
		return &newPackage, nil
	}

	data := make([]byte, newPackage.DataLength)
	dataNum, err2 := io.ReadFull(rawConn, data)

	if int32(dataNum) != newPackage.DataLength {
		return nil, fmt.Errorf("read data error, %v", err2)
	}

	newPackage.Data = data

	return &newPackage, nil
}

// output 数据编码
func (tcpfhp *TCPProtocol) EncodeWrite(protoId int32, data []byte, conn net.Conn) error {

	if conn == nil {
		return errors.New("con 为空")
	}

	pdata := ProtocolData{}
	pdata.ProtoId = protoId
	pdata.DataLength = int32(len(data))
	pdata.Data = data

	if err := binary.Write(conn, binary.LittleEndian, &pdata.ProtoId); err != nil {
		return fmt.Errorf("encodeWrite version error , %v", err)
	}

	if err := binary.Write(conn, binary.LittleEndian, &pdata.DataLength); err != nil {
		return fmt.Errorf("encodeWrite datalength error , %v", err)
	}

	if pdata.DataLength > 0 {
		if err := binary.Write(conn, binary.LittleEndian, &pdata.Data); err != nil {
			return fmt.Errorf("encodeWrite data error , %v", err)
		}
	}

	return nil
}

func (tcpfhp *TCPProtocol) Encode(c gnet.Conn, buf []byte) ([]byte, error) {
	return buf, nil
}

// 数据编码
func (tcpfhp *TCPProtocol) EncodeData(protoId int32, data []byte) ([]byte, error) {
	pdata := ProtocolData{}
	pdata.ProtoId = protoId
	pdata.DataLength = int32(len(data)) + DefaultHeadLength
	pdata.Data = data

	result := make([]byte, 0)

	buffer := bytes.NewBuffer(result)

	if err := binary.Write(buffer, binary.LittleEndian, &pdata.DataLength); err != nil {
		return nil, fmt.Errorf("encode datalength error , %v", err)
	}

	if err := binary.Write(buffer, binary.LittleEndian, &pdata.ProtoId); err != nil {
		return nil, fmt.Errorf("encode version error , %v", err)
	}

	if pdata.DataLength > 0 {
		if err := binary.Write(buffer, binary.LittleEndian, &pdata.Data); err != nil {
			return nil, fmt.Errorf("encode data error , %v", err)
		}
	}

	return buffer.Bytes(), nil
}
