package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"time"

	"web3/internal/auth"
	"web3/internal/cookie"
	"web3/internal/hash"

	"github.com/golang-jwt/jwt/v5"
)

func (h *Handler) AfterRegistrationPageHandler(w http.ResponseWriter, r *http.Request) {
	login, err := r.Cookie("login")
	if err != nil {
		http.Error(w, fmt.Sprintf("Internal server error: %v", err), http.StatusInternalServerError)
		return
	}
	pwd, err := r.Cookie("password")
	if err != nil {
		http.Error(w, fmt.Sprintf("Internal server error: %v", err), http.StatusInternalServerError)
		return
	}

	savedData := struct {
		Login    string
		Password string
	}{
		Login:    login.Value,
		Password: pwd.Value,
	}

	cookie.SetCookie(w, "login", "", cookie.CookieStoragePeriod)
	cookie.SetCookie(w, "password", "", cookie.CookieStoragePeriod)

	tmpl := template.Must(template.New("after_registration.html").ParseFiles("static/src/template/after_registration.html"))
	tmpl.Execute(w, savedData)
}

func (h *Handler) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	auth.SetAuthCookie(w, "", time.Now().Add(-1*time.Hour))
	var loginValue string
	login, err := r.Cookie("login")
	if err == nil {
		loginValue, _ = url.QueryUnescape(login.Value)
	}
	var passwordValue string
	pwd, err := r.Cookie("password")
	if err == nil {
		passwordValue, _ = url.QueryUnescape(pwd.Value)
	}
	var errorValue string
	auth_error, err := r.Cookie("auth_error")
	if err == nil {
		errorValue, _ = url.QueryUnescape(auth_error.Value)
	}

	savedData := struct {
		Login    string
		Password string
		Error    string
	}{
		Login:    loginValue,
		Password: passwordValue,
		Error:    errorValue,
	}

	cookie.SetCookie(w, "login", "", cookie.CookieStoragePeriod)
	cookie.SetCookie(w, "password", "", cookie.CookieStoragePeriod)
	cookie.SetCookie(w, "auth_error", "", cookie.CookieStoragePeriod)

	tmpl := template.Must(template.New("login_page.html").ParseFiles("static/src/template/login_page.html"))
	tmpl.Execute(w, savedData)
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var user auth.User
	user.Login = r.FormValue("login")
	user.Password = r.FormValue("password")

	userID, pwd_hash, err := h.Storage.GetPasswordHash(user.Login)

	if err != nil || !hash.CheckPassword(user.Password, pwd_hash) {
		cookie.SetCookie(w, "login", user.Login, cookie.CookieStoragePeriod)
		cookie.SetCookie(w, "password", user.Password, cookie.CookieStoragePeriod)
		cookie.SetCookie(w, "auth_error", "Ошибка авторизации: проверьте логин или пароль", cookie.CookieStoragePeriod)
		http.Redirect(w, r, "http://u69196.kubsu-dev.ru/login", http.StatusSeeOther)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &auth.Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(auth.JwtSecret)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	auth.SetAuthCookie(w, tokenString, time.Now().Add(24*time.Hour))

	http.Redirect(w, r, "http://u69196.kubsu-dev.ru/", http.StatusSeeOther)
}
