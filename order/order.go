package order

import (
	"fmt"

	fooditem "github.com/saint-rivers/nhamey-cli/fooditem"
)

type Order struct {
	Id        string              `json:"id"`
	Location  string              `json:"location"`
	FoodItems []fooditem.FoodItem `json:"foodItems"`
}

func (order *Order) Display() {
	fmt.Print(order)
}
