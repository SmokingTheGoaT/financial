package reit

import (
	"financial/types"
	"financial/utils/currency"
	"financial/utils/percent"
)

type (
	MortgageRequest struct {
		PropertySalePrice currency.Amount
		LoanToValueRatio  percent.Percent
		DownPayment       currency.Amount
		ClosingCosts      currency.Amount
		Principal         currency.Amount
		InterestRate      percent.Percent
		Term              types.Period
	}

	MortgageResponse struct {
		MonthlyMortgage currency.Amount
		MonthlyNet      currency.Amount
		AnnualizedNet   currency.Amount
		AnnualizedROI   percent.Percent
	}
)

func Mortgage(request MortgageRequest) (res MortgageResponse) {
	res = MortgageResponse{}
	//res.MonthlyMortgage = tvm.PMT()
	return
}
