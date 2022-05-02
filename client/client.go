package client

import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"log"
	"net/http"
	"os"

	"Crypto_Market_Updates/model"
)

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

func GetURL(url string) (model.NomicsResponse, error) {

	//Get function
	response, err := http.Get(url)

	//Error handling
	if err != nil {
		log.Fatal("GET URL Error! Please try again.")
	}
	defer response.Body.Close()
	//body, err := ioutil.ReadAll(response.Body)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	//Variable type model
	var cResp model.NomicsResponse
	//fmt.Printf("**TYPE*** %#v\n", &cResp)

	//JSON decoder
	err = json.NewDecoder(response.Body).Decode(&cResp)

	if err != nil {
		log.Fatal("UNMARSHAL Error! Please try again.", err)
	}

	return cResp, nil

}