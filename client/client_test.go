package client

import (
	"strings"
	"testing"
)

func TestAPIFetch(t *testing.T) {
	var currency = "CAD"
	var crypto = "BTC"
	got, err := APIFetch(currency, crypto)

	if err != nil {
		t.Errorf("Error in API Fetch %v", err)
	}
	if !strings.Contains(got, "Current Price") {
		t.Errorf("Test did not pass %v", got)
	}

}
