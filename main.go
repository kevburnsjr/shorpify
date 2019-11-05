package main

import (
	"fmt"
	"flag"

	"github.com/bold-commerce/go-shopify"
)

var apiKey = flag.String("k", "", "Api Key")
var apiPass = flag.String("p", "", "Api Password")

func main() {
	flag.Parse()

	app := goshopify.App{
		ApiKey: *apiKey,
		Password: *apiPass,
	}

	// Create a new API client
	client := goshopify.NewClient(app, "kburns-test-store", "")

	// Fetch the number of products.
	numProducts, err := client.Product.Count(nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(numProducts)

}
