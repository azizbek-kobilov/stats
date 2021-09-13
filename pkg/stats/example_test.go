package stats

import (
	"fmt"
	"github.com/kobilov-web-dev/bank/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			Amount: 10_000,
		},
		{
			Amount: 20_000,
		},
		{
			Amount: 30_000,
		},
	}

	fmt.Println(Avg(payments))
	// Output: 20000
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			Amount:   10_000,
			Category: "food",
		},
		{
			Amount:   20_000,
			Category: "food",
		},
		{
			Amount:   30_000,
			Category: "car",
		},
	}

	fmt.Println(TotalInCategory(payments, "food"))
	// Output: 30000
}
