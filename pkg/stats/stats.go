package stats

import (
	"github.com/kobilov-web-dev/bank/v2/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	count := 0
	for _, payment := range payments {
		if payment.Status != "FAIL" {
			sum += payment.Amount
			count++
		}
	}
	return sum / types.Money(count)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for _, payment := range payments {
		if payment.Category == category {
			if payment.Status != "FAIL" {
				sum += payment.Amount
			}
		}
	}
	return sum
}
