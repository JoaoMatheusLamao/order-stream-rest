package models

import (
	"errors"
	"strings"
	"time"
)

// Product is a struct that represents a product
type Product struct {
	SKU           string               `json:"sku,omitempty" bson:"sku"`
	Name          string               `json:"name,omitempty" bson:"name"`
	Stock         int64                `json:"stock,omitempty" bson:"stock"`
	Price         float64              `json:"price,omitempty" bson:"price"`
	InsertionDate time.Time            `json:"insertion_date" bson:"insertion_date"`
	UpdateHistory []UpdateHistoryEntry `json:"-" bson:"update_history"`
}

// UpdateHistoryEntry is a struct that represents an entry in the update history
type UpdateHistoryEntry struct {
	Date       time.Time `json:"date" bson:"date"`
	ProductOld string    `json:"productOld" bson:"product_old"`
}

// Validate is a method that validates the product
func (p *Product) Validate() error {
	var validationErrors []string

	if p.SKU == "" {
		validationErrors = append(validationErrors, "SKU is required")
	}

	if p.Name == "" {
		validationErrors = append(validationErrors, "Name is required")
	}

	if p.Price < 0 {
		validationErrors = append(validationErrors, "Price cannot be negative")
	}

	if len(validationErrors) > 0 {
		return errors.New("Validation errors: " + strings.Join(validationErrors, ", "))
	}

	return nil
}
