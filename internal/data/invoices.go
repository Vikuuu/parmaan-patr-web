package data

import "time"

type Invoice struct {
	ID         int64     `json:"id"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	To         string    `json:"to,omitempty"`
	From       string    `json:"from,omitempty"`
	Items      []string  `json:"items,omitempty"`
	TotalPrice uint32    `json:"total_price"`
}
