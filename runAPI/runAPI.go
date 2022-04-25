package runAPI

import (
	"Crypto_Market_Updates/client"
	"flag"
	"fmt"
	"log"
)

func RunApp() {

	currencyCode := flag.String(
		"currency", "CAD", "The code of the currency you would like to know the price of your crypto in",
	)

	cryptoCode := flag.String(
		"crypto", "BTC", "Input the crypto code you would like to know the price of",
	)
	flag.Parse()

	crypto, err := client.APIFetch(*currencyCode, *cryptoCode)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(crypto)
}
