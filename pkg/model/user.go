package model

import (
	"github.com/google/uuid"
	"github.com/mhthrh/taskManager/pkg/common"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Username  string    `json:"username" gorm:"type:varchar(25)"`
	Password  string    `json:"password"  gorm:"type:varchar(255)"`
	Email     string    `json:"email" gorm:"type:varchar(25)"`
	Mobile    string    `json:"mobile" gorm:"type:varchar(25)"`
	FirstName string    `gorm:"type:varchar(50)"`
	LastName  string    `json:"lastName" gorm:"type:varchar(50)"`
	createdAt time.Time `gorm:"type:timestamp"`
	updatedAt time.Time `gorm:"type:timestamp"`
	deletedAt time.Time `gorm:"type:timestamp"`
	Status    bool      `json:"status" gorm:"type:boolean"`
}

func NewUser(username string, password string, email string, mobile string, firstName string, lastName string) (*User, error) {
	if err := common.ValidateUsername(username); err != nil {
		return nil, err
	}
	if err := common.ValidatePassword(password); err != nil {
		return nil, err
	}
	if err := common.ValidateEmail(email); err != nil {
		return nil, err
	}
	if err := common.Mobile(mobile); err != nil {
		return nil, err
	}
	if err := common.Firstname(email); err != nil {
		return nil, err
	}
	if err := common.Lastname(lastName); err != nil {
		return nil, err
	}
	return &User{
		ID:        uuid.New(),
		Username:  username,
		Password:  password,
		Email:     email,
		Mobile:    mobile,
		FirstName: firstName,
		LastName:  lastName,
		createdAt: time.Now(),
		updatedAt: time.Now(),
		deletedAt: time.Time{},
		Status:    true,
	}, nil
}
