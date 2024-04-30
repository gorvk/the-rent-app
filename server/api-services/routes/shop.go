package routes

import (
	"net/http"

	controllers "github.com/gorvk/rent-app/server/api-services/controllers/shop"
)

func init() {
	http.HandleFunc("/api/shop/create", controllers.CreateShop)
	http.HandleFunc("/api/shop/get-shop", controllers.GetShopByEmail)
	http.HandleFunc("/api/shop/get-all-shops", controllers.GetAllShops)
	http.HandleFunc("/api/shop/update", controllers.UpdateShop)
	http.HandleFunc("/api/shop/delete", controllers.DeleteShop)
}
