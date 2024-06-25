package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorvk/rent-app/api-services/common"
	"github.com/gorvk/rent-app/api-services/common/constants"
	"github.com/gorvk/rent-app/api-services/common/types"
	models "github.com/gorvk/rent-app/api-services/models/order"
)

func GetCurrentUserOrders(w http.ResponseWriter, r *http.Request) {

	var err error
	err = common.CheckHttpResponseType(w, r, http.MethodGet)
	if err != nil {
		return
	}

	user, err := common.IsAuthenticated(r)
	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_HTTP_UNAUTHORIZED, http.StatusUnauthorized)
		return
	}

	rows, err := models.GetCurrentUserOrders(user.Id)
	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_DB_UNABLE_TO_GET_RECORD, http.StatusInternalServerError)
		return
	}

	orders := []types.OrderListView{}
	defer rows.Close()
	for rows.Next() {
		row := types.OrderListView{}
		err = rows.Scan(
			&row.Id,
			&row.ProductId,
			&row.BuyerId,
			&row.OrderStatus,
			&row.ProductName,
			&row.Price,
		)
		if err != nil {
			common.HandleHttpError(err, w, constants.ERROR_DB_UNABLE_TO_GET_RECORD, http.StatusInternalServerError)
			return
		}
		orders = append(orders, row)
	}

	var response types.RESPONSE_PARAMETERS
	response.Result = orders
	response.IsSuccess = true
	data, err := json.Marshal(response)
	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_HTTP_UNABLE_TO_PARSE_RESPONSE, http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_HTTP_UNABLE_TO_WRITE_RESPONSE, http.StatusInternalServerError)
		return
	}

}
