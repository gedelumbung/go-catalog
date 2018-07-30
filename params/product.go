package params

type ProductQueryParams struct {
	Page    int `json:"page" query:"page"`
	Limit   int `json:"limit" query:"limit"`
	OrderBy int `json:"order_by" query:"order_by"`
	Sort    int `json:"sort" query:"sort"`
}

type ProductRequest struct {
	Title         string `json:"title"`
	CategoryID    int    `json:"category_id"`
	Brand         string `json:"brand"`
	Price         int    `json:"price"`
	Description   string `json:"description"`
	Quantity      int    `json:"quantity"`
	TryOutfit     bool   `json:"try_outfit"`
	AvailableSize string `json:"available_sizes"`
}
