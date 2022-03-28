package fooditem

import (
	"fmt"

	user "github.com/saint-rivers/nhamey-cli/user"
)

type FoodItem struct {
	Id      string    `json:"id"`
	Name    string    `json:"name"`
	Amount  float32   `json:"amount"`
	Ordered user.User `json:"orderedBy"`
}

func (item *FoodItem) display() {
	fmt.Print(item)
}
