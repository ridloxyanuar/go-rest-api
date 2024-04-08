package helper

import (
	"ridloxyanuar/go-rest-api/model/domain"
	"ridloxyanuar/go-rest-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
