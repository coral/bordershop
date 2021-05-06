package main

import (
	"fmt"

	"github.com/coral/bordershop"
)

func main() {
	result, err := bordershop.GetCategory(9817)
	if err != nil {
		panic(err)
	}

	for _, product := range result.Products {

		fmt.Printf("%+v\n", product)
		break;
	}


}
