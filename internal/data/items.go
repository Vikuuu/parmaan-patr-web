package data

import (
	"database/sql"
	"errors"
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
	Price     int64     `json:"price"`
	Version   int64     `json:"version"`
}

func ValidateItem(v *validator.Validator, item *Item) {
	v.Check(item.Name != "", "name", "must be provided")
	v.Check(len(item.Name) <= 500, "name", "must not be more than 500 bytes long")

	v.Check(item.HsnSac != 0, "hsn_sac", "must be provided")

	v.Check(item.Gst != 0, "gst", "must be provided")
	v.Check(item.Gst <= 100, "gst", "must be less than 100")

	v.Check(item.Price != 0, "price", "must be provided")
}

// Define an ItemModel struct wrapping a sql.DB connection pool.
type ItemModel struct {
	DB *sql.DB
}

// Add placeholder methods for CRUD in the items table.
func (i ItemModel) Insert(item *Item) error {
	query := `
	INSERT INTO items (name, hsn, price, gst)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at, updated_at
	`

	// Create an args slice containing the values for the placeholder parameters from
	// the movie struct. Declaring this slice immediately next to our SQL query help to
	// make it nice and clear *what values are being used where* in the query.
	args := []any{item.Name, item.HsnSac, item.Price, item.Gst}

	return i.DB.QueryRow(query, args...).Scan(&item.ID, &item.CreatedAt, &item.UpdatedAt)
}

func (i ItemModel) Get(id int64) (*Item, error) {
	// To avoid making unnecessary DB call
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
	SELECT id, created_at, updated_at, name, hsn, price, gst, version
	FROM items
	WHERE id = $1
	`

	var item Item
	err := i.DB.QueryRow(query, id).Scan(
		&item.ID, &item.CreatedAt, &item.UpdatedAt,
		&item.Name, &item.HsnSac, &item.Price, &item.Gst,
		&item.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &item, nil
}

func (i ItemModel) Update(item *Item) error {
	query := `
	UPDATE items
	SET name = $1, hsn = $2, price = $3, gst = $4, updated_at = $5, version = version + 1
	where id = $6 AND version = $7
	RETURNING version
	`

	args := []any{
		item.Name, item.HsnSac, item.Price,
		item.Gst, time.Now(), item.ID, item.Version,
	}

	err := i.DB.QueryRow(query, args...).Scan(&item.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}

func (i ItemModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
	DELETE FROM items
	WHERE id = $1
	`

	result, err := i.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
