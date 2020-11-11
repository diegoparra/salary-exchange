package main

import (
	"fmt"
	"nf-generator/currency"
	"nf-generator/generate"
)

// Yahoo
//var url string = "https://go-swap-server.herokuapp.com/convert?from=EUR&to=BRL&amount=100&exchanger=yahoo&cacheTime=60s"

// TheMoneyConverter
const url string = "https://go-swap-server.herokuapp.com/convert?from=EUR&to=BRL&amount=100&exchanger=themoneyconverter&cacheTime=60s"

func main() {

	monthNumber := make(map[string]int)
	monthNumber["October"] = 10
	monthNumber["November"] = 11

	workdays := make(map[string]int)
	workdays["October"] = 22
	workdays["November"] = 21

	euroValue := currency.GetCurrency(url)

	fmt.Println("Cotação do Euro:", euroValue)

	generate.Generate(workdays, monthNumber, euroValue)
}
