package reit

import (
	"financial/enum"
	"financial/money"
	"financial/percent"
	"financial/pmt"
)

type (
	MortgageRequest struct {
		PropertySalePrice money.Amount
		LoanToValueRatio  percent.Percent
		DownPayment       money.Amount
		ClosingCosts      money.Amount
		Principal         money.Amount
		InterestRate      percent.Percent
		Term              enum.Period
	}

	MortgageResponse struct {
		MonthlyMortgage money.Amount
		MonthlyNet      money.Amount
		AnnualizedNet   money.Amount
		AnnualizedROI   percent.Percent
	}
)

func Mortgage(request MortgageRequest) (res MortgageResponse) {
	res = MortgageResponse{}
	res.MonthlyMortgage = pmt.New(pmt.Amortized(
		request.PropertySalePrice, request.InterestRate, request.Term))
	return
}
