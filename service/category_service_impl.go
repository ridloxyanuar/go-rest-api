package service

import (
	"context"
	"database/sql"
	"ridloxyanuar/go-rest-api/exception"
	"ridloxyanuar/go-rest-api/helper"
	"ridloxyanuar/go-rest-api/model/domain"
	"ridloxyanuar/go-rest-api/model/web"
	"ridloxyanuar/go-rest-api/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (c *CategoryServiceImpl) CreateCategory(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := c.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := c.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitorRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = c.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (c *CategoryServiceImpl) DeleteCategory(ctx context.Context, categoryId int) {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitorRollback(tx)

	category, err := c.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	c.CategoryRepository.Delete(ctx, tx, category)
}

func (c *CategoryServiceImpl) GetCategories(ctx context.Context) []web.CategoryResponse {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitorRollback(tx)

	categories := c.CategoryRepository.FindAll(ctx, tx)

	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, helper.ToCategoryResponse(category))
	}

	return categoryResponses

}

func (c *CategoryServiceImpl) GetCategoryById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitorRollback(tx)

	category, err := c.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (c *CategoryServiceImpl) UpdateCategory(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := c.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := c.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitorRollback(tx)

	category, err := c.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = request.Name

	category = c.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}
