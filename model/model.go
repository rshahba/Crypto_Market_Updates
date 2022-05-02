package model

// import (
// 	"fmt"
// 	"log"
// )

//Models the data we receive
type NomicsResponse struct {
	Name               string `json:"name"`
	CurrentPrice       string `json:"price"`
	MarketCapRank      string `json:"rank"`
	AllTimeHigh        string `json:"high"`
	CirculatingSupply  string `json:"circulating_supply"`
	NumExchangesTraded string `json:"num_exchanges"`
}
