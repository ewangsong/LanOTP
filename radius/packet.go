package radius

import (
	"crypto/md5"
	"encoding/binary"
	"errors"
)

//定义一个最大值的数据包
const MaxPacketLength = 4095

//一个数据包结构
type Packet struct {
	Code          Code
	Identifier    byte
	Authenticator [16]byte
	Secret        []byte
	Attributes
}

//解析数据包把属性域提取出来
func Parse(b, secret []byte) (*Packet, error) {
	if len(b) < 20 {
		return nil, errors.New("radius: packet not at least 20 bytes long")
	}

	length := int(binary.BigEndian.Uint16(b[2:4]))
	b = b[0:length]
	if length < 20 || length > MaxPacketLength || len(b) != length {
		return nil, errors.New("radius: invalid packet length")
	}

	attrs, err := ParseAttributes(b[20:])
	if err != nil {
		return nil, err
	}

	packet := &Packet{
		Code:       Code(b[0]),
		Identifier: b[1],
		Secret:     secret,
		Attributes: attrs,
	}
	copy(packet.Authenticator[:], b[4:20])
	return packet, nil
}

//返回一个新的数据包
func (p *Packet) Response(code Code) *Packet {
	q := &Packet{
		Code:       code,
		Identifier: p.Identifier,
		Secret:     p.Secret,
		Attributes: make(Attributes),
	}
	copy(q.Authenticator[:], p.Authenticator[:])
	return q
}

//编辑数据包

func (p *Packet) Encode() ([]byte, error) {
	attributesSize := p.Attributes.wireSize()
	if attributesSize == -1 {
		return nil, errors.New("invalid packet attribute length")
	}
	size := 20 + attributesSize //整个数据包的大小
	if size > MaxPacketLength {
		return nil, errors.New("encoded packet is too long")
	}

	b := make([]byte, size)
	b[0] = byte(p.Code)
	b[1] = byte(p.Identifier)
	binary.BigEndian.PutUint16(b[2:4], uint16(size))
	p.Attributes.encodeTo(b[20:])

	switch p.Code {
	case CodeAccessRequest:
	case CodeAccessAccept, CodeAccessReject, CodeAccountingRequest, CodeAccountingResponse, CodeAccessChallenge, CodeDisconnectRequest, CodeDisconnectACK, CodeDisconnectNAK, CodeCoARequest, CodeCoAACK, CodeCoANAK:
		hash := md5.New()
		hash.Write(b[:4])
		switch p.Code {
		case CodeAccountingRequest, CodeDisconnectRequest, CodeCoARequest:
			var nul [16]byte
			hash.Write(nul[:])
		default:
			hash.Write(p.Authenticator[:])
		}
		hash.Write(b[20:])
		hash.Write(p.Secret)
		hash.Sum(b[4:4:20])
	default:
		return nil, errors.New("radius: unknown Packet Code")
	}

	return b, nil
}
