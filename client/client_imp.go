package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type NomicsResponse []struct {
	Name               string `json:"name"`
	CurrentPrice       string `json:"price"`
	MarketCapRank      string `json:"rank"`
	AllTimeHigh        string `json:"high"`
	CirculatingSupply  string `json:"circulating_supply"`
	NumExchangesTraded string `json:"num_exchanges"`
}

func FiatCrypto(currency string, crypto string) (string, error) {

	fmt.Print("Enter currency: ")
	fmt.Scanf("%s", &currency)
	fmt.Print("Enter crypto: ")
	fmt.Scanf("%s", &crypto)

	URL := "https://api.nomics.com/v1/currencies/ticker?key=4fe2103af29f0c4acbb7a2ef6a7b9015c0b70c9a&interval=1d&ids=" + crypto + "&convert=" + currency

	switch {

	case len(currency) != 3 && len(crypto) != 3:
		log.Fatal("Currency and crypto codes can have 3 characters only (ex. -currencey = CAD , -crypto = BTC).")

	case len(crypto) < 3 || len(crypto) > 3:
		fmt.Println("Crypto code can have 3 characters only (ex. BTC).")
		os.Exit(1)

	case len(currency) < 3 || len(currency) > 3:
		log.Fatal("Currency code can have 3 characters only (ex. CAD).")
	}
	return URL, nil
}

type APIURL struct {
	APILink string
}

func (a *APIURL) GetUrlStr() (string, error) {
	url := a.APILink
	//Get function
	response, err := http.Get(url)

	//Error handling
	if err != nil {
		log.Fatal("GET URL Error! Please try again.")
	}
	defer response.Body.Close()

	var cResp NomicsResponse

	//JSON decoder
	err = json.NewDecoder(response.Body).Decode(&cResp)

	if err != nil {
		log.Fatal("UNMARSHAL Error! Please try again.", err)
	}

	return cResp.TextOutput(), nil

}

func PrintOutline(r ResponseFormat) string {
	result, err := r.GetUrlStr()
	if err != nil {
		log.Fatal("GetUrlstr Error! Please try again.", err)
	}
	fmt.Println("\n-------------\n", result)
	return result
}
func (c NomicsResponse) TextOutput() string {
	p := fmt.Sprintf(
		"Name: %s\nPrice: $ %s\nRank: %s\nHigh: $ %s\nCirculatingSupply: %s\nNumber of Traded Exchanges: %s\n",
		c[0].Name, c[0].CurrentPrice, c[0].MarketCapRank, c[0].AllTimeHigh, c[0].CirculatingSupply, c[0].NumExchangesTraded)
	return p
}
