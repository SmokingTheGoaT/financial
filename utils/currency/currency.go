package currency

import (
	"financial/utils"
	"fmt"
	"github.com/shopspring/decimal"
	"golang.org/x/text/currency"
)

type Amount struct {
	amount decimal.Decimal
	currency.Amount
}

func New(value interface{}, currencyUnit currency.Unit, opts ...interface{}) Amount {
	v := utils.ParseFloat(value, opts...)
	return Amount{
		decimal.NewFromFloat(v), currencyUnit.Amount(v),
	}
}

func (a Amount) Decimal() (d decimal.Decimal) {
	return a.amount
}

func (a Amount) String() string {
	return fmt.Sprintf("%s %s", a.Amount.Currency().String(), a.amount.String())
}

func (a Amount) Add(a2 Amount) (c Amount) {
	t := a.amount.Add(a2.Decimal())
	return Amount{
		t,
		a.Amount.Currency().Amount(t.String()),
	}
}

func (a Amount) Sub(a2 Amount) (c Amount) {
	t := a.amount.Sub(a2.Decimal())
	return Amount{
		t,
		a.Amount.Currency().Amount(t.String()),
	}
}

func (a Amount) Mul(a2 Amount) (c Amount) {
	t := a.amount.Mul(a2.Decimal())
	return Amount{
		t,
		a.Amount.Currency().Amount(t.String()),
	}
}

func (a Amount) Div(a2 Amount) (c Amount) {
	t := a.amount.Div(a2.Decimal())
	return Amount{
		t,
		a.Amount.Currency().Amount(t.String()),
	}
}
