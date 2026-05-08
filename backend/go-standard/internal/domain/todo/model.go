package todo

import "time"

type Todo struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      int        `json:"status"`
	Create_at   *time.Time `json:"create_at"`
	Update_at   *time.Time `json:"update_at"`
}
