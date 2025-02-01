package usecases

import (
	"github.com/MarcKVR/clean_arquitecture/domain/models"
	"github.com/MarcKVR/clean_arquitecture/domain/repositories"
)

type ExchangeRateUsecase struct {
	repo repositories.ExchangeRateRepository
}

func NewExchangeRateUsecase(repo repositories.ExchangeRateRepository) *ExchangeRateUsecase {
	return &ExchangeRateUsecase{repo: repo}
}

func (u *ExchangeRateUsecase) GetExchangeRate(base, target string) (*models.ExchangeRate, error) {
	return u.repo.GetExchangeRate(base, target)
}
