package models

type personalData struct {
	Name  string   `json:"fullName"`
	Email string   `json:"email"`
	Tel   string   `json:"phone"`
	Date  string   `json:"date"`
	Sex   string   `json:"sex"`
	Langs []string `json:"lang"`
	Bio   string   `json:"bio"`
}
