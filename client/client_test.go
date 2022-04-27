package client

import (

	"testing"
	"github.com/stretchr/testify/assert"
)


func TestAPIFetch(t *testing.T) {
	var currency = "CAD"
	var crypto = "BTC"
	got, err := APIFetch(currency, crypto)

	if err != nil {
		t.Errorf("Error in API Fetch %v", err)
	}

	assert.NotEqual(t, nil, got)


}
