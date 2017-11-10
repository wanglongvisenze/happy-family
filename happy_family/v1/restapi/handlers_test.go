package restapi

import (
	"testing"

	"github.com/go-openapi/strfmt"

	"github.com/wanglongvisenze/happy-family/happy_family/v1/models"
)

func TestGetEggs(t *testing.T) {
	eggs := getProduct("eggs")

	if eggs.Name != "eggs" || eggs.Quantity <= 0 {
		t.Fail()
	}
}

func TestByWife(t *testing.T) {
	eggs := getProduct("eggs")
	eggQuantity := eggs.Quantity
	milk := &models.Product{"milk", 1, 0}

	order := &models.Order{"", nil, strfmt.NewDateTime(), 0}
	order.Products = make([]*models.Product, 0)
	order.Products = append(order.Products, milk)

	if eggs.Quantity >= 0 {
		eggs.Quantity = 12
		order.Products = append(order.Products, &eggs)
	}

	filledOrder := placeOrder(order)

	if len(filledOrder.Products) != len(order.Products) {
		t.Fail()
	}

	eggs = getProduct("eggs")
	newEggQuantity := eggs.Quantity

	if newEggQuantity + 12 != eggQuantity && eggQuantity != 0 {
		t.Fail()
	}
}
