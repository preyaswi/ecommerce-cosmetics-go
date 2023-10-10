package usecase

import (
	"firstpro/repository"
	"firstpro/utils/models"
	"fmt"
)

func ShowAllProducts(page int, count int) ([]models.ProductBrief, error) {

	productsBrief, err := repository.ShowAllProducts(page, count)
	if err != nil {
		return []models.ProductBrief{}, err
	}
	for i := range productsBrief {
		p := &productsBrief[i]
		if p.Quantity == 0 {
			p.ProductStatus = "out of stock"
		} else {
			p.ProductStatus = "in stock"
		}
	}
	return productsBrief, nil
}
func ShowIndividualProducts(id int) (*models.ProductBrief, error) {
	product, err := repository.ShowIndividualProducts(id)
	fmt.Println("ahgfvgf", product)
	if err != nil {
		fmt.Println("ahcfaecf")
		return &models.ProductBrief{}, err
	}
	fmt.Println("321423")
	return product, nil

}
