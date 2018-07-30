package mysql

import (
	"github.com/gedelumbung/go-catalog/repository"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db                     *sqlx.DB
	productRepository      *productRepository
	productImageRepository *productImageRepository
}

func (s *Repository) Products() repository.ProductRepository {
	return s.productRepository
}

func (s *Repository) ProductImages() repository.ProductImageRepository {
	return s.productImageRepository
}

var _ repository.Repository = (*Repository)(nil)

func Connect(url string) (*Repository, error) {
	db, err := sqlx.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(20)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	s := &Repository{
		db:                     db,
		productRepository:      &productRepository{db: db},
		productImageRepository: &productImageRepository{db: db},
	}

	return s, nil
}
