package version

import (
	"math"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func New(c *gin.Context) (string, error) {
	ver := ""
	header := c.Request.Header.Get("Accept")
	header = strings.Join(strings.Fields(header), "")

	if strings.Contains(header, "version=") {
		ver = strings.Split(strings.SplitAfter(header, "version=")[1], ";")[0]
	}

	if v := c.Query("v"); v != "" {
		ver = v
	}

	if ver == "" {
		return "-1", nil
	}

	_, err := strconv.Atoi(strings.Join(strings.Split(ver, "."), ""))
	if err != nil {
		return "", err
	}

	return ver, nil
}

func Range(left string, op string, right string) bool {
	switch op {
	case "<":
		return (compare(left, right) == -1)
	case "<=":
		return (compare(left, right) <= 0)
	case ">":
		return (compare(left, right) == 1)
	case ">=":
		return (compare(left, right) >= 0)
	case "==":
		return (compare(left, right) == 0)
	}

	return false
}

func compare(left string, right string) int {
	// l > r : 1
	// l == r : 0
	// l < r : -1

	if left == "-1" {
		return 1
	} else if right == "-1" {
		return -1
	}

	lArr := strings.Split(left, ".")
	rArr := strings.Split(right, ".")
	lItems := len(lArr)
	rItems := len(rArr)
	min := int(math.Min(float64(lItems), float64(rItems)))

	for i := 0; i < min; i++ {
		l, _ := strconv.Atoi(lArr[i])
		r, _ := strconv.Atoi(rArr[i])

		if l != r {
			if l > r {
				return 1
			}

			return -1
		}
	}

	if lItems == rItems {
		return 0
	}

	if lItems < rItems {
		return 1
	}

	return -1
}
