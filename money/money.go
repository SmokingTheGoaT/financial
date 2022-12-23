package money

import (
	"github.com/shopspring/decimal"
	"golang.org/x/text/currency"
)

type Amount struct {
	currency.Amount
}

func (a Amount) Decimal() (d decimal.Decimal) {

	return
}

func (a Amount) Add() {

}

func (a Amount) Subtract() {

}

func (a Amount) Multiply() {

}

func (a Amount) Divide() {
	
}