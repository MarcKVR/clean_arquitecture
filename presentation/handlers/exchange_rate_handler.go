package handlers

import (
	"github.com/MarcKVR/clean_arquitecture/application/usecases"
	"github.com/gofiber/fiber/v2"
)

type ExchangeRateHandler struct {
	usecase *usecases.ExchangeRateUsecase
}

func NewExchangeRateHandler(app *fiber.App, usecase *usecases.ExchangeRateUsecase) {
	h := &ExchangeRateHandler{usecase: usecase}
	app.Get("/exchange-rate", h.GetExchangeRate)
}

func (h *ExchangeRateHandler) GetExchangeRate(c *fiber.Ctx) error {
	base := c.Query("base")
	target := c.Query("target")

	if base == "" || target == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "base and target query params are required",
		})
	}

	rate, err := h.usecase.GetExchangeRate(base, target)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(rate)
}
