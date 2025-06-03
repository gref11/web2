package auth

import (
	"net/http"
	"time"
)

func SetAuthCookie(w http.ResponseWriter, token string, expiration time.Time) {
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    token,
		Expires:  expiration,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})
}

func GetAuthToken(r *http.Request) (string, error) {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}
