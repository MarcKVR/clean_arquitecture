package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/MarcKVR/clean_arquitecture/domain/models"
	"github.com/MarcKVR/clean_arquitecture/domain/repositories"
)

type ExchangeRateRepositoryApi1 struct {
	apiURL string
}

func NewExchangeRateRepository(apiURL string) repositories.ExchangeRateRepository {
	return &ExchangeRateRepositoryApi1{apiURL: apiURL}
}

func (r *ExchangeRateRepositoryApi1) GetExchangeRate(base, target string) (*models.ExchangeRate, error) {
	url := r.apiURL + "/061c1baefab2369b22f5c82b/pair/" + base + "/" + target

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received error code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	formattedRate, ok := data["conversion_rate"].(float64)
	baseRate, ok1 := data["base_code"].(string)
	targetRate, ok2 := data["target_code"].(string)

	if !ok || !ok1 || !ok2 {
		return nil, fmt.Errorf("conversion operations fails")
	}

	return &models.ExchangeRate{
		BaseCurrency:   baseRate,
		TargetCurrency: targetRate,
		Rate:           formattedRate,
	}, nil
}
