package store

import (
	"fmt"
	"math/rand"
	"time"

	"gopkg.in/gomail.v2"
)

func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func SendOTP(email, otp string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "zihadbepari420@gmail.com") // Sender email
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Your OTP Code")
	m.SetBody("text/plain", "Your OTP is: "+otp)

	d := gomail.NewDialer("smtp.gmail.com", 587, "", "")

	return d.DialAndSend(m)
}
