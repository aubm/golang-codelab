package api

import (
	"encoding/json"
	"net/http"

	"github.com/aubm/golang-codelab/solutions/workshops/amount-converter/amounts"
)

type AmountHandlers struct {
	Converter interface {
		Convert(amount float64, currency, targetCurrency string) (float64, error)
	}
}

func (h *AmountHandlers) Convert(w http.ResponseWriter, r *http.Request) {
	userConvertRequest := convertRequest{}
	if err := json.NewDecoder(r.Body).Decode(&userConvertRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	convertedAmount, err := h.Converter.Convert(
		userConvertRequest.Amount,
		userConvertRequest.Currency,
		userConvertRequest.TargetCurrency,
	)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == amounts.ErrInvalidBaseCurrency || err == amounts.ErrTargetCurrencyNotFound {
			statusCode = http.StatusBadRequest
		}
		http.Error(w, err.Error(), statusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"converted_amount": convertedAmount,
	})
}

type convertRequest struct {
	Currency       string  `json:"currency"`
	TargetCurrency string  `json:"target_currency"`
	Amount         float64 `json:"amount"`
}
