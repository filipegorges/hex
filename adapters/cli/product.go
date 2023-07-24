package cli

import (
	"fmt"

	"github.com/filipegorges/hex/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, productPrice float64) (string, error) {
	switch action {
	case "create":
		product, err := service.Create(productName, productPrice)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Product ID %s with the name %s has the price of %f and is %s", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus()), nil
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return "", err
		}
		product, err = service.Enable(product)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Product %s has been enabled", product.GetName()), nil
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return "", err
		}
		product, err = service.Disable(product)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Product %s has been disabled", product.GetName()), nil
	default:
		product, err := service.Get(productId)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Product ID %s\nName %s\nPrice %f\nStatus %s\n", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus()), nil
	}
}
