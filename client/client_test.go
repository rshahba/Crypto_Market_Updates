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

	fmt.Print(response)

	//Error handling
	if err != nil {
		log.Fatal("GETURL Error! Please try again.")
	}
	defer response.Body.Close()
	
}


func TestAPIResponse(t *testing.T) {

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
		[{
			"name": "Litecoin",
			"price": "129.83668602",
			"rank": "24",
			"high": "471.89482175",
			"circulating_supply": "70198121",
			"num_exchanges": "356"
		  }]`))
	}))
	defer s.Close()
	
	r := model.NomicsResponse{{Name:"Litecoin", CurrentPrice:"129.83668602", MarketCapRank:"24", AllTimeHigh:"471.89482175", CirculatingSupply:"70198121", NumExchangesTraded:"356"},}


	resp, err := GetURL(s.URL)

	if !reflect.DeepEqual(resp, r) {
		t.Errorf("FAILED: expected %v, got %v\n",resp, r)
	}
	if !errors.Is(err, nil) {
		t.Errorf("Expected error FAILED: expected %v got %v\n", nil, err)
	}

}
