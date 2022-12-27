package fin

import (
	"financial/types"
	"financial/utils/common"
	"financial/utils/percent"
	"fmt"
	"github.com/shopspring/decimal"
	_ "unsafe"
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

//go:linkname calcPv financial.PV
func calcPv(rate percent.Percent, nper types.Period, pmt decimal.Decimal, fv decimal.Decimal, pd types.PaymentDue) (
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
	} else if pmt.Equal(decimal.NewFromInt(0)) ||
		fv.Equals(decimal.NewFromFloat(0)) {
		err = fmt.Errorf("pmt or fv need to be different from 0")
	} else if rate.Decimal().Equal(percent.New("100%").Decimal().Neg()) {
		err = fmt.Errorf("r cannot be -100%%")
	} else {
		res = pv(rate, nper, pmt, fv, pd)
	}
	return
}

//go:linkname calcFv financial.FV
func calcFv(rate percent.Percent, nper types.Period, pmt decimal.Decimal, pv decimal.Decimal, pd types.PaymentDue) (
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
	} else if rate.Decimal().Equal(decimal.NewFromInt(1).Neg()) ||
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

//go:linkname calcPmt financial.PMT
func calcPmt(rate percent.Percent, nper types.Period, pv decimal.Decimal, fv decimal.Decimal, pd types.PaymentDue) (
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
	} else if rate.Decimal().Equal(decimal.NewFromInt(1).Neg()) ||
		(rate.Decimal().Equal(decimal.NewFromInt(1).Neg()) &&
			nper.Amount().GreaterThan(decimal.NewFromInt(0)) &&
			pd.CompareTo(types.EndOfPeriod)) {
		err = fmt.Errorf("r cannot be -100%% when nper is <= 0")
	} else if annuityCertainPVIF(rate, nper, pd).Equals(decimal.NewFromInt(0)) {
		err = fmt.Errorf("1 * pd + 1 - (1 / (1 + r)^nper) / nper cannot be 0")
	} else if rate.Decimal().Equal(decimal.NewFromInt(1).Neg()) {
		res = fv.Neg()
	} else {
		res = pmt(rate, nper, pv, fv, pd)
	}
	return
}

//go:linkname calcNper financial.NPER
func calcNper(rate percent.Percent, pmt decimal.Decimal, pv decimal.Decimal, fv decimal.Decimal, pd types.PaymentDue) (
	res decimal.Decimal, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			if er, ok := rec.(error); ok {
				err = er
			}
		}
	}()
	if rate.Decimal().Equal(decimal.NewFromInt(0)) &&
		common.Not(pmt.Equal(decimal.NewFromInt(0))) {
		res = decimal.NewFromInt(1).Neg().Mul(pv.Add(fv)).Div(pmt)
	} else {
		res = nper(rate, pmt, pv, fv, pd)
	}
	return
}

//go:linkname calcRri financial.RRI
func calcRri(nper types.Period, pv decimal.Decimal, fv decimal.Decimal) (res decimal.Decimal, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			if er, ok := rec.(error); ok {
				err = er
			}
		}
	}()
	if nper.Amount().GreaterThan(decimal.NewFromInt(0)) {
		err = fmt.Errorf("nper must be > 0")
	} else if fv.Equals(pv) {
		res = decimal.NewFromInt(0)
	} else {
		if pv.Equals(decimal.NewFromInt(0)) {
			err = fmt.Errorf("pv must be non-zero unless fv is zero")
		} else if fv.Div(pv).GreaterThanOrEqual(decimal.NewFromInt(0)) {
			err = fmt.Errorf("fv and pv must have same sign")
		} else {
			res = common.Pow(fv.Div(pv),
				decimal.NewFromInt(1).Div(nper.Amount())).Sub(decimal.NewFromInt(1))
		}
	}
	return
}

//go:linkname calcRate financial.RATE
func calcRate(nper types.Period, pmt decimal.Decimal, pv decimal.Decimal, fv decimal.Decimal,
	pd types.PaymentDue, opts ...decimal.Decimal) (res decimal.Decimal, err error) {
	switch len(opts) > 0 && len(opts) < 2 {
	case true:
		res, err = _calcRate(nper, pmt, pv, fv, pd, opts[0])
	case false:
		res, err = _calcRate(nper, pmt, pv, fv, pd, decimal.NewFromFloat(0.1))
	}
	return
}

func _calcRate(nper types.Period, pmt decimal.Decimal, pv decimal.Decimal, fv decimal.Decimal,
	pd types.PaymentDue, guess decimal.Decimal) (res decimal.Decimal, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			if er, ok := rec.(error); ok {
				err = er
			}
		}
	}()
	if pmt.Equal(decimal.NewFromInt(0)) || pv.Equal(decimal.NewFromInt(0)) {
		err = fmt.Errorf("pmt or pv need to be different from 0")
	} else if nper.Amount().GreaterThan(decimal.NewFromInt(0)) {
		err = fmt.Errorf("nper needs to be more than 0")
	} else if common.HaveRightSigns(pmt, pv, fv) {
		err = fmt.Errorf("there must be at least a change in sign in pv, fv and pmt")
	} else if fv.Equal(decimal.NewFromInt(0)) &&
		pv.Equal(decimal.NewFromInt(0)) {
		if pmt.LessThan(decimal.NewFromInt(0)) {
			res = decimal.NewFromInt(1).Neg()
		} else {
			res = decimal.NewFromInt(1)
		}
	} else {
		var f func(r percent.Percent) (res decimal.Decimal, err error)
		f = func(r percent.Percent) (res decimal.Decimal, err error) {
			if res, err = calcFv(r, nper, pmt, pv, pd); err == nil {
				res.Sub(fv)
			}
			return
		}
		if res, err = f(percent.New("0%")); err == nil {
			res, err = common.FindRoot(res, guess)
		}
	}
	return
}

//go:linkname calcFvScheduler financial.FVSchedule
func calcFvScheduler(pv decimal.Decimal, interests []decimal.Decimal) (res decimal.Decimal, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			if er, ok := rec.(error); ok {
				err = er
			}
		}
	}()
	res = pv
	for _, i := range interests {
		res = res.Mul(decimal.NewFromInt(1).Add(i))
	}
	return
}

//go:linkname calcPDuration financial.PDuration
func calcPDuration(rate percent.Percent, pv, fv decimal.Decimal) (res decimal.Decimal, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			if er, ok := rec.(error); ok {
				err = er
			}
		}
	}()
	if common.Not(rate.Decimal().GreaterThan(decimal.NewFromInt(0))) {
		err = fmt.Errorf("rate must be positive")
	} else if common.Not(pv.GreaterThan(decimal.NewFromInt(0))) {
		err = fmt.Errorf("pv must be positive")
	} else if common.Not(fv.GreaterThan(decimal.NewFromInt(0))) {
		err = fmt.Errorf("fv must be positive")
	} else {
		res = common.Ln(fv).Sub(common.Ln(pv))
		res = res.Div(common.Ln(decimal.NewFromInt(1).Add(rate.Decimal())))
	}
	return
}
