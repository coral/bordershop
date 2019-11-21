package main

import (
	"fmt"

	"github.com/coral/bordershop"
)

func main() {
	result, err := bordershop.GetCategory(9819)
	if err != nil {
		panic(err)
	}

	for _, product := range result.Products {
		fmt.Println("-----   PRODUKT    ------")
		fmt.Println(product.DisplayName)
		fmt.Println(product.Price.AmountAsDecimal)
		fmt.Println(product.Image)
		fmt.Println()
	}

}
