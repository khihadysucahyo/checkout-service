package usecase

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/khihadysucahyo/checkout-service/domain"
)

type orderUsecase struct{}

func NewOrderUsecase() domain.OrderUsecase {
	return &orderUsecase{}
}

func checkStock(itemsStock domain.Items, item domain.RequestItem) bool {
	isAvailable := false

	for _, stockItem := range itemsStock {
		if stockItem.SKU == item.SKU {
			if stockItem.Quantity >= item.Quantity {
				isAvailable = true
			}
		}
	}

	return isAvailable
}

func getDetailItem(itemsStock domain.Items, item domain.RequestItem) domain.Item {
	var detailItem domain.Item

	for _, stockItem := range itemsStock {
		if stockItem.SKU == item.SKU {
			detailItem = stockItem
		}
	}

	return detailItem
}

func (c *orderUsecase) DoOrder(requestItems []domain.RequestItem) (domain.Order, error) {
	var err []error
	// GetItems
	itemsStock := domain.GetItems()

	// orderItems
	orderItems := []domain.ItemOrder{}

	// Check Stock and calculate total
	var total float64
	for i, item := range requestItems {
		if !checkStock(itemsStock, item) {
			err = append(err, errors.New("Item "+itemsStock[i].Name+" is not available / out of stock"))
		} else {

			detailItem := getDetailItem(itemsStock, item)

			// Buying more than 3 Alexa Speakers will have a 10% discount on all Alexa speakers
			totalItem := float64(item.Quantity) * detailItem.Price
			fmt.Println("totalItem", totalItem)
			fmt.Println("name: ", detailItem.Name)
			fmt.Println("price: ", detailItem.Price)
			fmt.Println("quantity: ", item.Quantity)
			fmt.Println("totalItem", totalItem)

			if item.SKU == "A304SD" && item.Quantity >= 3 {
				totalItem = totalItem - (totalItem * 0.1)
			}

			// Buy 3 Google Homes for the price of 2
			if item.SKU == "120P90" && item.Quantity == 3 {
				totalItem = totalItem - detailItem.Price
			}

			// Calculate total
			orderItems = append(orderItems, domain.ItemOrder{
				SKU:      detailItem.SKU,
				Name:     detailItem.Name,
				Price:    detailItem.Price,
				Quantity: item.Quantity,
				Total:    math.Round(totalItem*100) / 100,
			})
			total += float64(item.Quantity) * detailItem.Price

			// Each sale of a MacBook Pro comes with a free Raspberry Pi B
			if item.SKU == "43N23P" && checkStock(itemsStock, domain.RequestItem{SKU: "234234", Quantity: 1}) {
				orderItems = append(orderItems, domain.ItemOrder{
					SKU:      "234234",
					Name:     "Raspberry Pi B",
					Price:    0,
					Quantity: 1,
					Total:    0,
				})
			}

		}
	}

	if len(err) > 0 {
		//slice error to string
		var errString string
		for _, e := range err {
			errString += e.Error() + ";"
		}

		return domain.Order{}, errors.New(errString)
	}

	return domain.Order{
		ID:        "1",
		Customer:  "khihadysucahyo",
		Items:     orderItems,
		Total:     math.Round(total*100) / 100,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}
