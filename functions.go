package financial

import (
	_ "financial/fin"
	"financial/types"
	"github.com/shopspring/decimal"
)

// PV is the present value of an investment.
// href=https://support.microsoft.com/en-us/office/pv-function-23879d31-0e02-4321-be01-da16e8168cbd
func PV(rate types.Rate, nper, pmt, fv decimal.Decimal, pd types.PaymentDue) (res decimal.Decimal, err error)

// FV is the future value of an investment
// href=https://support.microsoft.com/en-us/office/fv-function-2eef9f44-a084-4c61-bdd8-4fe4bb1b71b3
func FV(rate types.Rate, nper, pmt, pv decimal.Decimal, pd types.PaymentDue) (res decimal.Decimal, err error)

// PMT is the periodic payment for an annuity
// href=https://support.microsoft.com/en-us/office/pmt-function-0214da64-9a63-4996-bc20-214433fa6441
func PMT(rate types.Rate, nper, pv, fv decimal.Decimal, pd types.PaymentDue) (res decimal.Decimal, err error)

// RRI returns an equivalent interest rate for the growth of an investment
// href=https://support.microsoft.com/en-us/office/rri-function-6f5822d8-7ef1-4233-944c-79e8172930f4
func RRI(nper, pv, fv decimal.Decimal) (res decimal.Decimal, err error)

// NPER is the number of periods for an investment
// href=https://support.microsoft.com/en-us/office/nper-function-240535b5-6653-4d2d-bfcf-b6a38151d815
func NPER(rate types.Rate, pmt, pv, fv decimal.Decimal, pd types.PaymentDue) (res decimal.Decimal, err error)

// RATE is the interest rate per period of an annuity
// href=https://support.microsoft.com/en-us/office/rate-function-9f665657-4a7e-4bb7-a030-83fc59e748ce
func RATE(nper, pmt, pv, fv decimal.Decimal, pd types.PaymentDue, opts ...decimal.Decimal) (
	res decimal.Decimal, err error)

// FVSchedule is the future value of an initial principal after applying a series of compound interest rates
// href=https://support.microsoft.com/en-us/office/fvschedule-function-bec29522-bd87-4082-bab9-a241f3fb251d
func FVSCHEDULE(pv decimal.Decimal, interests []decimal.Decimal) (res decimal.Decimal, err error)

// PDuration returns the number of periods required by an investment to reach a specified value.
// href=https://support.microsoft.com/en-us/office/pduration-function-44f33460-5be5-4c90-b857-22308892adaf
func PDURATION(rate types.Rate, pv, fv decimal.Decimal) (res decimal.Decimal, err error)

// IPMT is the interest payment for an investment for a given period
// href=https://support.microsoft.com/en-us/office/ipmt-function-5cce0ad6-8402-4a41-8d29-61a0b054cb6f
func IPMT(rate types.Rate, per, nper, pv, fv decimal.Decimal, pd types.PaymentDue) (res decimal.Decimal, err error)

// PPMT is the payment on the principal for an investment for a given period
// href=https://support.microsoft.com/en-us/office/ppmt-function-c370d9e3-7749-4ca4-beea-b06c6ac95e1b
func PPMT(rate types.Rate, per, nper, pv, fv decimal.Decimal, pd types.PaymentDue) (res decimal.Decimal, err error)

// CUMIPMT is the cumulative interest paid between two periods
// href=https://support.microsoft.com/en-us/office/cumipmt-function-61067bb0-9016-427d-b95b-1a752af0e606
func CUMIPMT(rate types.Rate, nper, pv, startPeriod, endPeriod decimal.Decimal, pd types.PaymentDue) (
	res decimal.Decimal, err error)

// CUMPRINC is the cumulative principal paid on a loan between two periods
// href=https://support.microsoft.com/en-us/office/cumprinc-function-94a4516d-bd65-41a1-bc16-053a6af4c04d
func CUMPRINC(rate types.Rate, nper, pv, startPeriod, endPeriod decimal.Decimal, pd types.PaymentDue) (
	res decimal.Decimal, err error)

// ISPMT calculates the interest paid during a specific period of an investment
// href=https://support.microsoft.com/en-us/office/ispmt-function-fa58adb6-9d39-4ce0-8f43-75399cea56cc
func ISPMT(rate types.Rate, per, nper, pv decimal.Decimal) (re decimal.Decimal, err error)
