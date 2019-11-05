package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/bold-commerce/go-shopify"
)

var apiKey = flag.String("k", "", "Api Key")
var apiPass = flag.String("p", "", "Api Password")

func main() {
	flag.Parse()

	app := goshopify.App{
		ApiKey:   *apiKey,
		Password: *apiPass,
	}

	// Create a new API client
	client := goshopify.NewClient(app, "kburns-test-store", "")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/html")
		// Fetch the number of products.
		numProducts, err := client.Product.Count(nil)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		fmt.Fprintf(w, "<p>Total Products: %d</p>", numProducts)

		// Fetch the number of products.
		products, err := client.Product.List(nil)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		for _, p := range products {
			fmt.Fprintf(w, "<p style='background: #e9e9e9; padding: 1em; '>%d: %s\n<br/>", p.ID, p.Title)
			fmt.Fprintf(w, "<img src='%s' width='200'/></p>\n", p.Image.Src)
		}

		fmt.Fprint(w, `<p>
			<a href='https://kburns-test-store.myshopify.com/' target='_blank'>https://kburns-test-store.myshopify.com/</a><br>
			<a href='https://github.com/kevburnsjr/shorpify' target='_blank'>https://github.com/kevburnsjr/shorpify</a>
		</p>`)
	})
	log.Fatal(http.ListenAndServe(":8060", nil))

}
