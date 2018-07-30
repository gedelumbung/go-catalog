package mysql

import (
	"database/sql"
	"errors"
	"time"

	"github.com/gedelumbung/go-catalog/helper"
	"github.com/gedelumbung/go-catalog/model"
	"github.com/gedelumbung/go-catalog/params"
	"github.com/gedelumbung/go-catalog/repository"
	"github.com/jmoiron/sqlx"
)

type productRepository struct {
	db *sqlx.DB
}

const selectProducts = `select 
	pi.filename "image",
	ifnull(pi.id, 0) "primary_image_id",
	p.*,
	c.id "category.id",
	c.title "category.title",
	c.created_at "category.created_at",
	c.updated_at "category.updated_at",
	c.deleted_at "category.deleted_at"
	from products p join categories c
	on p.category_id = c.id
 	left join (select id, product_id, filename from product_images where is_primary = 1) pi 
	on p.id = pi.product_id
	where p.deleted_at is null`

func (o *productRepository) All(params *params.ProductQueryParams) ([]*model.Product, int, error) {
	var (
		products     []*model.Product
		count, start int
	)
	products = []*model.Product{}

	err := o.db.QueryRowx(`select count(*) from products where deleted_at is null`).Scan(&count)
	if err != nil {
		return products, 0, err
	}

	start = (params.Page - 1) * params.Limit
	if start < 0 || params.Limit < 1 {
		return products, 0, errors.New("insufficient parameters")
	}

	err = o.db.Select(&products, selectProducts+` limit ? offset ?`, params.Limit, start)
	return products, count, nil
}

func (o *productRepository) FindByID(id int) (model.Product, error) {
	var product model.Product
	err := o.db.QueryRowx(selectProducts+` and p.id = ?`, id).StructScan(&product)
	if err == sql.ErrNoRows {
		return product, repository.ErrNotFound
	}
	if err != nil {
		return product, err
	}
	return product, err
}

func (o *productRepository) Store(model *model.Product) error {
	stmt, err := o.db.Preparex(`insert into products
		(category_id, title, brand, price, description, quantity, try_outfit, available_sizes, created_at, updated_at)
		values (?, ?, ?, ?, ?, ?, ?, ?, now(), now())`)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(
		model.CategoryID,
		model.Title,
		model.Brand,
		model.Price,
		model.Description,
		model.Quantity,
		model.TryOutfit,
		model.AvailableSize,
	)

	if err != nil {
		return err
	}

	now := time.Now()
	id, err := result.LastInsertId()
	model.ID = int(id)
	model.CreatedAt = helper.TimeToNullTime(now)
	model.UpdatedAt = helper.TimeToNullTime(now)

	return err
}

func (o *productRepository) Delete(id int) error {
	stmt, err := o.db.Preparex(`update products set deleted_at = now() where id = ?`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		id,
	)

	if err != nil {
		return err
	}

	return err
}
