package radius

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"errors"
	"net"
	"strconv"
)

// Attribute is a wire encoded RADIUS attribute.
type Attribute []byte

// String returns the given attribute as a string.
func String(a Attribute) string {
	return string(a)
}

// NewString returns a new Attribute from the given string. An error is returned
// if the string length is greater than 253.
func NewString(s string) (Attribute, error) {
	if len(s) > 253 {
		return nil, errors.New("string too long")
	}
	return Attribute(s), nil
}

// Bytes returns the given Attribute as a byte slice.
func Bytes(a Attribute) []byte {
	b := make([]byte, len(a))
	copy(b, []byte(a))
	return b
}

// NewBytes returns a new Attribute from the given byte slice. An error is
// returned if the slice is longer than 253.

func NewBytes(b []byte) (Attribute, error) {
	if len(b) > 253 {
		return nil, errors.New("value too long")
	}
	a := make(Attribute, len(b))
	copy(a, Attribute(b))
	return a, nil
}

// Integer64 returns the given attribute as an integer. An error is returned if
// the attribute is not 8 bytes long.
func Integer64(a Attribute) (uint64, error) {
	if len(a) != 8 {
		return 0, errors.New("invalid length")
	}
	return binary.BigEndian.Uint64(a), nil
}

// NewInteger64 creates a new Attribute from the given integer value.
func NewInteger64(i uint64) Attribute {
	v := make([]byte, 8)
	binary.BigEndian.PutUint64(v, i)
	return Attribute(v)
}

func UserPassword(a Attribute, secret, requestAuthenticator []byte) ([]byte, error) {
	if len(a) < 16 || len(a) > 128 {
		return nil, errors.New("invalid attribute length (" + strconv.Itoa(len(a)) + ")")
	}
	if len(secret) == 0 {
		return nil, errors.New("empty secret")
	}
	if len(requestAuthenticator) != 16 {
		return nil, errors.New("invalid requestAuthenticator length (" + strconv.Itoa(len(requestAuthenticator)) + ")")
	}

	dec := make([]byte, 0, len(a))

	hash := md5.New()
	hash.Write(secret)
	hash.Write(requestAuthenticator)
	dec = hash.Sum(dec)

	for i, b := range a[:16] {
		dec[i] ^= b
	}

	for i := 16; i < len(a); i += 16 {
		hash.Reset()
		hash.Write(secret)
		hash.Write(a[i-16 : i])
		dec = hash.Sum(dec)

		for j, b := range a[i : i+16] {
			dec[i+j] ^= b
		}
	}

	if i := bytes.IndexByte(dec, 0); i > -1 {
		return dec[:i], nil
	}
	return dec, nil
}
func IPAddr(a Attribute) (net.IP, error) {
	if len(a) != net.IPv4len {
		return nil, errors.New("invalid length")
	}
	b := make([]byte, net.IPv4len)
	copy(b, []byte(a))
	return b, nil
}

func Integer(a Attribute) (uint32, error) {
	if len(a) != 4 {
		return 0, errors.New("invalid length")
	}
	return binary.BigEndian.Uint32(a), nil
}

// NewInteger creates a new Attribute from the given integer value.
func NewInteger(i uint32) Attribute {
	v := make([]byte, 4)
	binary.BigEndian.PutUint32(v, i)
	return Attribute(v)
}
