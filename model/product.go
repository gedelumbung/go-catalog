package model

import (
	"database/sql"
	"encoding/json"
	"strconv"
	"time"

	"github.com/gedelumbung/go-catalog/helper"
	"github.com/go-sql-driver/mysql"
)

type Product struct {
	ID             int             `db:"id"`
	Title          string          `db:"title"`
	CategoryID     int             `db:"category_id"`
	Category       Category        `db:"category"`
	Image          *string         `db:"image"`
	PrimaryImageID *int            `db:"primary_image_id"`
	Images         []*ProductImage `db:"images"`
	Brand          string          `db:"brand"`
	Price          int             `db:"price"`
	Description    sql.NullString  `db:"description"`
	Quantity       int             `db:"quantity"`
	TryOutfit      bool            `db:"try_outfit"`
	AvailableSize  string          `db:"available_sizes"`
	CreatedAt      mysql.NullTime  `db:"created_at"`
	UpdatedAt      mysql.NullTime  `db:"updated_at"`
	DeletedAt      mysql.NullTime  `db:"deleted_at"`
}

func (o Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID             int             `json:"id"`
		Title          string          `json:"title"`
		Category       Category        `json:"category"`
		Image          *string         `json:"image,omitempty"`
		PrimaryImageID *int            `json:"primary_image_id,omitempty"`
		Images         []*ProductImage `json:"images,omitempty"`
		Brand          string          `json:"brand"`
		Price          int             `json:"price"`
		Description    string          `json:"description"`
		Quantity       int             `json:"quantity"`
		TryOutfit      bool            `json:"try_outfit"`
		AvailableSize  string          `json:"available_sizes"`
		CreatedAt      string          `json:"created_at"`
		UpdatedAt      string          `json:"updated_at"`
		DeletedAt      string          `json:"deleted_at,omitempty"`
	}{
		ID:            o.ID,
		Title:         o.Title,
		Category:      o.Category,
		Image:         o.GetUrl(),
		Images:        o.Images,
		Brand:         o.Brand,
		Price:         o.Price,
		Description:   helper.NullStringToString(o.Description),
		Quantity:      o.Quantity,
		TryOutfit:     o.TryOutfit,
		AvailableSize: o.AvailableSize,
		CreatedAt:     helper.NullTimeToString(o.CreatedAt, time.RFC3339),
		UpdatedAt:     helper.NullTimeToString(o.UpdatedAt, time.RFC3339),
		DeletedAt:     helper.NullTimeToString(o.DeletedAt, time.RFC3339),
	})
}

func (o Product) GetUrl() *string {
	var urlPointer = new(string)
	if *o.PrimaryImageID > 0 {
		imageID := strconv.Itoa(*o.PrimaryImageID)
		productID := strconv.Itoa(o.ID)
		url := `http://localhost:5050/v1/products/` + productID + `/images/` + imageID
		*urlPointer = url
		return urlPointer
	}
	return o.Image
}
