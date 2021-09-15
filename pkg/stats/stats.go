package stats

import 
"github.com/Oleg196508/bank/v2/pkg/types"

//Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
    sum := types.Money(0)
	number := 0
	for _, payment := range payments {
		if payment.Status != "FAIL" {
			number++
			sum += payment.Amount	
		}   
	}
	avgsum := sum / types.Money(number)

	return avgsum
}

//TotalInCategory находит сумму покупок в определённой категории
func TotalInCategory(payments []types.Payment, category types.Category)types.Money  {

	sum := types.Money(0)
	number := 0

	for _, payment := range payments {
		if payment.Category == category && payment.Status != "FAIL" {
			number++
			sum += payment.Amount
		}	
	}
	return sum
}
//FilterByCategory возвращает платежи в указанной категории
func FilterByCategory(payments []types.Payment, category types.Category)[]types.Payment  {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}	
	}
	return filtered
}
//CategoriesTotal возвращает сумму платежей по каждой категории
func CategoriesTotal(payments []types.Payment)map[types.Category]types.Money  {
	categories := map[types.Category]types.Money {}

	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
	}
	return categories
}

//CategoriesAvg считает среднюю сумму платежа по каждой категории
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money  {

	categories := map[types.Category]types.Money{}
    number := map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category] += payment.Amount //получили категория:платёж, категория:платёж...
		number[payment.Category]++ //получили категория:сколько раз складывали,...	
	}      
	for key, value := range categories {
		categories[key] = value / number[key] //В принципе value можно убрать
	 }
	return categories
}


