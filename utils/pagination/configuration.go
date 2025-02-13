package pagination

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"strconv"
	"test-case-roketin/common/constants"
)

type Pagination struct {
	Limit      *int   `json:"limit,omitempty;query:limit"`
	Page       *int   `json:"page,omitempty;query:page"`
	TotalRows  *int64 `json:"total_rows"`
	TotalPages *int   `json:"total_pages"`
}

func NewPagination() *Pagination {
	return &Pagination{}
}

func (p *Pagination) GetOffset() *int {
	if p.GetPage() == nil || p.GetLimit() == nil {
		return nil
	} else {
		var i = (*p.GetPage() - 1) * *p.GetLimit()

		return &i
	}
}

func (p *Pagination) GetLimit() *int {
	return p.Limit
}

func (p *Pagination) GetPage() *int {
	return p.Page
}

func (p *Pagination) GetPagination(c *gin.Context) (Pagination, string) {
	limit := p.GetLimit()
	page := p.GetPage()

	var search string

	query := c.Request.URL.Query()
	for key, val := range query {
		queryValue := val[len(val)-1]
		switch key {
		case "limit":
			temp, _ := strconv.Atoi(queryValue)
			limit = &temp
			break
		case "page":
			temp, _ := strconv.Atoi(queryValue)
			page = &temp
			break
		case "search":
			search = queryValue
			break
		}
	}

	if *limit == 0 {
		temp := 10
		limit = &temp
	}

	if *page == 0 {
		temp := 1
		page = &temp
	}

	return Pagination{Limit: limit, Page: page}, search
}

func Paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if pagination.GetPage() == nil || pagination.GetLimit() == nil {
		return func(db *gorm.DB) *gorm.DB {
			return db.Order(constants.UpdatedAt + " " + constants.Desc)
		}
	} else {
		var totalRows int64
		db.Model(value).Count(&totalRows)

		pagination.TotalRows = &totalRows
		totalPages := int(math.Ceil(float64(totalRows) / float64(*pagination.Limit)))
		pagination.TotalPages = &totalPages

		return func(db *gorm.DB) *gorm.DB {
			return db.Offset(*pagination.GetOffset()).Limit(*pagination.GetLimit()).Order(constants.UpdatedAt + " " + constants.Desc)
		}
	}
}
