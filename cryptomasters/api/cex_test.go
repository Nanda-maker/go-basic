package api_test

import (
	"go/cryptomasters/api"
	"testing"
)

func TestAPICall(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Error("Error was not found")
	}
}
