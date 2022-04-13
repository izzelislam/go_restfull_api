package main

import (
	"fmt"
	"go_restfull_api/app"
	"go_restfull_api/controller"
	"go_restfull_api/repository"
	"go_restfull_api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {

	validate := validator.New()
	db := app.NewDb()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	seerver := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	fmt.Println("Server started on port 3000")
	err := seerver.ListenAndServe()
	// helper.PanicIfError(err)
	fmt.Println("server error", err)
}
