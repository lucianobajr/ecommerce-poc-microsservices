package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lucianobajr/ecommerce-poc-microsservices/go-api/internal/entity"
	"github.com/lucianobajr/ecommerce-poc-microsservices/go-api/internal/service"
)

type WebProductHandler struct {
	ProductService *service.ProductService
}

func NewWebProductHandler(ProductService *service.ProductService) *WebProductHandler {
	return &WebProductHandler{ProductService: ProductService}
}

func (wph *WebProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := wph.ProductService.GetProducts()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	json.NewEncoder(w).Encode(products)
}

func (wph *WebProductHandler) GetProductByCategoryID(w http.ResponseWriter, r *http.Request) {
	categoryID := chi.URLParam(r, "categoryID")

	if categoryID == "" {
		http.Error(w, "categoryID is required", http.StatusBadRequest)
	}

	products, err := wph.ProductService.GetProductByCategoryID(categoryID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	json.NewEncoder(w).Encode(products)
}

func (wph *WebProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	result, err := wph.ProductService.CreateProduct(product.Name, product.Description, product.Price, product.CategoryID, product.ImageURL)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	json.NewEncoder(w).Encode(result)
}

func (wph *WebProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
	}

	product, err := wph.ProductService.GetProduct(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	json.NewEncoder(w).Encode(product)
}
