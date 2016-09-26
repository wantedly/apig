package db

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func (self *Parameter) Paginate(db *gorm.DB) (*gorm.DB, error) {
	if self == nil {
		return nil, errors.New("Parameter struct got nil.")
	}

	if self.IsLastID {
		if self.Order == "asc" {
			return db.Where("id > ?", self.LastID).Limit(self.Limit).Order("id asc"), nil
		}

		return db.Where("id < ?", self.LastID).Limit(self.Limit).Order("id desc"), nil
	}

	return db.Offset(self.Limit * (self.Page - 1)).Limit(self.Limit), nil
}

func (self *Parameter) SetHeaderLink(c *gin.Context, index int) error {
	if self == nil {
		return errors.New("Parameter struct got nil.")
	}

	var filters, preloads, sort string
	reqScheme := "http"

	if c.Request.TLS != nil {
		reqScheme = "https"
	}

	if len(self.Filters) != 0 {
		filters = self.GetRawFilterQuery()
	}

	if self.Preloads != "" {
		preloads = fmt.Sprintf("&preloads=%v", self.Preloads)
	}

	if self.Sort != "" {
		sort = fmt.Sprintf("&sort=%v", self.Sort)
	}

	if self.IsLastID {
		c.Header("Link", fmt.Sprintf("<%s://%v%v?limit=%v%v%v%v&last_id=%v&order=%v>; rel=\"next\"", reqScheme, c.Request.Host, c.Request.URL.Path, self.Limit, filters, preloads, sort, index, self.Order))
		return nil
	}

	if self.Page == 1 {
		c.Header("Link", fmt.Sprintf("<%s://%v%v?limit=%v%v%v%v&page=%v>; rel=\"next\"", reqScheme, c.Request.Host, c.Request.URL.Path, self.Limit, filters, preloads, sort, self.Page+1))
		return nil
	}

	c.Header("Link", fmt.Sprintf(
		"<%s://%v%v?limit=%v&page=%v>; rel=\"next\",<http://%v%v?limit=%v%v%v%v&page=%v>; rel=\"prev\"", reqScheme,
		c.Request.Host, c.Request.URL.Path, self.Limit, self.Page+1, c.Request.Host, c.Request.URL.Path, self.Limit, filters, preloads, sort, self.Page-1,
	))
	return nil
}
