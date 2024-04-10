package request

import "time"

type UpdateTodoRequest struct {
	ID        uint      `validate:"required"`
	Task      string    `validate:"required,min=1,max=255" json:"task"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Completed bool      `json:"completed"`
}
