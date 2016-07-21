package db

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Pagination struct {
	Limit   int
	Page    int
	Last_ID int
	Order   string
}

func (p *Pagination) Paginate(c *gin.Context) (*gorm.DB, error) {
	db := DBInstance(c)
	limit_query := c.DefaultQuery("limit", "25")
	page_query := c.DefaultQuery("page", "1")
	last_id_query := c.Query("last_id")
	p.Order = c.DefaultQuery("order", "desc")

	limit, err := strconv.Atoi(limit_query)
	if err != nil {
		return db, errors.New("invalid parameter.")
	}
	p.Limit = int(math.Max(1, math.Min(10000, float64(limit))))

	if last_id_query != "" {
		// pagination 1
		last_id, err := strconv.Atoi(last_id_query)
		if err != nil {
			return db, errors.New("invalid parameter.")
		}
		p.Last_ID = int(math.Max(0, float64(last_id)))
		if p.Order == "asc" {
			return db.Where("id > ?", p.Last_ID).Limit(p.Limit).Order("id asc"), nil
		} else {
			return db.Where("id < ?", p.Last_ID).Limit(p.Limit).Order("id desc"), nil
		}
	}

	// pagination 2
	page, err := strconv.Atoi(page_query)
	if err != nil {
		return db, errors.New("invalid parameter.")
	}
	p.Page = int(math.Max(1, float64(page)))
	return db.Offset(limit * (p.Page - 1)).Limit(p.Limit), nil
}

func (p *Pagination) SetHeaderLink(c *gin.Context, index int) {
	var link string
	if p.Last_ID != 0 {
		link = fmt.Sprintf("<http://%v%v?limit=%v&last_id=%v&order=%v>; rel=\"next\"", c.Request.Host, c.Request.URL.Path, p.Limit, index, p.Order)
	} else {
		if p.Page == 1 {
			link = fmt.Sprintf("<http://%v%v?limit=%v&page=%v>; rel=\"next\"", c.Request.Host, c.Request.URL.Path, p.Limit, p.Page+1)
		} else {
			link = fmt.Sprintf("<http://%v%v?limit=%v&page=%v>; rel=\"next\",<http://%v%v?limit=%v&page=%v>; rel=\"prev\"", c.Request.Host, c.Request.URL.Path, p.Limit, p.Page+1, c.Request.Host, c.Request.URL.Path, p.Limit, p.Page-1)
		}
	}
	c.Header("Link", link)
}
