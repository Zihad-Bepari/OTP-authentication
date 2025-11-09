package handlers

import (
	"fmt"
	"myproject/database"
	"myproject/store"
	"net/http"
	"text/template"
)

var index = template.Must(template.ParseFiles("frontend/index.html"))
var verification = template.Must(template.ParseFiles("frontend/verify.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		index.Execute(w, nil)
		return
	}

	email := r.FormValue("email")
	otp := store.GenerateOTP()

	err := store.SendOTP(email, otp)
	if err != nil {
		fmt.Fprintf(w, "Failed to send OTP: %v", err)
		return
	}

	database.OtpStore.Lock()
	database.OtpStore.Data[email] = otp
	database.OtpStore.Unlock()

	verification.Execute(w, email)
}
