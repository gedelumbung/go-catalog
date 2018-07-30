package api

import (
	"net/http"
	"strconv"

	"github.com/gedelumbung/go-catalog/helper"
	"github.com/gedelumbung/go-catalog/model"
	"github.com/gedelumbung/go-catalog/params"
	"github.com/gedelumbung/go-catalog/repository"
	"github.com/labstack/echo"
)

func (a *API) GetAllProducts(c echo.Context) error {
	var (
		products           []*model.Product
		strPage, strLimit  string
		page, limit, count int
		err                error
	)

	page = 1
	strPage = c.QueryParam("page")
	if len(strPage) > 0 {
		page, err = strconv.Atoi(strPage)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrRespond("client", "invalid request parameters", err.Error()))
		}
	}

	limit = defaultLimit
	strLimit = c.QueryParam("limit")
	if len(strLimit) > 0 {
		limit, err = strconv.Atoi(strLimit)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrRespond("client", "invalid request parameters", err.Error()))
		}
	}

	params := &params.ProductQueryParams{
		Limit: limit,
		Page:  page,
	}

	products, count, err = a.db.Products().All(params)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrRespond("server", "unable to load data from source", err))
	}

	return c.JSON(http.StatusOK, OKRespondWithMeta(products, map[string]interface{}{
		"pagination": NewPagination(count, page, limit),
	}))
}

func (a *API) GetProduct(c echo.Context) error {
	var (
		product model.Product
		images  []*model.ProductImage
		err     error
		id      int
	)
	errParams := map[string]string{}
	strId := c.Param("id")
	if id, err = strconv.Atoi(strId); err != nil {
		errParams["id"] = "invalid numeric value"
	}
	if len(errParams) > 0 {
		return c.JSON(http.StatusBadRequest, ErrRespond("client", "invalid request parameters", errParams))
	}

	product, err = a.db.Products().FindByID(id)

	if err != nil {
		if err.Error() == repository.ErrNotFound.Error() {
			return c.JSON(http.StatusNotFound, ErrRespond("server", "record not found", err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, ErrRespondString("server", "unable to load data from source", err.Error()))
	}

	images, err = a.db.ProductImages().AllByProductID(id)

	if err != nil {
		if err.Error() == repository.ErrNotFound.Error() {
			return c.JSON(http.StatusNotFound, ErrRespond("server", "record not found", err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, ErrRespondString("server", "unable to load data from source", err.Error()))
	}
	product.Images = images

	return c.JSON(http.StatusOK, OKRespond(product))
}

func (a *API) GetImage(c echo.Context) error {
	var (
		image       model.ProductImage
		err         error
		id, imageID int
	)
	errParams := map[string]string{}
	strId := c.Param("id")
	if id, err = strconv.Atoi(strId); err != nil {
		errParams["id"] = "invalid numeric value"
	}

	strImageId := c.Param("image_id")
	if imageID, err = strconv.Atoi(strImageId); err != nil {
		errParams["image_id"] = "invalid numeric value"
	}

	if len(errParams) > 0 {
		return c.JSON(http.StatusBadRequest, ErrRespond("client", "invalid request parameters", errParams))
	}

	_, err = a.db.Products().FindByID(id)

	if err != nil {
		if err.Error() == repository.ErrNotFound.Error() {
			return c.JSON(http.StatusNotFound, ErrRespond("server", "record not found", err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, ErrRespondString("server", "unable to load data from source", err.Error()))
	}

	image, err = a.db.ProductImages().FindByID(imageID)

	if err != nil {
		if err.Error() == repository.ErrNotFound.Error() {
			return c.JSON(http.StatusNotFound, ErrRespond("server", "record not found", err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, ErrRespondString("server", "unable to load data from source", err.Error()))
	}
	return c.File(`./storage/images/` + image.Filename)
}

func (a *API) StoreProduct(c echo.Context) error {
	errParams := map[string]string{}
	params := new(params.ProductRequest)
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ErrRespond("client", "invalid request parameters", err))
	}

	if len(params.Title) == 0 {
		errParams["title"] = "cannot be empty"
	}

	if len(params.Brand) == 0 {
		errParams["brand"] = "cannot be empty"
	}

	if params.Price == 0 {
		errParams["price"] = "cannot be empty"
	}

	if params.Quantity == 0 {
		errParams["quantity"] = "cannot be empty"
	}

	if len(params.AvailableSize) == 0 {
		errParams["available_sizes"] = "cannot be empty"
	}

	if len(errParams) > 0 {
		return c.JSON(http.StatusUnprocessableEntity, ErrRespond("client", "invalid request parameters", errParams))
	}

	product := model.Product{
		CategoryID:    params.CategoryID,
		Title:         params.Title,
		Brand:         params.Brand,
		Price:         params.Price,
		Description:   helper.StringToNullString(params.Description),
		Quantity:      params.Quantity,
		TryOutfit:     params.TryOutfit,
		AvailableSize: params.AvailableSize,
	}
	err := a.db.Products().Store(&product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrRespond("server", "error inserting new record", err))
	}
	return c.JSON(http.StatusOK, OKRespond(params))
}

func (a *API) DeleteProduct(c echo.Context) error {
	var (
		id  int
		err error
	)

	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ErrRespond("client", "invalid numeric value", err))
	}
	if err = a.db.Products().Delete(id); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrRespond("server", "failed delete record", err))
	}
	return c.NoContent(http.StatusNoContent)
}
