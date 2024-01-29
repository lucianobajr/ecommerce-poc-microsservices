package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lucianobajr/ecommerce-poc-microsservices/go-api/internal/database"
	"github.com/lucianobajr/ecommerce-poc-microsservices/go-api/internal/service"
	"github.com/lucianobajr/ecommerce-poc-microsservices/go-api/internal/webserver"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:dummypassword@tcp(localhost:3306)/catalog")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	productDB := database.NewProductDB(db)

	categoryService := service.NewCategoryService(*categoryDB)
	productService := service.NewProductService(*productDB)

	webCategoryHandler := webserver.NewWebCategoryHandler(categoryService)
	webProductHandler := webserver.NewWebProductHandler(productService)

	c := chi.NewRouter()

	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	// categories routes
	c.Get("/category/{id}", webCategoryHandler.GetCategory)
	c.Get("/category", webCategoryHandler.GetCategories)
	c.Post("/category", webCategoryHandler.CreateCategory)

	// products routes
	c.Get("/product/{id}", webProductHandler.GetProduct)
	c.Get("/product", webProductHandler.GetProducts)
	c.Post("/product", webProductHandler.CreateProduct)
	c.Get("/product/category/{categoryID}", webProductHandler.GetProductByCategoryID)

	fmt.Println("Server is running on port 8080")

	http.ListenAndServe(":8080", c)
}
