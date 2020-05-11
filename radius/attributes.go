package radius

import (
	"errors"
	"net"
)

var ErrNoAttribute = errors.New("radius: attribute not found")

//定义Type为int
type Type int

// TypeInvalid是一个可用于表示无效RADIUS的Type属性类型。
const TypeInvalid Type = -1

// Attributes是RADIUS属性类型的map，用于切片属性。
type Attributes map[Type][]Attribute

//解析属性包
func ParseAttributes(b []byte) (Attributes, error) {
	attrs := make(map[Type][]Attribute)

	for len(b) > 0 {
		if len(b) < 2 {
			return nil, errors.New("short buffer")
		}
		length := int(b[1])
		if length > len(b) || length < 2 || length > 255 {
			return nil, errors.New("invalid attribute length")
		}

		typ := Type(b[0])
		var value Attribute
		if length > 2 {
			value = make(Attribute, length-2)
			copy(value, b[2:])
		}
		attrs[typ] = append(attrs[typ], value)

		b = b[length:]
	}

	return attrs, nil
}

//查询是否含有某type属性
func (a Attributes) Lookup(key Type) (Attribute, bool) {
	m := a[key]
	if len(m) == 0 {
		return nil, false
	}
	return m[0], true
}

//get某属性的值
func (a Attributes) Get(key Type) Attribute {
	attr, _ := a.Lookup(key)
	return attr
}

//添加一个属性
func (a Attributes) Add(key Type, value Attribute) {
	a[key] = append(a[key], value)
}

//删除一个属性
func (a Attributes) Del(key Type) {
	delete(a, key)
}

// 删除同一属性并添加值
func (a Attributes) Set(key Type, value Attribute) {
	a[key] = append(a[key][:0], value)
}

//属性代码值
func (a Attributes) encodeTo(b []byte) {
	types := make([]int, 0, len(a))
	for typ := range a {
		if typ >= 1 && typ <= 255 {
			types = append(types, int(typ))
		}
	}

	for _, typ := range types {
		for _, attr := range a[Type(typ)] {
			if len(attr) > 255 {
				continue
			}
			size := 1 + 1 + len(attr)
			b[0] = byte(typ)
			b[1] = byte(size)
			copy(b[2:], attr)
			b = b[size:]
		}
	}
}

//属性字段的总长度
func (a Attributes) wireSize() (bytes int) {
	for typ, attrs := range a {
		if typ < 1 || typ > 255 {
			continue
		}
		for _, attr := range attrs {
			if len(attr) > 255 {
				return -1
			}
			// type field + length field + value field
			bytes += 1 + 1 + len(attr)
		}
	}
	return
}

func UserName_GetString(p *Packet) (value string) {
	value, _ = UserName_LookupString(p)
	return
}

func UserName_LookupString(p *Packet) (value string, err error) {
	a, ok := p.Lookup(UserName_Type)
	if !ok {
		err = ErrNoAttribute
		return
	}
	value = String(a)
	return
}

func UserPassword_GetString(p *Packet) (value string) {
	value, _ = UserPassword_LookupString(p)
	return
}

func UserPassword_LookupString(p *Packet) (value string, err error) {
	a, ok := p.Lookup(UserPassword_Type)
	if !ok {
		err = ErrNoAttribute
		return
	}
	var b []byte
	b, err = UserPassword(a, p.Secret, p.Authenticator[:])
	if err == nil {
		value = string(b)
	}
	return
}

func NASIPAddress_Get(p *Packet) (value net.IP) {
	value, _ = NASIPAddress_Lookup(p)
	return
}

func NASIPAddress_Lookup(p *Packet) (value net.IP, err error) {
	a, ok := p.Lookup(NASIPAddress_Type)
	if !ok {
		err = ErrNoAttribute
		return
	}
	value, err = IPAddr(a)
	return
}

func ReplyMessage_Add(p *Packet, value []byte) (err error) {
	var a Attribute
	a, err = NewBytes(value)
	if err != nil {
		return
	}
	p.Add(ReplyMessage_Type, a)
	return
}

func ReplyMessage_AddString(p *Packet, value string) (err error) {
	var a Attribute
	a, err = NewString(value)
	if err != nil {
		return
	}
	p.Add(ReplyMessage_Type, a)
	return
}

func ReplyMessage_Lookup(p *Packet) (value []byte, err error) {
	a, ok := p.Lookup(ReplyMessage_Type)
	if !ok {
		err = ErrNoAttribute
		return
	}
	value = Bytes(a)
	return
}

func ReplyMessage_LookupString(p *Packet) (value string, err error) {
	a, ok := p.Lookup(ReplyMessage_Type)
	if !ok {
		err = ErrNoAttribute
		return
	}
	value = String(a)
	return
}

func ReplyMessage_Set(p *Packet, value []byte) (err error) {
	var a Attribute
	a, err = NewBytes(value)
	if err != nil {
		return
	}
	p.Set(ReplyMessage_Type, a)
	return
}

func ReplyMessage_SetString(p *Packet, value string) (err error) {
	var a Attribute
	a, err = NewString(value)
	if err != nil {
		return
	}
	p.Set(ReplyMessage_Type, a)
	return
}
