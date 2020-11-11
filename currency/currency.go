package currency

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Currency create the struc that will contains the JSON value
type Currency struct {
	ExhangeValue float64 `json:"exchangeValue"`
}

// GetCurrency get the current EURO value
func GetCurrency(u string) (f float64) {

	res, err := http.Get(u)

	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var euro Currency

	if err := json.Unmarshal([]byte(b), &euro); err != nil {
		log.Fatal(err)
	}

	return euro.ExhangeValue

}
