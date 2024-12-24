package common

import (
	"regexp"
)

func ValidateEmail(email string) error {
	if email == "" {
		return emailEmpty
	}
	if len(email) > 20 {
		return emailTooLong
	}
	if !regexp.MustCompile(emailRegex).MatchString(email) {
		return emailMismatch
	}
	return nil
}
func ValidatePassword(password string) error {
	if password == "" {
		return passwordEmpty
	}
	if len(password) < 8 {
		return passwordShort
	}
	if !regexp.MustCompile(passwordRegex).MatchString(password) {
		return passwordMismatch
	}
	return nil
}
func ValidateUsername(username string) error {

	if username == "" {
		return userNameIsEmpty
	}
}
