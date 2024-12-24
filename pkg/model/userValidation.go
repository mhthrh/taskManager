package model

import (
	"errors"
	"github.com/mhthrh/taskManager/pkg/common"
)

func (u *User) Validate() error {
	if len(u.Username) > 25 {
		return common.userNameIsTooLong
	}
	if len(u.Username) == 0 {
		return errors.New("username is empty")
	}
	return nil
}
