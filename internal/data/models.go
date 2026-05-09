package data

import (
	"database/sql"
	"errors"
)

var ErrRecordNotFound = errors.New("record not found")

// Models struct for wrapping all the models struct.
type Models struct {
	Items ItemModel
}

func NewModels(db *sql.DB) Models {
	return Models{Items: ItemModel{DB: db}}
}
