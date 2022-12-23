package pmt

import (
	"financial/enum"
	"financial/money"
	"financial/percent"
)

type (
	pmtOpt func(pmt *pmt) (amount money.Amount)

	pmt struct {
		loanAmount           money.Amount
		periodicInterestRate percent.Percent
		paymentPeriods       enum.Period
	}
)

func New(fnOpts ...pmtOpt) (amount money.Amount) {
	obj := new(pmt)
	amount = fnOpts[0](obj)
	return
}

func Amortized(a money.Amount, pir percent.Percent, pn enum.Period) pmtOpt {
	return func(pmt *pmt) (amount money.Amount) {
		//amount.Amount = currency.USD.Amount()
		return
	}
}
