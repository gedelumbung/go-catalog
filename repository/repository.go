package repository

import (
	"errors"

	"github.com/gedelumbung/go-catalog/model"
	"github.com/gedelumbung/go-catalog/params"
)

var (
	ErrNotFound = errors.New("item not found")
)

type Repository interface {
	Products() ProductRepository
	ProductImages() ProductImageRepository
}

type ProductRepository interface {
	All(params *params.ProductQueryParams) ([]*model.Product, int, error)
	FindByID(id int) (model.Product, error)
	Store(*model.Product) error
	Delete(id int) error
}

type ProductImageRepository interface {
	AllByProductID(productID int) ([]*model.ProductImage, error)
	FindByID(id int) (model.ProductImage, error)
}
