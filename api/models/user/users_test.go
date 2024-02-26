package user

import "testing"

func TestPhoneSignUp(t *testing.T) {
	u := NewUser()
	u.PhoneSignUp("13800138000", "86")
}

func TestEmailSignUp(t *testing.T) {
	u := NewUser()
	u.EmailSignUp("abc@gmail.com", "123456")
}
