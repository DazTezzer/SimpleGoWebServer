package mappers

import (
	"bebeziansback/product/models"
	"bebeziansback/product/response"
	"fmt"
)

func MapProductCategoriesToResponse(categories []models.ProductCategory) response.ProductCategoryResponse {
	var categoryResponses []response.Category
	for _, category := range categories {
		categoryResponses = append(categoryResponses, response.Category{
			Name: category.Name,
		})
	}
	return response.ProductCategoryResponse{
		Categories: categoryResponses,
	}
}

func MapPopProductsToResponse(products []models.Product) response.PopProductsResponse {
	var productResponses []response.ProductResponsePart
	for _, product := range products {
		productResponse := response.ProductResponsePart{
			ProductID:       product.Id,
			Name:            product.Name,
			ImageURL:        product.Image.ImageUrl,
			Price:           fmt.Sprintf("%.2f", product.Price),
			DescriptionInfo: product.Description.AdditionalInfo,
		}
		productResponses = append(productResponses, productResponse)
	}

	return response.PopProductsResponse{Products: productResponses}
}

func MapProductsByCategoryToResponse(products []models.Product) response.ProductsByCategoryResponse {
	var productResponses []response.ProductResponsePart
	for _, product := range products {
		productResponse := response.ProductResponsePart{
			ProductID:       product.Id,
			Name:            product.Name,
			ImageURL:        product.Image.ImageUrl,
			Price:           fmt.Sprintf("%.2f", product.Price),
			DescriptionInfo: product.Description.AdditionalInfo,
		}
		productResponses = append(productResponses, productResponse)
	}

	return response.ProductsByCategoryResponse{Products: productResponses}
}

func MapProductToResponse(product models.Product, sizes []string) response.ProductResponse {
	productResponse := response.ProductResponse{
		ProductID:   product.Id,
		Name:        product.Name,
		ImageURL:    product.Image.ImageUrl,
		Price:       fmt.Sprintf("%.2f", product.Price),
		Description: product.Description,
		Size:        sizes,
	}

	return productResponse
}
