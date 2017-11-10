package restapi

import (
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/go-openapi/strfmt"

	"github.com/wanglongvisenze/happy-family/happy_family/v1/models"
)

var price = map[string]float64 {
	"milk": 1,
	"eggs": 1,
}

var stock = map[string]float64 {
	"milk": 1000,
	"eggs": 1000,
}

var orders = make([]*models.Order, 0)

func getProduct(name string) models.Product {
	return models.Product{name, stock[name], price[name]}
}

func placeOrder(order *models.Order) *models.Order {
	filledOrder := &models.Order{uuid.New().String(), make([]*models.Product, 0),
		strfmt.DateTime(time.Now()), 0}
	for _, prod := range order.Products {
		prod.Quantity = math.Min(prod.Quantity, stock[prod.Name])
		prod.UnitPrice = price[prod.Name]
		stock[prod.Name] -= prod.Quantity
		filledOrder.TotalPrice += prod.UnitPrice * prod.Quantity
		filledOrder.Products = append(filledOrder.Products, prod)
	}
	orders = append(orders, filledOrder)
	return filledOrder
}
