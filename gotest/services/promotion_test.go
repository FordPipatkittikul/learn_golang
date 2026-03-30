package services_test

import (
	"errors"
	"gotest/repositories"
	"gotest/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {
	
	type testCase struct {
		name 				string
		PurchaseMin 		int
		DiscountPercent		int
		amount 				int
		expected			int
	}

	cases := []testCase{
		{name: "applied 100", PurchaseMin: 100, DiscountPercent: 20, amount: 100, expected: 80},
		{name: "applied 200", PurchaseMin: 100, DiscountPercent: 20, amount: 200,expected: 160},
		{name: "applied 300", PurchaseMin: 100, DiscountPercent: 20, amount: 300,expected: 240},
		{name: "not applied", PurchaseMin: 100, DiscountPercent: 20, amount: 50,expected: 50},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {

			// Arrange
			promoRepo := repositories.NewPromotionRepositoryMock()
			promoRepo.On("GetPromotion").Return(repositories.Promotion{
				ID:					1,
				PurchaseMin: 		c.PurchaseMin,
				DiscountPercent: 	c.DiscountPercent,
			}, nil)
			promoService := services.NewPromotionService(promoRepo)

			// Act
			discount, _ := promoService.CalculateDiscount(c.amount)
	
			// Assert
			assert.Equal(t, c.expected, discount)
		})
	}

	t.Run("purchase amount zero", func(t *testing.T) { 
		// Arrange
		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{
			ID:					1,
			PurchaseMin: 		100,
			DiscountPercent: 	20,
		}, nil)
		promoService := services.NewPromotionService(promoRepo)

		// Act
		_, err := promoService.CalculateDiscount(0)

		// Assert
		assert.ErrorIs(t, err, services.ErrZeroAmount)
		promoRepo.AssertNotCalled(t, "GetPromotion")
	})

	t.Run("repository error", func(t *testing.T) { 
		// Arrange
		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{}, errors.New(""))
		promoService := services.NewPromotionService(promoRepo)

		// Act
		_, err := promoService.CalculateDiscount(100)

		// Assert
		assert.ErrorIs(t, err, services.ErrRepository)
	})
}