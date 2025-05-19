package models

import "time"

// модель форм
type Form struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Schema    string    `json:"schema"`
	CreatedAt time.Time `json:"created_at"`
}

// модель для работы с формой
type Response struct {
	ID        int       `json:"id"`
	FormID    int       `json:"form_id"`
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
	Status    string    `json:"status"`
}
