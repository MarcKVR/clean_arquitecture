package repositories

import "github.com/MarcKVR/clean_arquitecture/domain/models"

type ExchangeRateRepository interface {
	GetExchangeRate(base, target string) (*models.ExchangeRate, error)
}
