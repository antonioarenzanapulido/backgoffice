package data

import (
	"backgoffice/internal/validator"
	"math"
	"strings"
)

type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafelist []string
}

func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Page >= 1 && f.PageSize <= 10_000_000, "page", "must be between 1 and 10.000.000")
	v.Check(f.PageSize >= 1 && f.PageSize <= 100, "page_size", "must be between 1 and 100")
	v.Check(validator.PermittedValue(f.Sort, f.SortSafelist...), "sort", "invalid sort value")
}

func (f Filters) sortColumn() string {
	for _, safeValue := range f.SortSafelist {
		if f.Sort == safeValue {
			return strings.TrimPrefix(f.Sort, "-")
		}
	}

	panic("unsafe sort paramter: " + f.Sort)
}

func (f Filters) sortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}
	return "ASC"
}

func (f Filters) limit() int {
	return f.PageSize
}

func (f Filters) offset() int {
	return (f.Page - 1) * f.PageSize
}

type Metadata struct {
	CurrentPage int `json:"current_page,omitempty"`
	PageSize    int `json:"page_size,omitempty"`
	FirstPage   int `json:"first_page,omitempty"`
	LastPage    int `json:"last_page,omitempty"`
}

func calculateMetadata(totalRecords, page, pageSize int) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}

	return Metadata{
		CurrentPage: page,
		PageSize:    pageSize,
		FirstPage:   1,
		LastPage:    int(math.Ceil(float64(totalRecords) / float64(pageSize))),
	}
}
