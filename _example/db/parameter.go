package db

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	defaultLimit = "25"
	defaultPage  = "1"
	defaultOrder = "desc"
)

type Parameter struct {
	Filters  map[string]string
	Preloads string
	Sort     string
	Limit    int
	Page     int
	LastID   int
	Order    string
	IsLastID bool
}

func NewParameter(c *gin.Context, model interface{}) (*Parameter, error) {
	parameter := &Parameter{}

	if err := parameter.initialize(c, model); err != nil {
		return nil, err
	}

	return parameter, nil
}

func (self *Parameter) initialize(c *gin.Context, model interface{}) error {
	self.Filters = filterToMap(c, model)
	self.Preloads = c.Query("preloads")
	self.Sort = c.Query("sort")

	limit, err := validate(c.DefaultQuery("limit", defaultLimit))
	if err != nil {
		return err
	}

	self.Limit = int(math.Max(1, math.Min(10000, float64(limit))))
	page, err := validate(c.DefaultQuery("page", defaultPage))
	if err != nil {
		return err
	}

	self.Page = int(math.Max(1, float64(page)))
	lastID, err := validate(c.Query("last_id"))
	if err != nil {
		return err
	}

	if lastID != -1 {
		self.IsLastID = true
		self.LastID = int(math.Max(0, float64(lastID)))
	}

	self.Order = c.DefaultQuery("order", defaultOrder)
	return nil
}

func validate(s string) (int, error) {
	if s == "" {
		return -1, nil
	}

	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return num, nil
}
