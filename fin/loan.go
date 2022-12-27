package fin

import (
	"financial/types"
	"github.com/shopspring/decimal"
	_ "unsafe"
)

func ipmt(r types.Rate, per, nper, pv, fv decimal.Decimal, pd types.PaymentDue) (ip decimal.Decimal) {
	ip = decimal.NewFromInt(1).Neg().Mul(
		pv.Mul(
			fvif(r, per.Sub(decimal.NewFromInt(1))).
				Mul(r.Decimal()).
				Add(pmt(r, nper, pv, fv, types.EndOfPeriod).Mul(fvif(r, per.Sub(decimal.NewFromInt(1))))).
				Sub(decimal.NewFromInt(1))))
	switch pd.CompareTo(types.EndOfPeriod) {
	case true:
		break
	case false:
		ip.Div(decimal.NewFromInt(1).Add(r.Decimal()))
	}
	return
}

func ppmt(r types.Rate, per, nper, pv, fv decimal.Decimal, pd types.PaymentDue) (pm decimal.Decimal) {
	return pmt(r, nper, pv, fv, pd).Sub(ipmt(r, per, nper, pv, fv, pd))
}

func ispmt(r types.Rate, per, nper, pv decimal.Decimal) (ipm decimal.Decimal) {
	coupon := pv.Neg().Mul(r.Decimal())
	return coupon.Sub(coupon.Div(nper.Mul(per)))
}
