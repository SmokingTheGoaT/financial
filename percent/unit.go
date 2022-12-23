package percent

import (
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

func New(i interface{}) Percent {
	v := parseValue(i)
	return &unit{
		v: v,
	}
}

func (u *unit) Decimal() decimal.Decimal {
	return u.v
}

func (u *unit) String() (value string) {
	value = fmt.Sprintf("")
	return
}

func parseValue(i interface{}) (d decimal.Decimal) {
	//TODO: refactor the whole function.
	var err error
	switch i.(type) {
	case string:
		//TODO: check for '%' within string and remove it.
		//		convert string into a float and divide by float64(100)
		d, err = decimal.NewFromString(i.(string))
		if err != nil {
			panic(err)
		}
	case float32:
		d = decimal.NewFromFloat32(i.(float32))
	case float64:
		d = decimal.NewFromFloat(i.(float64))
	case int32:
		d = decimal.NewFromInt32(i.(int32))
	case int64:
		d = decimal.NewFromInt(i.(int64))
	default:
		panic(fmt.Errorf("the only types available for percent conversion are string number, " +
			"float32, float64, int32 and int64"))
	}
	return
}
