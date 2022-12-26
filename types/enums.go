package types

import (
	"fmt"
	"github.com/shopspring/decimal"
)

var (
	timeUnits = [...]string{
		"month(s)",
		"quarter(s)",
		"year(s)",
	}
)

type Term uint16

const (
	Monthly Term = iota
	Quarterly
	Yearly
)

func (tu Term) String() string {
	return timeUnits[tu]
}

func (tu Term) Term(value uint64) Period {
	v := decimal.NewFromInt(int64(value))
	return Period{v, tu}
}

type periodOpt func(p *Period)

type Period struct {
	amount decimal.Decimal
	Term
}

func ToMonths() periodOpt {
	return func(p *Period) {
		switch p.Term.String() {
		case Monthly.String():
			break
		case Quarterly.String():
			m := decimal.NewFromInt(3)
			p.amount = p.amount.Mul(m)
			p.Term = Monthly
		case Yearly.String():
			m := decimal.NewFromInt(12)
			p.amount = p.amount.Mul(m)
			p.Term = Monthly
		}
	}
}

func ToQuarters() periodOpt {
	return func(p *Period) {
		switch p.Term.String() {
		case Monthly.String():
			m := decimal.NewFromInt(3)
			p.amount = p.amount.Div(m)
			p.Term = Quarterly
		case Quarterly.String():
			break
		case Yearly.String():
			m := decimal.NewFromInt(4)
			p.amount = p.amount.Mul(m)
			p.Term = Quarterly
		}
	}
}

func ToYears() periodOpt {
	return func(p *Period) {
		switch p.Term.String() {
		case Monthly.String():
			m := decimal.NewFromInt(12)
			p.amount = p.amount.Div(m)
			p.Term = Yearly
		case Quarterly.String():
			m := decimal.NewFromInt(4)
			p.amount = p.amount.Div(m)
			p.Term = Yearly
		case Yearly.String():
			break
		}
	}
}

func (t Period) Convert(opts ...periodOpt) Period {
	if len(opts) > 0 && len(opts) < 2 {
		for _, o := range opts {
			o(&t)
		}
	} else {
		panic(fmt.Errorf("only one periodOpt can be included, if period conversion is desired"))
	}
	return t
}

func (t Period) Amount() decimal.Decimal {
	return t.amount
}

func (t Period) String() string {
	return fmt.Sprintf("%s %s", t.amount.String(), t.Term.String())
}

type PaymentDue int

const (
	EndOfPeriod PaymentDue = iota
	BeginningOfPeriod
)

func (pd PaymentDue) String() string {
	return [...]string{"EndOfPeriod", "BeginningOfPeriod"}[pd]
}

func (pd PaymentDue) CompareTo(pd2 PaymentDue) bool {
	return pd == pd2
}
