package main
// Pointers allows us to use things by reference.You do not get copy but get reference.
// 

import (
	"fmt"
)

var shippingDays = 30

func main() {
	shippingDaysAdjustments(shippingDays)
	fmt.Println("after shippingDaysAdjustments calls", shippingDays)

	shippingDaysAdjustmentsPtr(&shippingDays)
	fmt.Println("after shippingDaysAdjustments calls", shippingDays)
	fmt.Println("after shippingDaysAdjustments calls", &shippingDays)
}

func shippingDaysAdjustments(shippingDays int) {
	shippingDays += 10
}

func shippingDaysAdjustmentsPtr(shippingDays *int) {
	*shippingDays += 10
}