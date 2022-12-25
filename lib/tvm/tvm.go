package tvm

import (
	"financial/utils/common"
	"financial/utils/percent"
	"financial/utils/types"
	"github.com/shopspring/decimal"
)

func fvif(rate percent.Percent, nper types.Period) (fvif decimal.Decimal) {
	fvif = rate.Decimal().
		Add(decimal.NewFromInt(1)).
		Pow(nper.Amount())
	return
}

func pvif(rate percent.Percent, nper types.Period) (pvif decimal.Decimal) {
	pvif = decimal.NewFromInt(1).
		Div(fvif(rate, nper))
	return
}

func annuityCertainPVIF(rate percent.Percent, nper types.Period, pd types.PaymentDue) (
	apvif decimal.Decimal) {
	if rate.Decimal().IsZero() {
		apvif = nper.Amount()
	} else {
		apvif = decimal.NewFromInt(1).
			Add(rate.Decimal().Mul(decimal.NewFromInt(int64(pd)))).
			Mul(decimal.NewFromInt(1).Sub(pvif(rate, nper))).
			Div(rate.Decimal())
	}
	return
}

func annuityCertainFVIF(rate percent.Percent, nper types.Period, pd types.PaymentDue) (
	afvif decimal.Decimal) {
	afvif = annuityCertainPVIF(rate, nper, pd).
		Mul(fvif(rate, nper))
	return
}

func nperif(rate percent.Percent, pmt decimal.Decimal, value decimal.Decimal, pd types.PaymentDue) (
	nperif decimal.Decimal) {
	nperif = value.Mul(rate.Decimal()).
		Add(pmt.Mul(decimal.NewFromInt(1).Add(rate.Decimal().Mul(decimal.NewFromInt(int64(pd))))))
	return
}

func pv(rate percent.Percent, nper types.Period, pmt decimal.Decimal, fv decimal.Decimal, pd types.PaymentDue) (
	pv decimal.Decimal) {
	neg := decimal.NewFromInt(1).Neg()
	pv = neg.Mul(fv.Mul(pvif(rate, nper).Add(pmt.Mul(annuityCertainPVIF(rate, nper, pd)))))
	return
}

func fv(rate percent.Percent, nper types.Period, pmt decimal.Decimal, pv decimal.Decimal, pd types.PaymentDue) (
	fv decimal.Decimal) {
	neg := decimal.NewFromInt(1).Neg()
	fv = neg.Mul(pv.Mul(fvif(rate, nper).Add(pmt.Mul(annuityCertainFVIF(rate, nper, pd)))))
	return
}

func pmt(rate percent.Percent, nper types.Period, pv decimal.Decimal, fv decimal.Decimal, pd types.PaymentDue) (
	pmt decimal.Decimal) {
	neg := decimal.NewFromInt(1).Neg()
	pmt = neg.Mul(pv.Add(fv.Mul(pvif(rate, nper)))).Div(annuityCertainPVIF(rate, nper, pd))
	return
}

func nper(rate percent.Percent, pmt decimal.Decimal, pv decimal.Decimal, fv decimal.Decimal, pd types.PaymentDue) (
	nper decimal.Decimal) {
	nper = common.Ln(nperif(rate, pmt, fv.Neg(), pd).Div(nperif(rate, pmt, pv, pd))).
		Div(common.Ln(rate.Decimal().Add(decimal.NewFromInt(1))))
	return
}
