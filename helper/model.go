package helper

import (
	"go_restfull_api/model/domain"
	"go_restfull_api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryresponses []web.CategoryResponse
	for _, category := range categories {
		categoryresponses = append(categoryresponses, ToCategoryResponse(category))
	}
	return categoryresponses
}
