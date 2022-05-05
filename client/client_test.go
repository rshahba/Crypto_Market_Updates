package client

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	//"Crypto_Market_Updates/mocks"
	
	"log"
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
	
	response, err := http.Get(UrlTest)

	//fmt.Print(response)
	//OK 200

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
			"price": "130.31892984",
			"rank": "25",
			"high": "471.89482189",
			"circulating_supply": "70212558",
			"num_exchanges": "356"
		  }]`))
	}))
	defer s.Close()
		r := "Name: Litecoin\nPrice: $ 130.31892984\nRank: 25\nHigh: $ 471.89482189\nCirculatingSupply: 70212558\nNumber of Traded Exchanges: 356\n"

	resp, err := GetUrlStr(s.URL)

	if !reflect.DeepEqual(resp, r) {
		t.Errorf("FAILED: expected %v, got %v\n",resp, r)
	}
	if !errors.Is(err, nil) {
		t.Errorf("Expected error FAILED: expected %v got %v\n", nil, err)
	}

}

func TestGetURL(t *testing.T){
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
		[{
			"name": "Litecoin",
			"price": "130.31892984",
			"rank": "25",
			"high": "471.89482189",
			"circulating_supply": "70212558",
			"num_exchanges": "356"
		  }]`))
	}))
	defer s.Close()
	r := "Name: Litecoin\nPrice: $ 130.31892984\nRank: 25\nHigh: $ 471.89482189\nCirculatingSupply: 70212558\nNumber of Traded Exchanges: 356\n"	
	resp, err := GetUrlStr(s.URL)
	if err != nil {
		log.Fatal("GET URL Error! Please try again.")
	}
	assert.Equal(t, r, resp)
}

// func TestNoAPIResoponse(t *testing.T){
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockClient := mocks.NewMockresponseFormat(mockCtrl)

// 	mockClient.EXPECT().responseFormat().Return(nil).Times(1)
	
// }