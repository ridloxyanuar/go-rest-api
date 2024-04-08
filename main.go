package main

import (
	"net/http"
	"ridloxyanuar/go-rest-api/app"
	"ridloxyanuar/go-rest-api/controller"
	"ridloxyanuar/go-rest-api/exception"
	"ridloxyanuar/go-rest-api/helper"
	"ridloxyanuar/go-rest-api/middleware"
	"ridloxyanuar/go-rest-api/repository"
	"ridloxyanuar/go-rest-api/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validator := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validator)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	//error handler
	router.PanicHandler = exception.ErrorHandler

	// run to server
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
		// ErrorLog: log.New(os.Stderr, "log: ", log.Lshortfile),
		// ReadTimeout: 5 * time.Second,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
