package amounts

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrInvalidBaseCurrency    = errors.New("invalid base currency")
	ErrTargetCurrencyNotFound = errors.New("target currency not found")
)

type Converter struct {
	FixerApiUrl string
}

func (c *Converter) Convert(amount float64, currency, targetCurrency string) (float64, error) {
	rateList, err := c.fetchRate(currency)
	if err != nil {
		return 0, err
	}

	conversionRate, ok := rateList[targetCurrency]
	if !ok {
		return 0, ErrTargetCurrencyNotFound
	}

	return amount * conversionRate, nil
}

func (c *Converter) fetchRate(baseCurrency string) (rates, error) {
	resp, err := http.Get(fmt.Sprintf("%s/latest?base=%s", c.FixerApiUrl, baseCurrency))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode > 399 && resp.StatusCode < 500 {
			return nil, ErrInvalidBaseCurrency
		}
		return nil, fmt.Errorf("status code is %v", resp.StatusCode)
	}

	body := fixerApiResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body.Rates, nil
}

type fixerApiResponse struct {
	Rates rates `json:"rates"`
}

type rates map[string]float64
