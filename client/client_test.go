package client

import (
	"net/http"
	"net/http/httptest"

	"testing"

	"Crypto_Market_Updates/mocks"

	"github.com/stretchr/testify/assert"

	"log"

	gomock "github.com/golang/mock/gomock"
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

func TestPrintOutline(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mock := mocks.NewMockResponseFormat(mockCtrl)

	resultWant := "Correct!"
	mock.EXPECT().GetUrlStr().Return(resultWant, nil).Times(1)

	resultGot := PrintOutline(mock)

	assert.Equal(t, resultWant, resultGot, "Correct should be printed!")
}
