package client

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"Crypto_Market_Updates/model"
	"log"
	"fmt"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
func TestHandler(t *testing.T) {
	UrlTest, err := FiatCrypto("CAD", "LTC")
	if err != nil {
		log.Println(err)
	}
	req := httptest.NewRequest(http.MethodGet, UrlTest, nil)
	w := httptest.NewRecorder()

	Handler(w, req)

	want, got := http.StatusOK, w.Result().StatusCode
	if want != got {
		t.Fatalf("Expected a %d, instead got: %d", want, got)
	}

	fmt.Print(UrlTest)
	
	response, err := http.Get(UrlTest)

	//Error handling
	if err != nil {
		log.Fatal("GETURL Error! Please try again.")
	}
	defer response.Body.Close()
	
}

type Tests struct {
	name string
	server *httptest.Server
	response *model.NomicsResponse
	expectedError error
}

func TestAPIResponse(t *testing.T) {

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
		{
			"id": "LTC",
			"currency": "LTC",
			"symbol": "LTC",
			"name": "Litecoin",
			"logo_url": "https://s3.us-east-2.amazonaws.com/nomics-api/static/images/currencies/ltc.svg",
			"status": "active",
			"price": "129.67366368",
			"price_date": "2022-05-02T00:00:00Z",
			"price_timestamp": "2022-05-02T14:32:00Z",
			"circulating_supply": "70198121",
			"max_supply": "84000000",
			"market_cap": "9102847499",
			"market_cap_dominance": "0.0039",
			"num_exchanges": "356",
			"num_pairs": "3773",
			"num_pairs_unmapped": "601",
			"first_candle": "2013-03-25T00:00:00Z",
			"first_trade": "2013-03-25T00:00:00Z",
			"first_order_book": "2018-08-29T00:00:00Z",
			"rank": "24",
			"rank_delta": "1",
			"high": "471.89482175",
			"high_timestamp": "2021-05-09T00:00:00Z",
			"1d": {
			  "volume": "992583800.59",
			  "price_change": "5.15494063",
			  "price_change_pct": "0.0414",
			  "volume_change": "36519560.14",
			  "volume_change_pct": "0.0382",
			  "market_cap_change": "362763679.58",
			  "market_cap_change_pct": "0.0415"
			}
		  }`))
	}))
	defer s.Close()
	r := &model.NomicsResponse{Name:"Litecoin", CurrentPrice:"129.83668602", MarketCapRank:"24", AllTimeHigh:"471.89482175", CirculatingSupply:"70198121", NumExchangesTraded:"356"}

	resp, err := GetURL(s.URL)
	
	if !reflect.DeepEqual(resp, r) {
		t.Errorf("FAILED: expected %v, got %v\n",resp, r)
	}
	if !errors.Is(err, nil) {
		t.Errorf("Expected error FAILED: expected %v got %v\n", nil, err)
	}

}
