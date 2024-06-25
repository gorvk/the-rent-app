package routes

import (
	"net/http"

	controllers "github.com/gorvk/rent-app/api-services/controllers/order"
)

func init() {
	http.HandleFunc("/api/order/create", controllers.CreateOrder)
	http.HandleFunc("/api/order/get-current-user-orders", controllers.GetCurrentUserOrders)
}
