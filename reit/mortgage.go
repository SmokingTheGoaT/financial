package reit

import (
	"financial/utils/currency"
	"financial/utils/percent"
	"financial/utils/types"
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
	//res.MonthlyMortgage = loan.New(
	//	loan.Amortized(
	//		request.PropertySalePrice,
	//		request.InterestRate,
	//		request.Term,
	//	),
	//)
	return
}
