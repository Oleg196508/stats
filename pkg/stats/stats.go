package stats

import 
"github.com/Oleg196508/bank/pkg/types"

//Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
    sum := types.Money(0)
	number := 0
	for _, payment := range payments {
		number++
		sum += payment.Amount
	    
	}
	avgsum := sum / types.Money(number)

	return avgsum
}

//TotalInCategory находит сумму покупок в определённой категории
func TotalInCategory(payments []types.Payment, category types.Category)types.Money  {

	sum := types.Money(0)
	number := 0

	for _, payment := range payments {
		if payment.Category == category {
			number++
			sum += payment.Amount
		}	
	}
	return sum
}