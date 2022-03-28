package main

import (
	"fmt"

	"github.com/saint-rivers/nhamey-cli/user"
)

func main() {
	fmt.Print("hello world\n")
	user := user.User{}
	fmt.Print(user.Firstname)

	// newOrder()
}

func newOrder() {
	// order := new(order.Order){"001", "intratevy", nil}
	// order.Display()
}
