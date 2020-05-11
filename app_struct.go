package main

import "fmt"

type product struct {
	productName  string `json:product_name`
	productPrice int    `json:product_price`
	description  string `json:description`
}

func main() {
	// productA := product{
	// 	productName:  "kecap",
	// 	productPrice: 10000,
	// 	description:  "kecap 100 ml",
	// }

	// fmt.Println(productA)
	products := []product{
		{productName: "kecap", productPrice: 1000, description: "kecap 100ml"},
		{productName: "garam", productPrice: 5000, description: "garam 200g"},
	}

	fmt.Println(products[0])
	fmt.Println(products[1])

	for i, prod := range products {
		fmt.Println(i, prod)
	}
}
