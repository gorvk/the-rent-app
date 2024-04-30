package routes

import (
	"net/http"

	controllers "github.com/gorvk/rent-app/server/api-services/controllers/product"
)

func init() {
	http.HandleFunc("/api/product/create", controllers.CreateProduct)
	http.HandleFunc("/api/product/get-product", controllers.GetProductById)
	http.HandleFunc("/api/product/get-all-products", controllers.GetAllProducts)
	http.HandleFunc("/api/product/update", controllers.UpdateProduct)
	http.HandleFunc("/api/product/delete", controllers.DeleteProduct)
}
