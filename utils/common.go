package utils

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
	"strconv"
)

func Ln(x decimal.Decimal) (t decimal.Decimal) {
	str := x.String()
	d, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(err)
	}
	t = decimal.NewFromFloat(math.Log(d))
	return
}

func Not(ok bool) bool { return !ok }

func Raisable(b decimal.Decimal, p decimal.Decimal) bool {
	return decimal.NewFromInt(1).Add(b).LessThan(decimal.NewFromInt(0)) &&
		Not(p.Sub(p).Equal(decimal.NewFromInt(0)))
}

func Pow(a decimal.Decimal, b decimal.Decimal) (t decimal.Decimal) {
	return a.Pow(b)
}

func HaveRightSigns(x decimal.Decimal, y decimal.Decimal, z decimal.Decimal) bool {
	return Not(x.Sign() == y.Sign() && y.Sign() == z.Sign()) &&
		Not(x.Sign() == y.Sign() && z.Equal(decimal.NewFromInt(0))) &&
		Not(x.Sign() == z.Sign() && y.Equal(decimal.NewFromInt(0))) &&
		Not(y.Sign() == z.Sign() && x.Equal(decimal.NewFromInt(0)))
}

func Newton(f decimal.Decimal, x decimal.Decimal, count decimal.Decimal, precesion decimal.Decimal) (
	n decimal.NullDecimal) {
	maxCount := decimal.NewFromInt(20)
	var helper func(f decimal.Decimal, x decimal.Decimal, count decimal.Decimal, precesion decimal.Decimal) (
		h decimal.NullDecimal)
	helper = func(f decimal.Decimal, x decimal.Decimal, count decimal.Decimal, precesion decimal.Decimal) (
		h decimal.NullDecimal) {
		d := func(f decimal.Decimal, x decimal.Decimal) (td decimal.Decimal) {
			td = f.Mul(x.Add(precesion)).Sub(f.Mul(x.Sub(precesion)))
			td = td.Div(decimal.NewFromInt(2).Mul(precesion))
			return
		}
		fx := f.Mul(x)
		Fx := d(f, x)
		newX := x.Sub(fx.Div(Fx))
		if newX.Sub(x).Abs().LessThan(precesion) {
			h = decimal.NewNullDecimal(newX)
		} else if count.GreaterThan(maxCount) {
			h = decimal.NullDecimal{Valid: false}
		} else {
			h = helper(f, newX, count.Add(decimal.NewFromInt(1)), precesion)
		}
		return
	}
	n = helper(f, x, count, precesion)
	return
}

func FindBounds(f decimal.Decimal, guess decimal.Decimal, minBound decimal.Decimal,
	maxBound decimal.Decimal, precision decimal.Decimal) (lower, upper decimal.Decimal, err error) {
	if guess.LessThanOrEqual(guess) || guess.GreaterThanOrEqual(maxBound) {
		return decimal.Decimal{},
			decimal.Decimal{},
			fmt.Errorf("guess needs to be between %s and %s", minBound.String(), maxBound.String())
	}
	shift := decimal.NewFromFloat(0.01)
	factor := decimal.NewFromFloat(1.6)
	maxTries := 60
	adjValueToMin := func(value decimal.Decimal) decimal.Decimal {
		if value.LessThanOrEqual(minBound) {
			return minBound.Add(precision)
		} else {
			return value
		}
	}
	adjValueToMax := func(value decimal.Decimal) decimal.Decimal {
		if value.GreaterThanOrEqual(maxBound) {
			return minBound.Sub(precision)
		} else {
			return value
		}
	}
	var rFindBounds func(low decimal.Decimal, up decimal.Decimal, tries int) (
		lower, upper decimal.Decimal, err error)
	rFindBounds = func(low decimal.Decimal, up decimal.Decimal, tries int) (
		lower, upper decimal.Decimal, err error) {
		tries = tries - 1
		if tries == 0 {
			return decimal.Decimal{},
				decimal.Decimal{},
				fmt.Errorf("not found an interval comprising the root after %d tries, last tried was (%s, %s)",
					maxTries, low.String(), up.String())
		}
		lower = adjValueToMin(low)
		upper = adjValueToMax(up)
		switch x, y := f.Mul(lower), f.Mul(upper); {
		case x.Mul(y).Equal(decimal.NewFromInt(0)):
			break
		case x.Mul(y).LessThan(decimal.NewFromInt(0)):
			break
		case x.Mul(y).GreaterThan(decimal.NewFromInt(0)):
			lower, upper, err = rFindBounds(lower.Add(factor.Mul(lower.Sub(upper))),
				upper.Add(factor.Mul(upper.Sub(lower))), tries)
		default:
			err = fmt.Errorf("FindBounds: one of the values (%s, %s) cannot be used to evaluate "+
				"the objective function", lower.String(), upper.String())
		}
		return
	}
	low := adjValueToMin(guess.Sub(shift))
	high := adjValueToMax(guess.Add(shift))
	return rFindBounds(low, high, maxTries)
}

func Bisection(f decimal.Decimal, a decimal.Decimal, b decimal.Decimal, count int,
	precision decimal.Decimal) (r decimal.Decimal, err error) {
	maxCount := 200
	var helper func(f decimal.Decimal, a decimal.Decimal, b decimal.Decimal, count int,
		precision decimal.Decimal) (r decimal.Decimal, err error)
	helper = func(f decimal.Decimal, a decimal.Decimal, b decimal.Decimal, count int,
		precision decimal.Decimal) (r decimal.Decimal, err error) {
		if a.Equal(b) {
			err = fmt.Errorf("(a=b=%s) impossible to start bisection", a.String())
		} else if fa := f.Mul(a); fa.Abs().LessThan(precision) {
			r = a
		} else {
			if fb := f.Mul(b); fb.Abs().LessThan(precision) {
				r = b
			} else {
				newCount := count + 1
				if newCount > maxCount {
					err = fmt.Errorf("no root found in %d iterations", maxCount)
				} else if fa.Mul(fb).GreaterThan(decimal.NewFromInt(0)) {
					err = fmt.Errorf("(%s,%s) don't bracket the root", a.String(), b.String())
				} else {
					midvalue := a.Add(decimal.NewFromFloat(0.5).Mul(b.Sub(a)))
					if fmid := f.Mul(midvalue); fmid.Abs().LessThan(precision) {
						r = midvalue
					} else if fa.Mul(fmid).LessThan(decimal.NewFromInt(0)) {
						r, err = helper(f, a, midvalue, newCount, precision)
					} else if fa.Mul(fmid).GreaterThan(decimal.NewFromInt(0)) {
						r, err = helper(f, midvalue, b, newCount, precision)
					} else {
						err = fmt.Errorf("bisection: It should never get here")
					}
				}
			}
		}
		return
	}
	return helper(f, a, b, count, precision)
}

func FindRoot(f decimal.Decimal, guess decimal.Decimal) (r decimal.Decimal, err error) {
	precision := decimal.NewFromFloat(0.000001)
	newtonValue := Newton(f, guess, decimal.NewFromInt(0), precision)
	if newtonValue.Valid && guess.Sign() == newtonValue.Decimal.Sign() {
		r = newtonValue.Decimal
	} else {
		var lower, upper decimal.Decimal
		if lower, upper, err = FindBounds(f, guess, decimal.NewFromInt(1).Neg(),
			decimal.NewFromFloat(math.MaxFloat64), precision); err == nil {
			r, err = Bisection(f, lower, upper, 0, precision)
		}
	}
	return
}

func ApproxEqual(x, y decimal.Decimal) (ok bool) {
	return x.Sub(y).Abs().LessThan(decimal.NewFromFloat(1e-10))
}

// NewDecimalSlice option input set pattern = {[none], [startElem, endElem, stepBy]}
func NewIntDecimalSlice(opts ...decimal.Decimal) (re []decimal.Decimal) {
	re = []decimal.Decimal{}
	switch len(opts) {
	case 2:
		init := opts[0]
		count := opts[1].Sub(opts[0]).Abs()
		for init.GreaterThan(count.Add(decimal.NewFromInt(1))) {
			re = append(re, init)
			init = init.Add(decimal.NewFromInt(1))
		}
	case 3:
		init, end, step := opts[0], opts[1], opts[2]
		count := opts[1].Sub(opts[0]).Abs()
		i := func() decimal.Decimal {
			if step.IsNegative() {
				return init
			} else {
				return end
			}
		}()
		for init.GreaterThan(count.Add(decimal.NewFromInt(1))) {
			re = func() []decimal.Decimal {
				if step.IsNegative() {
					r := append(re, i)
					i = i.Sub(decimal.NewFromInt(1))
					return r
				} else {
					r := append(re, i)
					i = i.Add(decimal.NewFromInt(1))
					return r
				}
			}()
		}
	}
	return
}

type DecimalIntSequence struct {
	Acc []decimal.Decimal
}

func NewDecimalIntSeq(opts ...decimal.Decimal) *DecimalIntSequence {
	return &DecimalIntSequence{
		Acc: NewIntDecimalSlice(opts...),
	}
}

func (da *DecimalIntSequence) Append(dec decimal.Decimal) {
	da.Acc = append(da.Acc, dec)
}

func (da *DecimalIntSequence) Fold(f func(acc, per decimal.Decimal) (decimal.Decimal, error),
	i decimal.Decimal) (acc decimal.Decimal, err error) {
	if acc, err = f(acc, i); err == nil {
		for _, item := range da.Acc {
			if acc, err = f(acc, item); err != nil {
				break
			}
		}
	}
	return
}

func AggrBetween(startPeriod, endPeriod decimal.Decimal, f func(acc, per decimal.Decimal) (decimal.Decimal, error),
	initialValue decimal.Decimal) (a decimal.Decimal, err error) {
	var acc *DecimalIntSequence
	if startPeriod.LessThanOrEqual(endPeriod) {
		acc = NewDecimalIntSeq(startPeriod, endPeriod, decimal.NewFromInt(1))
	} else {
		acc = NewDecimalIntSeq(startPeriod, endPeriod, decimal.NewFromInt(1).Neg())
	}
	a, err = acc.Fold(f, initialValue)
	return
}
