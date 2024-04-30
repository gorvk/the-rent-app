package routes

import (
	"net/http"

	controllers "github.com/gorvk/rent-app/server/api-services/controllers/user"
)

func init() {
	http.HandleFunc("/api/user/get-all", controllers.GetAllUsers)
	http.HandleFunc("/api/user/delete", controllers.DeleteUser)
	http.HandleFunc("/api/user/update", controllers.UpdateUser)
}
