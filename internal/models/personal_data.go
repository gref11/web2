package models

import (
	"net/http"
)

type UserData struct {
	Name  string   `json:"fullName"`
	Email string   `json:"email"`
	Tel   string   `json:"phone"`
	Date  string   `json:"date"`
	Sex   string   `json:"sex"`
	Langs []string `json:"lang"`
	Bio   string   `json:"bio"`
}

func UserDataParse(r *http.Request) (UserData, error) {
	err := r.ParseForm()
	if err != nil {
		return UserData{}, err
	}

	var uData UserData
	uData.Name = r.FormValue("fullName")
	uData.Email = r.FormValue("email")
	uData.Tel = r.FormValue("phone")
	uData.Date = r.FormValue("date")
	uData.Sex = r.FormValue("sex")
	uData.Langs = r.Form["lang"]
	uData.Bio = r.FormValue("bio")

	return uData, nil
}
