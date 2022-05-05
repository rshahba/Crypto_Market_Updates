package runAPI

import (
	"Crypto_Market_Updates/client"
	"flag"
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

	urlK, err := client.FiatCrypto(*currencyCode, *cryptoCode)

	if err != nil {
		log.Println(err)
	}

	APIFetchRes := &client.APIURL{APILink: urlK}
	client.PrintOutline(APIFetchRes)
}
