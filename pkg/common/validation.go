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
func ValidateUsername(userName string) error {

	if len(userName) > 25 {
		return userNameIsTooLong
	}
	if len(userName) == 0 {
		return userNameIsEmpty
	}
	return nil
}
func Mobile(number string) error {
	if number == "" {
		return mobileIsEmpty
	}
	if len(number) > 11 {
		return mobileIsTooLong
	}
	if !regexp.MustCompile(mobileNumberRegex).MatchString(number) {
		return mobileIsMismatch
	}
	return nil
}
func Firstname(name string) error {
	if name == "" {
		return firstNameIsEmpty
	}
	if len(name) > 25 {
		return firstNameIsTooLong
	}
	return nil
}
func Lastname(name string) error {
	if name == "" {
		return lastNameIsEmpty
	}
	if len(name) > 25 {
		return lastNameIsTooLong
	}
	return nil
}
