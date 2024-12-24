package model

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ScheduledAt time.Time `json:"Scheduled_at"`
	Started     bool      `json:"started"`
	IsDone      bool      `json:"is_done"`
	Logs        []TaskLog `json:"logs"`
}
type TaskLog struct {
	Id       uuid.UUID `json:"id"`
	Time     time.Time `json:"time"`
	Assignee string    `json:"assignee"`
}

func NewTask(name string, description string, scheduledAt time.Time, assignee string) Task {
	return Task{
		Id:          uuid.UUID{},
		Name:        name,
		Description: description,
		ScheduledAt: scheduledAt,
		Started:     false,
		IsDone:      false,
		Logs:        nil,
	}
}
