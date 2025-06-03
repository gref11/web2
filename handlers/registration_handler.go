package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"web3/internal/auth"
	"web3/internal/cookie"
	"web3/internal/hash"
	"web3/internal/models"
	"web3/internal/validation"
	"web3/scripts"
)

func (h *Handler) RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"in": scripts.In,
	}

	savedData, _ := cookie.LoadFormData(r)

	tmpl := template.Must(template.New("registration.html").Funcs(funcMap).ParseFiles("static/src/template/registration.html"))
	tmpl.Execute(w, savedData)
}

func (h *Handler) UserAddHandler(w http.ResponseWriter, r *http.Request) {
	scripts.EnableCORS(w)
	cookie.ClearErrorCookies(w)

	if r.Method != "POST" {
		http.Redirect(w, r, "http://u69196.kubsu-dev.ru/signup", http.StatusSeeOther)
		return
	}

	userData, err := models.UserDataParse(r)
	if err != nil {
		http.Error(w, "Ошибка декодирования", http.StatusBadRequest)
		return
	}

	cookie.SetCookie(w, "name_value", userData.Name, cookie.CookieStoragePeriod)
	cookie.SetCookie(w, "tel_value", userData.Tel, cookie.CookieStoragePeriod)
	cookie.SetCookie(w, "email_value", userData.Email, cookie.CookieStoragePeriod)
	cookie.SetCookie(w, "sex_value", userData.Sex, cookie.CookieStoragePeriod)
	cookie.SetCookie(w, "langs_value", strings.Join(userData.Langs, ","), cookie.CookieStoragePeriod)
	userData.Bio = scripts.ConvertToMnemonic(userData.Bio)
	cookie.SetCookie(w, "bio_value", userData.Bio, cookie.CookieStoragePeriod)

	if validation_errors, err := validation.ValidateUserData(userData); err != nil {
		for cookie_name, cookie_value := range validation_errors {
			cookie.SetCookie(w, cookie_name, cookie_value, cookie.CookieStoragePeriod)
		}

		http.Redirect(w, r, "http://u69196.kubsu-dev.ru/signup", http.StatusSeeOther)
		return
	}

	userID, err := h.Storage.InsertUserData(userData)

	if err != nil {
		http.Error(w, fmt.Sprintf("Internal server error: %v", err), http.StatusInternalServerError)
		return
	}

	login, pwd := auth.Generate(userID)
	pwdHash, err := hash.HashPassword(pwd)
	if err != nil {
		http.Error(w, fmt.Sprintf("Internal server error: %v", err), http.StatusInternalServerError)
		return
	}

	if err := h.Storage.NewUserAuth(userID, login, pwdHash); err != nil {
		http.Error(w, fmt.Sprintf("Internal server error: %v", err), http.StatusInternalServerError)
		return
	}

	cookie.SetCookie(w, "login", login, cookie.CookieStoragePeriod)
	cookie.SetCookie(w, "password", pwd, cookie.CookieStoragePeriod)

	http.Redirect(w, r, "http://u69196.kubsu-dev.ru/success_signup", http.StatusSeeOther)
}
