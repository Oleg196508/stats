package stats

import (
	"fmt"
	"github.com/Oleg196508/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:     1,
	        Amount: 10,
			Category: "",
			Status: "OK",
		},
		{
			ID:     99,
	        Amount: 60,
			Category: "",
			Status: "OK",
		},
		{
			ID: 3,
			Amount: 20,
			Category: "",
			Status: "OK",
		}, 
		{
			ID: 3,
			Amount: 20,
			Category: "",
			Status: "FAIL",
		}, 
	}
	avgsum := Avg(payments)
	fmt.Println(avgsum)

	//Output: 30	
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{ ID:     1, Amount: 10, Category: "avto",   Status: "FAIL",  },
		{ ID:     2, Amount: 20, Category: "avto",   Status: "OK"     },
		{ ID:     3, Amount: 30, Category: "phone",  Status: "OK"     },
		{ ID:     4, Amount: 40, Category: "phone",  Status: "OK"     },
		{ ID:     5, Amount: 50, Category: "phone",  Status: "OK"     },
		{ ID:     6, Amount: 60, Category: "school", Status: "OK"     },	
	}

	fmt.Println(TotalInCategory(payments, "phone"))
	
	//Output: 120	
}