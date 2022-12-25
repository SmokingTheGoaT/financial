package percent

import (
	"financial/utils"
	"fmt"
	"github.com/shopspring/decimal"
)

type (
	unit struct {
		v decimal.Decimal
	}

	Percent interface {
		Decimal() decimal.Decimal
		String() (value string)
	}
)

func New(i interface{}, opts ...interface{}) Percent {
	v := newDecimal(i, opts...)
	return &unit{
		v: v,
	}
}

func (u *unit) Decimal() decimal.Decimal {
	return u.v
}

func (u *unit) String() (value string) {
	d2 := decimal.NewFromInt(100)
	value = fmt.Sprintf("%s%%", u.v.Mul(d2).String())
	return
}

func newDecimal(i interface{}, opts ...interface{}) (d decimal.Decimal) {
	var u float64
	if len(opts) > 0 {
		d := opts[0]
		u = utils.ParseFloat(d, utils.RemoveStrings(map[string]string{"%": ""}))
	} else {
		u = 100
	}
	pf := utils.ParseFloat(i, utils.RemoveStrings(map[string]string{"%": ""}))
	v := pf / u
	d = decimal.NewFromFloat(v)
	return
}
