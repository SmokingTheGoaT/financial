package fin

import (
	"financial/types"
	"financial/utils/percent"
	"github.com/shopspring/decimal"
)

func ipmt(r percent.Percent, per decimal.Decimal, nper types.Period,
	pv, fv decimal.Decimal, pd types.PaymentDue) (ip decimal.Decimal) {

	return
}
