package usecase

import (
	"errors"
	"firstpro/domain"
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
func AddCategory(category domain.Category) (domain.Category, error) {
	productResponse, err := repository.AddCategory(category)
	if err != nil {
		return domain.Category{}, err
	}
	return productResponse, nil
}
func UpdateCategory(current string, new string) (domain.Category, error) {
	result, err := repository.CheckCategory(current)
	if err != nil {
		return domain.Category{}, err
	}
	if !result {
		return domain.Category{}, errors.New("there is no category as you mentioned")
	}
	newCat, err := repository.UpdateCategory(current, new)
	if err != nil {
		return domain.Category{}, err
	}
	return newCat, err
}
func DeleteCategory(categoryID string) error {
	err := repository.DeleteCategory(categoryID)
	if err != nil {
		return err
	}
	return nil
}
