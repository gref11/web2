package cookie

import (
	"net/http"
	"net/url"
	"strings"

	"web3/internal/models"
)

var CookieStoragePeriod = 365 * 24 * 60 * 60

func ClearErrorCookies(w http.ResponseWriter) {
	cookiesToClear := []string{
		"name_error", "email_error", "tel_error",
		"name_value", "email_value", "tel_value",
	}

	for _, name := range cookiesToClear {
		SetCookie(w, name, "", -1)
	}
}

func SetCookie(w http.ResponseWriter, name, value string, maxAge int) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    url.QueryEscape(value),
		MaxAge:   maxAge,
		Path:     "/",
		HttpOnly: true,
	})
}

func LoadFormData(r *http.Request, data *models.RegistrationData) {
	for _, cookie := range r.Cookies() {
		switch cookie.Name {
		case "name_error":
			data.Errors["Name"], _ = url.QueryUnescape(cookie.Value)
		case "email_error":
			data.Errors["Email"], _ = url.QueryUnescape(cookie.Value)
		case "tel_error":
			data.Errors["Tel"], _ = url.QueryUnescape(cookie.Value)
		case "name_value":
			data.Name, _ = url.QueryUnescape(cookie.Value)
		case "email_value":
			data.Email, _ = url.QueryUnescape(cookie.Value)
		case "tel_value":
			data.Tel, _ = url.QueryUnescape(cookie.Value)
		case "sex_value":
			data.Sex = cookie.Value
		case "date_value":
			data.Date = cookie.Value
		case "langs_value":
			langs, _ := url.QueryUnescape(cookie.Value)
			data.Langs = strings.Split(langs, ",")
		case "bio_value":
			data.Bio, _ = url.QueryUnescape(cookie.Value)
		}
	}
}
