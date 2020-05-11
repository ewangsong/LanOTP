package otp

import (
	"time"
)

//TOTP 结构体
type TOTP struct {
	Time          time.Time // 时间计数器
	Period        uint8     // 时间窗口通常30
	WindowBack    uint8     // How many steps HOTP will go backwards to validate a token
	WindowForward uint8     // How many steps HOTP will go forward to validate a token
	BasicOtp
}

func (t *TOTP) setDefaults() {
	if t.Time.IsZero() {
		t.Time = time.Now()
	}
	if t.Length == 0 {
		t.Length = DefaultLength
	}

	if t.Period == 0 {
		t.Period = DefaultPeriod
	}
	if t.WindowBack == 0 {
		t.WindowBack = DefaultWindowBack
	}
	if t.WindowForward == 0 {
		t.WindowForward = DefaultWindowForward
	}
}

//Get totp token
func (t *TOTP) Get() string {
	t.setDefaults()
	t.BasicOtp.Counter = uint64(t.Time.Unix() / int64(t.Period))
	return t.GetOtpToken()
}

// Now is a fluent interface to set the TOTP generator's time to the current date/time
func (t *TOTP) Now() *TOTP {
	t.Time = time.Now()
	return t
}

// Verify a token with the current settings, including the WindowBack and WindowForward
func (t TOTP) Verify(token string) bool {
	t.setDefaults()
	givenTime := t.Time
	for i := int(t.WindowBack) * -1; i <= int(t.WindowForward); i++ {
		t.Time = givenTime.Add(time.Second * time.Duration(int(t.Period)*i))
		if t.Get() == token {
			return true
		}
	}
	return false
}
