package response

import "time"

type TodoResponse struct {
	ID        uint      `json:"id,omitempty"`
	Task      string    `json:"name,omitempty"`
	StartTime time.Time `json:"startTime,omitempty"`
	EndTime   time.Time `json:"endTime,omitempty"`
	Completed bool      `json:"completed,omitempty"`
}
