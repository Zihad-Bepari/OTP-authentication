package database

import "sync"

type OTPEntry struct {
	Email string
	OTP   string
}

var OtpStore = struct {
	sync.RWMutex
	Data map[string]string
}{Data: make(map[string]string)}
