package model

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/gedelumbung/go-catalog/helper"
	"github.com/go-sql-driver/mysql"
)

type ProductImage struct {
	ID        int    `db:"id"`
	ProductID int    `db:"product_id"`
	Filename  string `db:"filename"`
	IsPrimary bool   `db:"is_primary"`
	Url       string
	CreatedAt mysql.NullTime `db:"created_at"`
	UpdatedAt mysql.NullTime `db:"updated_at"`
	DeletedAt mysql.NullTime `db:"deleted_at"`
}

func (o ProductImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID        int    `json:"id"`
		ProductID int    `json:"product_id"`
		Filename  string `json:"filename"`
		IsPrimary bool   `json:"is_primary"`
		Url       string `json:"url"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		DeletedAt string `json:"deleted_at,omitempty"`
	}{
		ID:        o.ID,
		ProductID: o.ProductID,
		Filename:  o.Filename,
		IsPrimary: o.IsPrimary,
		Url:       o.GetUrl(),
		CreatedAt: helper.NullTimeToString(o.CreatedAt, time.RFC3339),
		UpdatedAt: helper.NullTimeToString(o.UpdatedAt, time.RFC3339),
		DeletedAt: helper.NullTimeToString(o.DeletedAt, time.RFC3339),
	})
}

func (o ProductImage) GetUrl() string {
	if len(o.Filename) > 0 {
		imageID := strconv.Itoa(o.ID)
		productID := strconv.Itoa(o.ProductID)
		return `http://localhost:5050/v1/products/` + productID + `/images/` + imageID
	}
	return o.Filename
}
