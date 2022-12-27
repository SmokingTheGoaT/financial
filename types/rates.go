package types

import (
	"financial/utils"
	"fmt"
	"github.com/shopspring/decimal"
)

//TODO: Refactor the following type into a rate unit instead.
//The new Rate type should include the following, rate Period.

type (
	unit struct {
		rate decimal.Decimal
	}

	Rate interface {
		Decimal() decimal.Decimal
		Percent() string
	}
)

func New(i interface{}, opts ...interface{}) Rate {
	v := newDecimal(i, opts...)
	return &unit{
		rate: v,
	}
}

func (u *unit) Decimal() decimal.Decimal {
	return u.rate
}

func (u *unit) Percent() (value string) {
	d2 := decimal.NewFromInt(100)
	value = fmt.Sprintf("%s%%", u.rate.Mul(d2).String())
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
