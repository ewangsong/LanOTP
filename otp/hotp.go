package otp

//HOTP 结构体
type HOTP struct {
	BasicOtp
}

//Get 获取htop token
func (h *HOTP) HotpGet() string {
	return h.GetOtpToken()
}
