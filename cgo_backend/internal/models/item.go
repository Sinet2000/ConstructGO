package models

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name          string             `json:"name"`
	SKU           string             `json:"sku"`
	Image         string             `json:"image"`
	Price         float64            `json:"price"`
	Weight        float64            `json:"weight,omitempty"`
	Quantity      int                `json:"quantity"`
	ItemType      string             `json:"itemType"` // Could be "PerItem", "PerSize", "PerWeight" etc.
	PurchasePrice float64            `json:"purchasePrice,omitempty"`

	Height float64 `json:"height,omitempty"`
	Width  float64 `json:"width,omitempty"`
	Length float64 `json:"length,omitempty"`
}

// IsAvailable checks if the item quantity is more than 0
func (i *Item) IsAvailable() bool {
	return i.Quantity > 0
}

// DeductQuantity reduces the item quantity by a given amount
func (i *Item) DeductQuantity(amount int) error {
	if amount <= i.Quantity {
		i.Quantity -= amount
		return nil
	}
	return errors.New("insufficient quantity")
}

// AddQuantity increases the quantity of the item by a given amount
func (i *Item) AddQuantity(amount int) {
	i.Quantity += amount
}

// UpdatePrice updates the selling price of the item
func (i *Item) UpdatePrice(newPrice float64) {
	i.Price = newPrice
}

// UpdatePurchasePrice updates the purchase price of the item
func (i *Item) UpdatePurchasePrice(newPrice float64) {
	i.PurchasePrice = newPrice
}

func (i *Item) PricePerM2() (float64, error) {
	if i.Height == 0 || i.Width == 0 {
		return 0, errors.New("invalid dimensions")
	}

	areaM2 := i.Height * i.Width
	pricePerM2 := i.Price / areaM2

	return pricePerM2, nil
}

// UpdateItem updates the details of the item
func (i *Item) UpdateItem(updatedItem Item) {
	i.Name = updatedItem.Name
	i.SKU = updatedItem.SKU
	i.Image = updatedItem.Image
	i.Price = updatedItem.Price
	i.Weight = updatedItem.Weight
	i.Quantity = updatedItem.Quantity
	i.ItemType = updatedItem.ItemType
	i.PurchasePrice = updatedItem.PurchasePrice
	i.Height = updatedItem.Height
	i.Width = updatedItem.Width
	i.Length = updatedItem.Length
}

// CalculateVolume calculates the volume of the item
func (i *Item) CalculateVolume() float64 {
	return i.Height * i.Width * i.Length
}

// CalculatePricePerVolume calculates the price per unit volume
func (i *Item) CalculatePricePerVolume() float64 {
	volume := i.CalculateVolume()
	if volume == 0 {
		return 0
	}
	return i.Price / volume
}

// UpdateDimensions updates the dimensions of the item
func (i *Item) UpdateDimensions(height, width, length float64) {
	i.Height = height
	i.Width = width
	i.Length = length
}

// CalculateWeightPerVolume calculates the weight per unit volume
func (i *Item) CalculateWeightPerVolume() float64 {
	volume := i.CalculateVolume()
	if volume == 0 {
		return 0
	}
	return i.Weight / volume
}
