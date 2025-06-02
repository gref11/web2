package validation

import (
	"errors"

	"web3/internal/models"
	"web3/internal/regexps"
)

func ValidateUserData(userData models.UserData) (map[string]string, error) {
	hasError := false
	validation_errors := make(map[string]string, 0)

	if !regexps.NameRegex.MatchString(userData.Name) {
		validation_errors["name_error"] = "Имя может содержать только символы латиницы, кириллицы и пробелы, и не может превышать 150 символов"
		hasError = true
	}

	if !regexps.TelRegex.MatchString(userData.Tel) {
		validation_errors["tel_error"] = "Телефон должен быть формата +01234567890"
		hasError = true
	}

	if !regexps.EmailRegex.MatchString(userData.Email) {
		validation_errors["email_error"] = "Email должен быть формата adress@domen.tld"
		hasError = true
	}

	if hasError {
		return validation_errors, errors.New("User data validation error")
	}

	return validation_errors, nil
}
