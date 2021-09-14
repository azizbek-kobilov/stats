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

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categoriesSumm := map[types.Category]types.Money{}
	categoriesCount := map[types.Category]int{}

	for _, payment := range payments {
		categoriesSumm[payment.Category] += payment.Amount
		categoriesCount[payment.Category] += 1
	}

	categoriesAvg := map[types.Category]types.Money{}

	for key, value := range categoriesSumm {
		categoriesAvg[key] += value / types.Money(categoriesCount[key])
	}

	return categoriesAvg
}

func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money,) map[types.Category]types.Money {
	diff := map[types.Category]types.Money{}

	for key, value := range first {
		diff[key] = second[key] - value
	}

	for key := range second {
		if _, ok := diff[key]; !ok {
			diff[key] = second[key]
		}
	}

	return diff
}
