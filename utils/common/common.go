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

func Pow(a decimal.Decimal, b decimal.Decimal) (t decimal.Decimal) {
	return a.Pow(b)
}

func HaveRightSigns(x decimal.Decimal, y decimal.Decimal, z decimal.Decimal) bool {
	return Not(x.Sign() == y.Sign() && y.Sign() == z.Sign()) &&
		Not(x.Sign() == y.Sign() && z.Equal(decimal.NewFromInt(0))) &&
		Not(x.Sign() == z.Sign() && y.Equal(decimal.NewFromInt(0))) &&
		Not(y.Sign() == z.Sign() && x.Equal(decimal.NewFromInt(0)))
}
