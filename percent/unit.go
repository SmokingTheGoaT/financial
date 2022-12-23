package percent

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
	"strings"
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

func New(i interface{}) Percent {
	v := newDecimal(i)
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

func newDecimal(i interface{}) (d decimal.Decimal) {
	pf := parseFloat(i)
	v := pf / float64(100)
	d = decimal.NewFromFloat(v)
	return
}

func parseFloat(i interface{}) (d float64) {
	var err error
	switch i.(type) {
	case string:
		str := i.(string)
		str = strings.Trim(str, "%")
		d, err = strconv.ParseFloat(str, 64)
		if err != nil {
			panic(d)
		}
	case int:
		d = float64(i.(int))
	case int8:
		d = float64(i.(int8))
	case int16:
		d = float64(i.(int16))
	case int32:
		d = float64(i.(int32))
	case int64:
		d = float64(i.(int64))
	case uint:
		d = float64(i.(uint))
	case uint8:
		d = float64(i.(uint8))
	case uint16:
		d = float64(i.(uint16))
	case uint32:
		d = float64(i.(uint32))
	case uint64:
		d = float64(i.(uint64))
	case float32:
		d = float64(i.(float32))
	default:
		panic(fmt.Errorf("interface does not support a numeric string, int, uint or float32"))
	}
	return
}
