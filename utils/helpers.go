package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type (
	StringOption func(str string) string
)

func RemoveStrings(m map[string]string) StringOption {
	return func(str string) (res string) {
		res = str
		for i, v := range m {
			res = strings.Replace(str, i, v, -1)
		}
		return
	}
}

func ParseFloat(i interface{}, opts ...interface{}) (d float64) {
	var err error
	switch i.(type) {
	case string:
		str := i.(string)
		if len(opts) > 0 {
			opt, ok := opts[0].(StringOption)
			if !ok {
				panic(fmt.Errorf("trying to parse a string float with added operation on string, " +
					"missing StringOption"))
			}
			str = opt(str)
		}
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
	case float64:
		d = i.(float64)
	default:
		panic(fmt.Errorf("interface does not support a numeric string, int, uint or float32"))
	}
	return
}
