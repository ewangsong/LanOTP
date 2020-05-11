package otp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// BasicOtp 初始结构体
type BasicOtp struct {
	Secret  string //密钥
	Length  uint8  //otp长度
	Counter uint64 //counter移动元素计数器 计数器
}

//GenerateRandomSecret 生成密钥
func GenerateRandomSecret() string {
	//base32字符集
	randnum := "234567ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var bytes = make([]byte, 32)
	rand.Seed(time.Now().UnixNano())
	for i := range bytes {
		bytes[i] = randnum[rand.Intn(32)]
	}
	return (string(bytes))
}

//counterToBytes 计数器转换成bytes
func counterToBytes(counter uint64) (text []byte) {
	text = make([]byte, 8)
	for i := 7; i >= 0; i-- {
		text[i] = byte(counter & 0xff)
		counter = counter >> 8
	}
	return
}

//算法流程
func hmacSHA1AndTruncate(key, text []byte) int {
	H := hmac.New(sha1.New, key)
	H.Write([]byte(text))
	hash := H.Sum(nil)
	//返回计算int值
	offset := int(hash[len(hash)-1] & 0xf)
	return ((int(hash[offset]) & 0x7f) << 24) |
		((int(hash[offset+1] & 0xff)) << 16) |
		((int(hash[offset+2] & 0xff)) << 8) |
		(int(hash[offset+3]) & 0xff)

}

//GetOtpToken 获取token值
func (o *BasicOtp) GetOtpToken() string {
	//base32
	secretBytes, _ := base32.StdEncoding.DecodeString(o.Secret)
	text := counterToBytes(o.Counter)
	//获取十进制
	binary := hmacSHA1AndTruncate(secretBytes, text)
	otpint := int64(binary) % int64(math.Pow10(int(o.Length)))
	otpstring := fmt.Sprintf(fmt.Sprintf("%%0%dd", o.Length), otpint)
	return otpstring
}
