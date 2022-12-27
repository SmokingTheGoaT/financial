package financial

import (
	_ "financial/fin"
	"financial/types"
	"financial/utils/percent"
	"github.com/shopspring/decimal"
)

// PV is the present value of an investment.
// href=https://support.microsoft.com/en-us/office/pv-function-23879d31-0e02-4321-be01-da16e8168cbd
func PV(rate percent.Percent, nper types.Period, pmt decimal.Decimal, fv decimal.Decimal, pd types.PaymentDue) (
	res decimal.Decimal, err error)

// FV is the future value of an investment
// href=https://support.microsoft.com/en-us/office/fv-function-2eef9f44-a084-4c61-bdd8-4fe4bb1b71b3
func FV(rate percent.Percent, nper types.Period, pmt decimal.Decimal, pv decimal.Decimal, pd types.PaymentDue) (
	res decimal.Decimal, err error)

// PMT is the periodic payment for an annuity
// href=https://support.microsoft.com/en-us/office/pmt-function-0214da64-9a63-4996-bc20-214433fa6441
func PMT(rate percent.Percent, nper types.Period, pv decimal.Decimal, fv decimal.Decimal, pd types.PaymentDue) (
	res decimal.Decimal, err error)

// RRI returns an equivalent interest rate for the growth of an investment
// href=https://support.microsoft.com/en-us/office/rri-function-6f5822d8-7ef1-4233-944c-79e8172930f4
func RRI(nper types.Period, pv decimal.Decimal, fv decimal.Decimal) (res decimal.Decimal, err error)

// NPER is the number of periods for an investment
// href=https://support.microsoft.com/en-us/office/nper-function-240535b5-6653-4d2d-bfcf-b6a38151d815
func NPER(rate percent.Percent, pmt decimal.Decimal, pv decimal.Decimal, fv decimal.Decimal, pd types.PaymentDue) (
	res decimal.Decimal, err error)

// RATE is the interest rate per period of an annuity
// href=https://support.microsoft.com/en-us/office/rate-function-9f665657-4a7e-4bb7-a030-83fc59e748ce
func RATE(nper types.Period, pmt decimal.Decimal, pv decimal.Decimal, fv decimal.Decimal,
	pd types.PaymentDue, guess decimal.Decimal) (res decimal.Decimal, err error)

// FVSchedule is the future value of an initial principal after applying a series of compound interest rates
// href=https://support.microsoft.com/en-us/office/fvschedule-function-bec29522-bd87-4082-bab9-a241f3fb251d
func FVSchedule(pv decimal.Decimal, interests []decimal.Decimal) (res decimal.Decimal, err error)

// PDuration returns the number of periods required by an investment to reach a specified value.
// href=https://support.microsoft.com/en-us/office/pduration-function-44f33460-5be5-4c90-b857-22308892adaf
func PDuration(rate percent.Percent, pv, fv decimal.Decimal) (res decimal.Decimal, err error)
