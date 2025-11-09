package handlers

import (
	"myproject/database"
	"net/http"
	"text/template"
)

var result = template.Must(template.ParseFiles("frontend/result.html"))

func VerifyHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	inputOTP := r.FormValue("otp")

	database.OtpStore.RLock()
	savedOTP, exists := database.OtpStore.Data[email]
	database.OtpStore.RUnlock()

	if !exists {
		result.Execute(w, "No OTP found for this email.")
		return
	}

	if inputOTP == savedOTP {
		result.Execute(w, "OTP Verified Successfully!")
		// Delete OTP after successful verification
		database.OtpStore.Lock()
		delete(database.OtpStore.Data, email)
		database.OtpStore.Unlock()
	} else {
		result.Execute(w, "Incorrect OTP, try again!")
	}
}
