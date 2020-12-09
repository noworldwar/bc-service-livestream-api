package utils

import (
	"strconv"
	"strings"

	cur "github.com/PGITAb/bc-proto-enums-currency-go"
	"github.com/gin-gonic/gin"
)

func CheckPostFormData(c *gin.Context, vals ...string) string {
	for _, v := range vals {
		if strings.TrimSpace(c.PostForm(v)) == "" {
			return v
		}
	}
	return ""
}

func CheckMapData(data map[string]string, vals ...string) string {
	for _, v := range vals {
		if strings.TrimSpace(data[v]) == "" {
			return v
		}
	}
	return ""
}

func CheckQueryData(c *gin.Context, vals ...string) string {
	for _, v := range vals {
		if strings.TrimSpace(c.Query(v)) == "" {
			return v
		}
	}
	return ""
}

func IsInt(vals ...string) (b bool, s string) {
	for _, v := range vals {
		if v != "" {
			_, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				msg := v + " non-integer"
				return false, msg
			}
		}
	}

	return true, ""
}

func IsInt64(vals ...string) (b bool, s string) {
	for _, v := range vals {
		if v != "" {
			_, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				msg := v + " non-integer"
				return false, msg
			}
		}
	}

	return true, ""
}

// func CheckToken(token string, c *gin.Context) error {
// 	staff, err := model.RedisDB.Get(token).Result()
// 	if err != nil {
// 		return err
// 	}
// 	_ = model.RedisDB.Expire(token, model.RedisEX).Err()
// 	c.Set("Staff", staff)
// 	return nil
// }

func ContainsInt64(vals []int64, val int64) bool {
	for _, v := range vals {
		if v == val {
			return true
		}
	}

	return false
}

func CheckCurrency(currency string) (bool, cur.Currency) {
	i := cur.Currency_value[currency]
	if i == 0 && currency != "EUR" {
		return false, cur.Currency_EUR
	}
	return true, cur.Currency(cur.Currency_value[currency])
}
