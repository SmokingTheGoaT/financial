package common

import (
	"github.com/shopspring/decimal"
	"math"
	"strconv"
)

func Ln(x decimal.Decimal) (t decimal.Decimal) {
	str := x.String()
	d, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(err)
	}
	t = decimal.NewFromFloat(math.Log(d))
	return
}
