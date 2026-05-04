package data

import "time"

type Item struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Name      string    `json:"name"`
	HsnSac    int64     `json:"hsn_sac"`
	Gst       int64     `json:"gst"`
	Price     string    `json:"price"`
}
