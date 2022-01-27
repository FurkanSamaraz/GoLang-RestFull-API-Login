package main

import (
	"fmt"
	"net/http"

	help "github.com/FurkanSamaraz/IsEmpty"
	_ "github.com/lib/pq"
)

var uName, pwd, pwdConfirm string

func bosSignup(w http.ResponseWriter, r *http.Request) {
	uNameCheck := help.IsEmpty(uName)
	pwdCheck := help.IsEmpty(pwd)
	pwdConfirmCheck := help.IsEmpty(pwdConfirm)

	if uNameCheck || pwdCheck || pwdConfirmCheck {
		fmt.Fprintf(w, "Error BOŞ! \n")
	}
}

func bosLogin(w http.ResponseWriter, r *http.Request) {
	uNameCheck := help.IsEmpty(uName)
	pwdCheck := help.IsEmpty(pwd)

	if uNameCheck || pwdCheck {
		fmt.Fprintf(w, "Error BOŞ! \n")
	}
}
func signup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	uName = r.FormValue("username")
	pwd = r.FormValue("password")
	pwdConfirm = r.FormValue("passworConfirm")

	bosSignup(w, r)

	if pwd == pwdConfirm {
		fmt.Fprintf(w, "Şifre Doğru")
	} else {
		fmt.Fprintf(w, "Şifre Yalnış")
	}
}
func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	uName = r.FormValue("username")
	pwd = r.FormValue("password")
	dbname := "furkan"
	dbpassword := "123"

	bosLogin(w, r)

	if dbpassword == pwd && dbname == uName {
		fmt.Fprintf(w, "Giriş Başarılı")
	} else {
		fmt.Fprintf(w, "Giriş Başarısız")
	}

}
func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/login", login)
	http.ListenAndServe(":8080", mux)
}
