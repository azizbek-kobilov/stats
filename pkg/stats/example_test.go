package stats

import (
	"fmt"
	"github.com/kobilov-web-dev/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			Amount: 10_000,
			Status: "FAIL",
		},
		{
			Amount: 20_000,
			Status: "OK",
		},
		{
			Amount: 30_000,
			Status: "INPROGRESS",
		},
	}

	fmt.Println(Avg(payments))
	// Output: 25000
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			Amount:   10_000,
			Category: "food",
			Status: "FAIL",
		},
		{
			Amount:   20_000,
			Category: "food",
			Status: "OK",
		},
		{
			Amount:   30_000,
			Category: "car",
			Status: "INPROGRESS",
		},
	}

	fmt.Println(TotalInCategory(payments, "food"))
	// Output: 20000
}
