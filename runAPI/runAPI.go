package runAPI

import (
	"Crypto_Market_Updates/client"
	"Crypto_Market_Updates/model"
	"flag"
	"fmt"
	"log"

)

func TextOutput(nPs model.NomicsResponse) string {
	
	if nPs.AllTimeHigh == "" {
		log.Fatal("Error in currency or crypto codes. Please check for correct spelling.")
	}

	p := fmt.Sprintf(
		"\n---------------------\nName: %s\nCurrent Price : $%s\nMarket Cap Rank: %s\nAll Time High: $%s\nCirculating Supply: %s\nNumber of Exchanges Traded: %s\n",
		nPs.Name, nPs.CurrentPrice, nPs.MarketCapRank, nPs.AllTimeHigh, nPs.CirculatingSupply, nPs.NumExchangesTraded)
	return p
}

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
	CryptoS, errr := client.GetURL(urlK)
	if errr != nil {
		log.Println(err)
	}
	fmt.Printf("%#v", CryptoS)
	fmt.Println(TextOutput(CryptoS))
}
