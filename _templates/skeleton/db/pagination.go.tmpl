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

	var pretty, filters, preloads string
	reqScheme := "http"

	if c.Request.TLS != nil {
		reqScheme = "https"
	}

	if _, ok := c.GetQuery("pretty"); ok {
		pretty = "&pretty"
	}

	if len(self.Filters) != 0 {
		filters = self.GetRawFilterQuery()
	}

	if self.Preloads != "" {
		preloads = fmt.Sprintf("&preloads=%v", self.Preloads)
	}

	if self.IsLastID {
		c.Header("Link", fmt.Sprintf("<%s://%v%v?limit=%v%s%s&last_id=%v&order=%v%s>; rel=\"next\"", reqScheme, c.Request.Host, c.Request.URL.Path, self.Limit, filters, preloads, index, self.Order, pretty))
		return nil
	}

	if self.Page == 1 {
		c.Header("Link", fmt.Sprintf("<%s://%v%v?limit=%v%s%s&page=%v%s>; rel=\"next\"", reqScheme, c.Request.Host, c.Request.URL.Path, self.Limit, filters, preloads, self.Page+1, pretty))
		return nil
	}

	c.Header("Link", fmt.Sprintf(
		"<%s://%v%v?limit=%v%s%s&page=%v%s>; rel=\"next\",<%s://%v%v?limit=%v%s%s&page=%v%s>; rel=\"prev\"", reqScheme,
		c.Request.Host, c.Request.URL.Path, self.Limit, filters, preloads, self.Page+1, pretty, reqScheme, c.Request.Host, c.Request.URL.Path, self.Limit, filters, preloads, self.Page-1, pretty))
	return nil
}
