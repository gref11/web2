package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"web3/internal/cookie"
	"web3/internal/models"
	"web3/internal/validation"
	"web3/scripts"
)

func (h *Handler) UserPageHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	funcMap := template.FuncMap{
		"in": scripts.In,
	}

	var savedData models.RegistrationData

	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		savedData.AuthorizationError = "Для просмотра пользователя необходимо авторизоваться"
	} else {

		savedData, ok = cookie.LoadUpdatedFormData(r)

		if !ok {
			userData, err := h.Storage.GetUserByID(userID)
			if err != nil {
				http.Error(w, fmt.Sprintf("Internal server error: %v", err), http.StatusInternalServerError)
				return
			}
			savedData.Name = userData.Name
			savedData.Tel = userData.Tel
			savedData.Email = userData.Email
			savedData.Sex = userData.Sex
			savedData.Bio = userData.Bio
			savedData.Langs = userData.Langs
		}
	}

	tmpl := template.Must(template.New("personal_page.html").Funcs(funcMap).ParseFiles("static/src/template/personal_page.html"))
	tmpl.Execute(w, savedData)
}

func (h *Handler) UserUpdateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, ok := ctx.Value("user_id").(string)

	if !ok {
		cookie.SetCookie(w, "auth_error", "Для просмотра пользователя необходимо авторизоваться", cookie.CookieStoragePeriod)
		http.Redirect(w, r, "http://u69196.kubsu-dev.ru/", http.StatusSeeOther)
	}
	cookie.ClearErrorCookies(w)

	if r.Method != "POST" {
		http.Redirect(w, r, "http://u69196.kubsu-dev.ru/", http.StatusSeeOther)
		return
	}

	userData, err := models.UserDataParse(r)
	if err != nil {
		http.Error(w, "Ошибка декодирования", http.StatusBadRequest)
		return
	}

	cookie.SetCookie(w, "updated_name_value", userData.Name, cookie.CookieStoragePeriod)
	cookie.SetCookie(w, "updated_tel_value", userData.Tel, cookie.CookieStoragePeriod)
	cookie.SetCookie(w, "updated_email_value", userData.Email, cookie.CookieStoragePeriod)
	cookie.SetCookie(w, "updated_sex_value", userData.Sex, cookie.CookieStoragePeriod)
	cookie.SetCookie(w, "updated_langs_value", strings.Join(userData.Langs, ","), cookie.CookieStoragePeriod)
	userData.Bio = scripts.ConvertToMnemonic(userData.Bio)
	cookie.SetCookie(w, "updated_bio_value", userData.Bio, cookie.CookieStoragePeriod)

	if validation_errors, err := validation.ValidateUserData(userData); err != nil {
		for cookie_name, cookie_value := range validation_errors {
			cookie.SetCookie(w, cookie_name, cookie_value, cookie.CookieStoragePeriod)
		}

		http.Redirect(w, r, "http://u69196.kubsu-dev.ru/", http.StatusSeeOther)
		return
	}

	if err := h.Storage.UpdateUserData(userID, userData); err != nil {
		http.Error(w, fmt.Sprintf("Internal server error: %v", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "http://u69196.kubsu-dev.ru/", http.StatusSeeOther)
}
