package model

import (
	"fmt"
	"log"
)

//Cryptoresponse is exported and models the data we receive
type NomicsResponse []struct {
	Name               string `json:"name"`
	CurrentPrice       string `json:"price"`
	MarketCapRank      string `json:"rank"`
	AllTimeHigh        string `json:"high"`
	CirculatingSupply  string `json:"circulating_supply"`
	NumExchangesTraded string `json:"num_exchanges"`
}

//TextOutput is exported and formats the data to plain text
func (c NomicsResponse) TextOutput() string {
	if c[0].AllTimeHigh == "" {
		log.Fatal("Error in currency or crypto codes. Please check for correct spelling.")
	}

	p := fmt.Sprintf(
		"\n---------------------\nName: %s\nCurrent Price : $%s\nMarket Cap Rank: %s\nAll Time High: $%s\nCirculating Supply: %s\nNumber of Exchanges Traded: %s\n",
		c[0].Name, c[0].CurrentPrice, c[0].MarketCapRank, c[0].AllTimeHigh, c[0].CirculatingSupply, c[0].NumExchangesTraded)
	return p
}
