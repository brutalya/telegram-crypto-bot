package crypto

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

// GetCryptoPrice fetches the current price of a specific cryptocurrency.
func GetCryptoPrice(cryptoName string) (float64, error) {
	client := resty.New()
	url := "https://api.coingecko.com/api/v3/simple/price"

	params := map[string]string{
		"ids":           cryptoName,
		"vs_currencies": "usd",
	}

	resp, err := client.R().SetQueryParams(params).Get(url)
	if err != nil {
		return 0.0, err
	}

	if resp.StatusCode() == 200 {
		var price map[string]map[string]float64
		// Parse the JSON response into a map
		err := json.Unmarshal(resp.Body(), &price)
		if err != nil {
			return 0.0, err
		}

		return price[cryptoName]["usd"], nil
	}

	return 0.0, nil
}
