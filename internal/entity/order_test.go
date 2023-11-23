package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfItGetsAnError(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validade(), "id is required")
}

func TestIfItGetsAnErrorWhenPriceIsZero(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.Validade(), "price must be greater than zero")
}

func TestIfItGetsAnErrorWhenTaxIsZero(t *testing.T) {
	order := Order{ID: "123", Price: 10.0}
	assert.Error(t, order.Validade(), "invalid tax")
}

func TestFinalPrice(t *testing.T) {
	order := Order{ID: "123", Price: 10.0, Tax: 1.0}
	assert.NoError(t, order.Validade())
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 11.0, order.FinalPrice)
}
