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

	fmt.Println(result.Products[0].DisplayName)
	fmt.Println(result.Products[0].Price)

}
