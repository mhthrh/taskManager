package common

import (
	"errors"
	"testing"
)

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		email string
		err   error
	}{
		{
			email: "",
			err:   emailEmpty,
		}, {
			email: "aieyfguryfbuwhbfwihfbuyfblhqjndkje@b.com",
			err:   emailTooLong,
		}, {
			email: "kiri.com",
			err:   emailMismatch,
		}, {
			email: "kiri.gmail.com",
			err:   emailMismatch,
		}, {
			email: "mhthrhr@gmail.com",
			err:   nil,
		},
	}

	for _, test := range tests {
		if err := ValidateEmail(test.email); !errors.Is(err, test.err) {
			t.Errorf("ValidateEmail(%q): want %v but get %v", test.email, err, test.err)
		}
	}
}

func TestValidateUsername(t *testing.T) {
	tests := []struct {
		name     string
		userName string
		err      error
	}{
		{
			name:     "username-test-1",
			userName: "",
			err:      userNameIsEmpty,
		}, {
			name:     "username-test-2",
			userName: "kjfenwkjfnjfwkjfnkjenj32843800000",
			err:      userNameIsTooLong,
		}, {
			name:     "username-test-3",
			userName: "mohsen",
			err:      nil,
		},
	}

	for _, test := range tests {
		if err := ValidateUsername(test.userName); !errors.Is(err, test.err) {
			t.Errorf("ValidateUsername(%q): want %v but get %v", test.userName, test.err, err)
		}
	}
}

func TestMobile(t *testing.T) {
	tests := []struct {
		name   string
		mobile string
		err    error
	}{
		{
			name:   "mobile-test-1",
			mobile: "",
			err:    mobileIsEmpty,
		}, {
			name:   "mobile-test-2",
			mobile: "83175y41989581984985194",
			err:    mobileIsTooLong,
		}, {
			name:   "mobile-test-3",
			mobile: "16047277989",
			err:    nil,
		},
	}
	for _, test := range tests {
		if err := Mobile(test.mobile); !errors.Is(err, test.err) {
			t.Errorf("Mobile(%q): want %v but get %v", test.mobile, test.err, err)
		}
	}
}
func TestFirstname(t *testing.T) {
	tests := []struct {
		testName string
		name     string
		err      error
	}{
		{
			testName: "firstNameTest-1",
			name:     "",
			err:      firstNameIsEmpty,
		}, {
			testName: "firstNameTest-2",
			name:     "ijfnwhjbfiwhbfiwhbiewbibiebieub",
			err:      firstNameIsTooLong,
		}, {
			testName: "firstNameTest-3",
			name:     "mohsen",
			err:      nil,
		},
	}
	for _, test := range tests {
		if err := Firstname(test.name); !errors.Is(err, test.err) {
			t.Errorf("Firstname(%q): want %v but get %v", test.name, test.err, err)
		}
	}
}
func TestLastname(t *testing.T) {
	tests := []struct {
		testName string
		name     string
		err      error
	}{
		{
			testName: "lastNameTest-1",
			name:     "",
			err:      lastNameIsEmpty,
		}, {
			testName: "lastNameTest-2",
			name:     "ijfnwhjbfiwhbfiwhbiewbibieub",
			err:      lastNameIsTooLong,
		}, {
			testName: "lastNameTest-3",
			name:     "mohsen",
		},
	}
	for _, test := range tests {
		if err := Lastname(test.name); !errors.Is(err, test.err) {
			t.Errorf("Lastname(%q): want %v but get %v", test.name, test.err, err)
		}
	}
}
