package service

import (
	"bebeziansback/product/mappers"
	"bebeziansback/product/models"
	"bebeziansback/product/response"
	"bebeziansback/server/config"
)

func GetAllCategories() (response.ProductCategoryResponse, error) {
	var categories []models.ProductCategory
	if err := config.DB.
		Find(&categories).Error; err != nil {
		return response.ProductCategoryResponse{}, err
	}
	response := mappers.MapProductCategoriesToResponse(categories)
	return response, nil
}

func GetPopProducts() (response.PopProductsResponse, error) {
	var products []models.Product
	if err := config.DB.
		Preload("Image").
		Preload("Description").
		Preload("Category").
		Limit(20).
		Find(&products).Error; err != nil {
		return response.PopProductsResponse{}, err
	}
	response := mappers.MapPopProductsToResponse(products)
	return response, nil
}

func GetProductsByCategory(name string) (response.ProductsByCategoryResponse, error) {
	var products []models.Product
	if err := config.DB.
		Joins("JOIN productcategory ON productcategory.id = products.category_id").
		Preload("Image").
		Preload("Description").
		Preload("Category").
		Where("productcategory.name = ?", name).
		Find(&products).Error; err != nil {
		return response.ProductsByCategoryResponse{}, err
	}
	response := mappers.MapProductsByCategoryToResponse(products)
	return response, nil
}

func GetProductById(id int) (response.ProductResponse, error) {
	var product models.Product
	var sizes []string

	if err := config.DB.
		Preload("Image").
		Preload("Description").
		Preload("Category").
		Where("products.id = ?", id).
		Find(&product).Error; err != nil {
		return response.ProductResponse{}, err
	}

	if err := config.DB.Table("size").
		Where("product_id = ?", id).
		Pluck("size", &sizes).Error; err != nil { 
		return response.ProductResponse{}, err
	}
	productResponse := mappers.MapProductToResponse(product,sizes)

	return productResponse, nil
}