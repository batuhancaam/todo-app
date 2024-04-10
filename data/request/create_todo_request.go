package request

import "time"

type CreateTodoRequest struct {
	Task      string    `validate:"required" json:"task"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Completed bool      `json:"completed"`
}
