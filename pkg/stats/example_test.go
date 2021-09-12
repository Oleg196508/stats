package stats

import (
	"fmt"
	"github.com/Oleg196508/bank/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:     1,
	        Amount: 10,
			Category: "",
		},
		{
			ID:     99,
	        Amount: 60,
			Category: "",
		},
		{
			ID: 3,
			Amount: 20,
			Category: "",
		}, 
	}
	avgsum := Avg(payments)
	fmt.Println(avgsum)

	//Output: 30	
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{ ID:     1, Amount: 10, Category: "avto",   },
		{ ID:     2, Amount: 20, Category: "avto",   },
		{ ID:     3, Amount: 30, Category: "phone",  },
		{ ID:     4, Amount: 40, Category: "phone",  },
		{ ID:     5, Amount: 50, Category: "phone",  },
		{ ID:     6, Amount: 60, Category: "school", },	
	}

	fmt.Println(TotalInCategory(payments, "phone"))
	
	//Output: 120	
}