package financial

import (
	"financial/lib/tvm"
	"financial/types"
	"financial/utils/percent"
	"github.com/shopspring/decimal"
)

func PV(rate percent.Percent, nper types.Period, pmt decimal.Decimal, fv decimal.Decimal, pd types.PaymentDue) (
	res decimal.Decimal, err error) {
	return tvm.PV(rate, nper, pmt, fv, pd)
}

func FV(rate percent.Percent, nper types.Period, pmt decimal.Decimal, pv decimal.Decimal, pd types.PaymentDue) (
	res decimal.Decimal, err error) {
	return tvm.FV(rate, nper, pmt, pv, pd)
}