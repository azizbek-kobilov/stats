package stats

import (
	"reflect"
	"testing"
	"github.com/kobilov-web-dev/bank/v2/pkg/types"
)

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 1_000_00},
		{ID: 2, Category: "food", Amount: 2_000_00},
		{ID: 3, Category: "auto", Amount: 3_000_00},
		{ID: 4, Category: "auto", Amount: 4_000_00},
		{ID: 5, Category: "fun", Amount: 5_000_00},
	}
	expected := map[types.Category]types.Money{
		"auto": 2_666_66,
		"food": 2_000_00,
		"fun": 5_000_00,
	}
	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestPeriodsDynamic(t *testing.T) {
	firstPeriod := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	secondPeriod := map[types.Category]types.Money{
		"auto": 10,
		"food": 25,
		"mobile": 5,
	}
	expected := map[types.Category]types.Money{
		"auto": 0,
		"food": 5,
		"mobile": 5,
	}
	result := PeriodsDynamic(firstPeriod, secondPeriod)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}
