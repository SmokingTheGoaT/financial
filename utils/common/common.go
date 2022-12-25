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

func Not(ok bool) bool { return !ok }

func Raisable(b decimal.Decimal, p decimal.Decimal) bool {
	return Not(decimal.NewFromInt(1).Add(b).LessThan(decimal.NewFromInt(0)) &&
		Not(p.Sub(p).Equal(decimal.NewFromInt(0))))
}
