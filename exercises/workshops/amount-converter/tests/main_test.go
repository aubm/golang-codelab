package amount_converter_test

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strings"
	"testing"
)

const host = "http://localhost:8080"

func TestShouldRespondCorrectly(t *testing.T) {
	requestUrl := host + "/convert"
	var baseEurAmount float64 = 5
	expectedContentType := "application/json"

	requestBody := fmt.Sprintf(`{
	  "currency": "EUR",
	  "target_currency": "CZK",
	  "amount": %v
	}`, baseEurAmount)

	eurResp, err := http.Post(requestUrl, "application/json", strings.NewReader(requestBody))
	if err != nil {
		t.Errorf("performing post request failed\nrequest url: %v\nrequest body: %v\ngot error: %v", requestUrl, requestBody, err)
		return
	}

	if contentType := eurResp.Header.Get("Content-Type"); contentType != "application/json" {
		t.Errorf("expected response content type to be %v, but got %v", expectedContentType, contentType)
	}

	eurRespBody := converterResponse{}
	if err := json.NewDecoder(eurResp.Body).Decode(&eurRespBody); err != nil {
		t.Errorf("json decoding the response failed\nrequest url: %v\nrequest body: %v\ngot error: %v", requestUrl, requestBody, err)
		return
	}

	convertedCzkAmount := eurRespBody.ConvertedAmount
	if convertedCzkAmount <= baseEurAmount {
		t.Errorf("it is unlikely that %v EUR give %v CZK, it seems that the convertion operation failed", baseEurAmount, convertedCzkAmount)
	}

	requestBody = fmt.Sprintf(`{
	  "currency": "CZK",
	  "target_currency": "EUR",
	  "amount": %v
	}`, convertedCzkAmount)

	czkResp, err := http.Post(requestUrl, "application/json", strings.NewReader(requestBody))
	if err != nil {
		t.Errorf("performing post request failed\nrequest url: %v\nrequest body: %v\ngot error: %v", requestUrl, requestBody, err)
		return
	}

	if contentType := czkResp.Header.Get("Content-Type"); contentType != "application/json" {
		t.Errorf("expected response content type to be %v, but got %v", expectedContentType, contentType)
	}

	czkRespBody := converterResponse{}
	if err := json.NewDecoder(czkResp.Body).Decode(&czkRespBody); err != nil {
		t.Errorf("json decoding the response failed\nrequest url: %v\nrequest body: %v\ngot error: %v", requestUrl, requestBody, err)
		return
	}

	convertedEurAmount := math.Floor(czkRespBody.ConvertedAmount + .5)
	if baseEurAmount != convertedEurAmount {
		t.Errorf("converting back CZK to EUR should give the original amount in EUR but it did not, got %v EUR", convertedEurAmount)
	}
}

func TestShouldRespondWithBadRequestStatusWhenJsonIsInvalid(t *testing.T) {
	requestUrl := host + "/convert"
	requestBody := `{`

	resp, err := http.Post(requestUrl, "application/json", strings.NewReader(requestBody))
	if err != nil {
		t.Errorf("performing post request failed\nrequest url: %v\nrequest body: %v\ngot error: %v", requestUrl, requestBody, err)
		return
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected response status code to be %v, but got %v", http.StatusBadRequest, resp.StatusCode)
	}
}

func TestShouldRespondWithBadRequestStatusWhenBaseCurrencyIsInvalid(t *testing.T) {
	requestUrl := host + "/convert"
	requestBody := `{
	  "currency": "JOHNDOE",
	  "target_currency": "CZK",
	  "amount": 5
	}`

	resp, err := http.Post(requestUrl, "application/json", strings.NewReader(requestBody))
	if err != nil {
		t.Errorf("performing post request failed\nrequest url: %v\nrequest body: %v\ngot error: %v", requestUrl, requestBody, err)
		return
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected response status code to be %v, but got %v", http.StatusBadRequest, resp.StatusCode)
	}
}

func TestShouldRespondWithBadRequestStatusWhenTargetCurrencyIsInvalid(t *testing.T) {
	requestUrl := host + "/convert"
	requestBody := `{
	  "currency": "EUR",
	  "target_currency": "JOHNDOE",
	  "amount": 5
	}`

	resp, err := http.Post(requestUrl, "application/json", strings.NewReader(requestBody))
	if err != nil {
		t.Errorf("performing post request failed\nrequest url: %v\nrequest body: %v\ngot error: %v", requestUrl, requestBody, err)
		return
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected response status code to be %v, but got %v", http.StatusBadRequest, resp.StatusCode)
	}
}

type converterResponse struct {
	ConvertedAmount float64 `json:"converted_amount"`
}
