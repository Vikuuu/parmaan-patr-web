package data

import (
	"time"

	"github.com/Vikuuu/parmaan-patr-web/internal/validator"
)

type Item struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Name      string    `json:"name"`
	HsnSac    int64     `json:"hsn_sac"`
	Gst       int64     `json:"gst"`
	Price     string    `json:"price"`
}

func ValidateItem(v *validator.Validator, item *Item) {
	v.Check(item.Name != "", "name", "must be provided")
	v.Check(len(item.Name) <= 500, "name", "must not be more than 500 bytes long")

	v.Check(item.HsnSac != 0, "hsn_sac", "must be provided")

	v.Check(item.Gst != 0, "gst", "must be provided")
	v.Check(item.Gst <= 100, "gst", "must be less than 100")

	v.Check(item.Price != "", "price", "must be provided")
}
