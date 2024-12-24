package common

import (
	"errors"
	"fmt"
)

const (
	emailRegex        = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	passwordRegex     = `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`
	mobileNumberRegex = `^\+?[1-9]\d{1,14}$`
)

var (
	emailEmpty         = errors.New("email is empty")
	emailTooLong       = errors.New("email is too long")
	emailMismatch      = errors.New("invalid email")
	passwordEmpty      = errors.New("password is empty")
	passwordShort      = errors.New("password is too short")
	passwordMismatch   = errors.New(fmt.Sprintf("invalid password, reason %s", "At least one lowercase letter, At least one uppercase letter,At least one digit,  At least one special character, At least 8 characters long, allowing letters, digits, and specified special characters."))
	userNameIsEmpty    = errors.New("username is empty")
	userNameIsTooLong  = errors.New("username is too long")
	firstNameIsEmpty   = errors.New("first name is empty")
	firstNameIsTooLong = errors.New("first name is too long")
	lastNameIsEmpty    = errors.New("last name is empty")
	lastNameIsTooLong  = errors.New("last name is too long")
	mobileIsEmpty      = errors.New("mobile is empty")
	mobileIsTooLong    = errors.New("mobile is too long")
	mobileIsMismatch   = errors.New("mobile is mismatch")
)
