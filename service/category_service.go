package service

import (
	"context"
	"ridloxyanuar/go-rest-api/model/web"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	GetCategories(ctx context.Context) []web.CategoryResponse
	DeleteCategory(ctx context.Context, categoryId int)
	UpdateCategory(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	GetCategoryById(ctx context.Context, categoryId int) web.CategoryResponse
}
