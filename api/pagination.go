package api

const defaultLimit = 10

type Pagination struct {
	Total       int `json:"total"`
	Limit       int `json:"limit"`
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
}

func NewPagination(total, page, limit int) *Pagination {
	lastPage := calculateLastPage(limit, total)
	return &Pagination{Total: total, Limit: limit, CurrentPage: page, LastPage: lastPage}
}

func calculateLastPage(perPage, total int) int {
	pages := total / perPage
	if total%perPage > 0 {
		return pages + 1
	}
	return pages
}
