package mysql

import (
	"database/sql"

	"github.com/gedelumbung/go-catalog/model"
	"github.com/gedelumbung/go-catalog/repository"
	"github.com/jmoiron/sqlx"
)

type productImageRepository struct {
	db *sqlx.DB
}

const selectProductImages = `select * from product_images pi where pi.deleted_at is null`

func (o *productImageRepository) AllByProductID(productID int) ([]*model.ProductImage, error) {
	var (
		images []*model.ProductImage
		count  int
	)
	images = []*model.ProductImage{}

	err := o.db.QueryRowx(`select count(*) from product_images where deleted_at is null`).Scan(&count)
	if err != nil {
		return images, err
	}

	err = o.db.Select(&images, selectProductImages+` and pi.product_id = ?`, productID)
	return images, nil
}

func (o *productImageRepository) FindByID(id int) (model.ProductImage, error) {
	var image model.ProductImage
	err := o.db.QueryRowx(selectProductImages+` and pi.id = ?`, id).StructScan(&image)
	if err == sql.ErrNoRows {
		return image, repository.ErrNotFound
	}
	if err != nil {
		return image, err
	}
	return image, err
}
