package stats

import (
	"testing"
	"reflect"
    "github.com/Oleg196508/bank/v2/pkg/types")

func TestFilterByCategory_nil(t *testing.T) {
	var payments []types.Payment
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")	
	}	
}
func TestFilterByCategory_empty(t *testing.T) {
	var payments []types.Payment
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")	
	}	
}
func TestFilterByCategory_notFound(t *testing.T) {
	payments := []types.Payment {
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun" },
	}
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")	
	}	
}

func TestFilterByCategory_foundOne(t *testing.T) {
	payments := []types.Payment {
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun" },
	}
	expected := []types.Payment {
		{ID: 2, Category: "food"},	
	}
	result := FilterByCategory(payments, "food") 

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)	
	}	
}

func TestFilterByCategory_foundMultiple(t *testing.T) {
	payments := []types.Payment {
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun" },
	}
	expected := []types.Payment {
		{ID: 1, Category: "auto"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
	}
	result := FilterByCategory(payments, "auto") 

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)	
	}	
}
func TestCategoriesTotal(t *testing.T) {
	payments := []types.Payment {
		{ID: 1, Category: "auto", Amount: 1000_00},
		{ID: 2, Category: "food", Amount:  2000_00},
		{ID: 3, Category: "auto", Amount:  3000_00},
		{ID: 4, Category: "auto", Amount:  4000_00},
		{ID: 5, Category: "fun",  Amount:  5000_00},
	}
	expected := map[types.Category]types.Money {
		"auto": 8000_00,
		"food": 2000_00,
		"fun": 5000_00,
		
	}
	result := CategoriesTotal(payments) 

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)	
	}	
}
func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment {
		{ID: 1, Category: "auto", Amount:  1000_00},
		{ID: 2, Category: "food", Amount:  2000_00},
		{ID: 3, Category: "auto", Amount:  3000_00},
		{ID: 4, Category: "auto", Amount:  4000_00},
		{ID: 5, Category: "fun",  Amount:  5000_00},
		{ID: 6, Category: "fun",  Amount:  3000_00},
		{ID: 7, Category: "auto",  Amount: 2000_00},
	}
	expected := map[types.Category]types.Money {
		"auto": 2500_00,
		"food": 2000_00,
		"fun":  4000_00,
		
	}
	result := CategoriesAvg(payments) 

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)	
	}	
}
func TestPeriodsDynamic(t *testing.T) {
	first := map[types.Category]types.Money {
	    "auto": 10_000_00,
		"food": 20_000_00,
		"phone": 9_000_00,	
	}
	second := map[types.Category]types.Money {
	    "auto": 5000_00,
		"food": 3000_00,
		"mobile": 5000_00,

	}
	
	expected := map[types.Category]types.Money {
		"auto": -5000_00,
		"food": -17_000_00,
		"mobile": 5000_00,
		"phone": -9_000_00,	
	}
	result := PeriodsDynamic(first, second) 

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)	
	}	
}

