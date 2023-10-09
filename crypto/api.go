package crypto

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/go-resty/resty/v2"
)

func GetSimplePriceParams() map[string]string {
	params := map[string]string{
		"ids":                 "",
		"vs_currencies":       "usd",
		"include_market_cap":  "true",
		"include_24hr_vol":    "true",
		"include_24hr_change": "true",
	}
	return params
}

// ListSupportedCryptos returns a list of supported cryptocurrencies from the API.
func ListSupportedCryptos() ([]string, error) {
	client := resty.New()
	url := "https://api.coingecko.com/api/v3/coins/list"

	var cryptos []struct {
		ID     string `json:"id"`
		Symbol string `json:"symbol"`
	}

	_, err := client.R().SetResult(&cryptos).Get(url)
	if err != nil {
		return nil, err
	}

	var cryptoNames []string
	for _, crypto := range cryptos {
		cryptoNames = append(cryptoNames, crypto.Symbol)
	}

	return cryptoNames, nil
}

// FetchCryptoDataList fetches data for a list of cryptocurrencies.
func FetchCryptoDataList(cryptoList []string) (map[string]interface{}, error) {
	client := resty.New()
	url := "https://api.coingecko.com/api/v3/simple/price"

	params := GetSimplePriceParams()

	// Compose the ids parameter with a comma-separated list of cryptocurrencies
	for _, crypto := range cryptoList {
		params["ids"] += crypto + ","
	}

	resp, err := client.R().SetQueryParams(params).Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == 200 {
		// Parse the JSON response into a map
		var data map[string]interface{}
		err := json.Unmarshal(resp.Body(), &data)
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	return nil, nil
}

// FetchCryptoData fetches info about a specific cryptocurrency
func FetchCryptoData(cryptoName string) (map[string]interface{}, error) {
	client := resty.New()
	url := "https://api.coingecko.com/api/v3/simple/price"

	params := GetSimplePriceParams()

	// catch responce
	resp, err := client.R().SetQueryParams(params).Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == 200 {
		// Parse the JSON response into a map
		var data map[string]interface{}
		err := json.Unmarshal(resp.Body(), &data)
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	return nil, nil
}

// CryptoData represents cryptocurrency data.
type CryptoData struct {
	Name         string  `json:"name"`
	Symbol       string  `json:"symbol"`
	MarketCap    float64 `json:"market_cap"`
	CurrentPrice float64 `json:"current_price"`
}

// FetchTopCryptosWithPrices fetches data for the top N cryptocurrencies by market capitalization
// including current prices.
func FetchTopCryptosWithPrices(count int) ([]CryptoData, error) {
	url := "https://api.coingecko.com/api/v3/coins/markets"
	params := fmt.Sprintf("?vs_currency=usd&order=market_cap_desc&per_page=%d&page=1&sparkline=false", count)
	// vs_currency=usd&order=market_cap_desc&per_page=10&page=1&sparkline=false&price_change_percentage=24h&locale=en

	resp, err := http.Get(url + params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []CryptoData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

func FetchTopCryptos(count int) ([]CryptoData, error) {
	data, err := FetchTopCryptosWithPrices(count)
	if err != nil {
		return nil, err
	}

	// Sort cryptocurrencies by market capitalization in descending order
	sort.Slice(data, func(i, j int) bool {
		return data[i].MarketCap > data[j].MarketCap
	})

	return data, nil
}
