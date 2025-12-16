package api

import (
	"encoding/json"
	"fmt"
	"go/cryptomasters/datatypes"
	"io"
	"net/http"
	"strings"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	if len(currency) != 3 {
		return nil, fmt.Errorf("3 characters required; %d received", len(currency))
	}

	upCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))
	if err != nil {
		return nil, err
	}
	var response CEXResponse
	if res.StatusCode == http.StatusOK {
		bodyByte, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyByte, &response)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, fmt.Errorf("status code recieved: %v", res.StatusCode)
	}
	rate := datatypes.Rate{Currency: currency, Price: float64(response.Bid)}
	return &rate, nil
}
