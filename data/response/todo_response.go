package response

import "time"

type TodoResponse struct {
	ID        uint      `json:"id"`
	Task      string    `json:"name"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Completed bool      `json:"completed"`
}
