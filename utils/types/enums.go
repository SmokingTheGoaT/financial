package types

import (
	"fmt"
	"github.com/shopspring/decimal"
)

var (
	timeUnits = [...]string{
		"monthly",
		"quarterly",
		"yearly",
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

func (tu Term) Term(value int32) Period {
	v := decimal.NewFromInt32(value)
	return Period{v, tu}
}

type Period struct {
	amount decimal.Decimal
	Term
}

func (t Period) Amount() decimal.Decimal {
	return t.amount
}

func (t Period) String() string {
	return fmt.Sprintf("%d %s", t.amount, t.Term.String())
}

type PaymentDue int

const (
	EndOfPeriod PaymentDue = iota
	BeginningOfPeriod
)

func (pd PaymentDue) String() string {
	return [...]string{"EndOfPeriod", "BeginningOfPeriod"}[pd]
}
