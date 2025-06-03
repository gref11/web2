package cookie

import (
	"net/http"
	"net/url"
	"strings"

	"web3/internal/models"
)

var CookieStoragePeriod = 60 * 60

func ClearErrorCookies(w http.ResponseWriter) {
	cookiesToClear := []string{
		"name_error", "email_error", "tel_error", "auth_error",
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

func LoadFormData(r *http.Request) (models.RegistrationData, bool) {
	data := models.RegistrationData{
		Errors: make(map[string]string),
	}
	ok := false
	for _, cookie := range r.Cookies() {
		switch cookie.Name {
		case "name_error":
			data.Errors["Name"], _ = url.QueryUnescape(cookie.Value)
			ok = true
		case "email_error":
			data.Errors["Email"], _ = url.QueryUnescape(cookie.Value)
			ok = true
		case "tel_error":
			data.Errors["Tel"], _ = url.QueryUnescape(cookie.Value)
			ok = true
		case "name_value":
			data.Name, _ = url.QueryUnescape(cookie.Value)
		case "email_value":
			data.Email, _ = url.QueryUnescape(cookie.Value)
		case "tel_value":
			data.Tel, _ = url.QueryUnescape(cookie.Value)
		case "sex_value":
			data.Sex = cookie.Value
		case "langs_value":
			langs, _ := url.QueryUnescape(cookie.Value)
			data.Langs = strings.Split(langs, ",")
		case "bio_value":
			data.Bio, _ = url.QueryUnescape(cookie.Value)
		}
	}

	return data, ok
}

func LoadUpdatedFormData(r *http.Request) (models.RegistrationData, bool) {
	data := models.RegistrationData{
		Errors: make(map[string]string),
	}
	ok := false
	for _, cookie := range r.Cookies() {
		switch cookie.Name {
		case "updated_name_error":
			data.Errors["Name"], _ = url.QueryUnescape(cookie.Value)
			ok = true
		case "updated_email_error":
			data.Errors["Email"], _ = url.QueryUnescape(cookie.Value)
			ok = true
		case "updated_tel_error":
			data.Errors["Tel"], _ = url.QueryUnescape(cookie.Value)
			ok = true
		case "updated_name_value":
			data.Name, _ = url.QueryUnescape(cookie.Value)
		case "updated_email_value":
			data.Email, _ = url.QueryUnescape(cookie.Value)
		case "updated_tel_value":
			data.Tel, _ = url.QueryUnescape(cookie.Value)
		case "updated_sex_value":
			data.Sex = cookie.Value
		case "updated_langs_value":
			langs, _ := url.QueryUnescape(cookie.Value)
			data.Langs = strings.Split(langs, ",")
		case "updated_bio_value":
			data.Bio, _ = url.QueryUnescape(cookie.Value)
		}
	}

	return data, ok
}
