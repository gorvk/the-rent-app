package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorvk/rent-app/api-services/common"
	"github.com/gorvk/rent-app/api-services/common/constants"
	customTypes "github.com/gorvk/rent-app/api-services/common/types"
	models "github.com/gorvk/rent-app/api-services/models/shop"
)

func GetAllShops(w http.ResponseWriter, r *http.Request) {

	var err error
	err = common.CheckHttpResponseType(w, r, http.MethodGet)
	if err != nil {
		return
	}

	rows, err := models.GetAllShops()
	if err != nil {
		common.HandleDbError(err, w, constants.ERROR_DB_UNABLE_TO_GET_RECORD, http.StatusInternalServerError)
		return
	}

	var shops []customTypes.Shop

	defer rows.Close()
	for rows.Next() {
		row := customTypes.Shop{}
		rows.Scan(
			&row.Id,
			&row.OwnerId,
			&row.ShopName,
			&row.Email,
			&row.PhoneNumber,
			&row.MapLocation,
			&row.ShopType,
			&row.ShopDescription,
		)
		shops = append(shops, row)
	}

	var response customTypes.RESPONSE_PARAMETERS
	response.Result = shops
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
