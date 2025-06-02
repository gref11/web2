package models

type RegistrationData struct {
	Name   string
	Email  string
	Tel    string
	Date   string
	Sex    string
	Langs  []string
	Bio    string
	Errors map[string]string
}
