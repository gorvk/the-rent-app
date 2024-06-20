package routes

import (
	"net/http"

	controllers "github.com/gorvk/rent-app/api-services/controllers/order"
)

func init() {
	http.HandleFunc("/api/order/create-order", controllers.CreateOrder)
}
