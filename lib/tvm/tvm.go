package tvm

import (
	"financial/utils/common"
	"financial/utils/percent"
	"financial/utils/types"
	"fmt"
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

func PV(rate percent.Percent, nper types.Period, pmt decimal.Decimal, fv decimal.Decimal, pd types.PaymentDue) (
	res decimal.Decimal, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			if er, ok := rec.(error); ok {
				err = er
			}
		}
	}()
	if common.Raisable(rate.Decimal(), nper.Amount()) {
		err = fmt.Errorf("r is not raisable to nper (r is less than -1 and nper not an integer")
	} else if common.Not(pmt.Equal(decimal.NewFromInt(0))) ||
		common.Not(fv.Equals(decimal.NewFromFloat(0))) {
		err = fmt.Errorf("pmt or fv need to be different from 0")
	} else if !rate.Decimal().Equal(percent.New("100%").Decimal().Neg()) {
		err = fmt.Errorf("r cannot be -100%%")
	} else {
		res = pv(rate, nper, pmt, fv, pd)
	}
	return
}

func FV(rate percent.Percent, nper types.Period, pmt decimal.Decimal, pv decimal.Decimal, pd types.PaymentDue) (
	res decimal.Decimal, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			if er, ok := rec.(error); ok {
				err = er
			}
		}
	}()
	if common.Raisable(rate.Decimal(), nper.Amount()) {
		err = fmt.Errorf("r is not raisable to nper (r is negative and nper not an integer")
	} else if common.Not(rate.Decimal().Equal(decimal.NewFromInt(1).Neg())) ||
		(rate.Decimal().Equals(decimal.NewFromFloat(-1)) &&
			nper.Amount().GreaterThan(decimal.NewFromInt(0))) {
		err = fmt.Errorf("r cannot be -100%% when nper is <= 0")
	} else {
		if rate.Decimal().Equal(decimal.NewFromInt(1).Neg()) &&
			pd.CompareTo(types.BeginningOfPeriod) {
			res = decimal.NewFromInt(1).Neg().Mul(pv.Mul(fvif(rate, nper)))
		} else if rate.Decimal().Equal(decimal.NewFromInt(1).Neg()) &&
			pd.CompareTo(types.EndOfPeriod) {
			res = decimal.NewFromInt(1).Neg().Mul(pv.Mul(fvif(rate, nper).Add(pmt)))
		} else {
			res = fv(rate, nper, pmt, pv, pd)
		}
	}
	return
}

func PMT(rate percent.Percent, nper types.Period, pv decimal.Decimal, fv decimal.Decimal, pd types.PaymentDue) (
	res decimal.Decimal, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			if er, ok := rec.(error); ok {
				err = er
			}
		}
	}()
	if common.Raisable(rate.Decimal(), nper.Amount()) {
		err = fmt.Errorf("r is not raisable to nper (r is negative and nper not an integer")
	} else if common.Not(rate.Decimal().Equal(decimal.NewFromInt(1).Neg())) || (rate.Decimal().Equal(decimal.NewFromInt(1).Neg()) &&
		nper.Amount().GreaterThan(decimal.NewFromInt(0)) &&
		pd.CompareTo(types.EndOfPeriod)) {
		err = fmt.Errorf("r cannot be -100%% when nper is <= 0")
	} else if common.Not(annuityCertainPVIF(rate, nper, pd).Equals(decimal.NewFromInt(0))) {
		err = fmt.Errorf("1 * pd + 1 - (1 / (1 + r)^nper) / nper has to be <> 0")
	} else if rate.Decimal().Equal(decimal.NewFromInt(1).Neg()) {
		res = fv.Neg()
	} else {
		res = pmt(rate, nper, pv, fv, pd)
	}
	return
}
