package session

import (
	Users "authentication/api/pkg/types/users"
	ORM "authentication/api/src/system/db"
	"authentication/api/src/system/jwt"
	Passwords "authentication/api/src/system/passwords"

	"encoding/json"
	"log"
	"net/http"
	"os"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.FormValue("email")
	password := r.FormValue("password")

	if len(email) < 1 || len(password) < 1 {
		http.Error(w, "Email and password are required.", http.StatusUnauthorized)
		return
	}

	user := Users.User{Email: email}
	err := ORM.FindBy(db, &user)
	if err != nil || user.Id < 1 {
		log.Println(err)
		http.Error(w, "Credentials do not match.", http.StatusUnauthorized)
		return
	}

	if !Passwords.IsValid(user.Password, password) {
		http.Error(w, "Credentials do not match.", http.StatusUnauthorized)
		return
	}

	token := jwt.GetToken(user.Id)
	login := LoginData{User: user, Token: token}

	http.SetCookie(w, &http.Cookie{
		Name:       os.Getenv("COOKIE_NAME"),
		Value:      token,
		Path:       "/",
		RawExpires: "0",
	})

	packet, err := json.Marshal(login)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to marshal json.", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}
