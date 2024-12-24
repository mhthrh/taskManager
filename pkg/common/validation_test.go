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
