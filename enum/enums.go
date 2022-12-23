package enum

import (
	"fmt"
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
	return &period{value, tu}
}

type period struct {
	amount int32
	Term
}

type Period interface {
	Amount() int32
	String() string
}

func (t *period) Amount() int32 {
	return t.amount
}

func (t *period) String() string {
	return fmt.Sprintf("%d %s", t.amount, t.Term.String())
}
