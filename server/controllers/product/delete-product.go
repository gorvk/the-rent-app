package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorvk/rent-app/api-services/common"
	"github.com/gorvk/rent-app/api-services/common/constants"
	customTypes "github.com/gorvk/rent-app/api-services/common/types"
	productModel "github.com/gorvk/rent-app/api-services/models/product"
	shopModel "github.com/gorvk/rent-app/api-services/models/shop"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	var err error
	err = common.CheckHttpResponseType(w, r, http.MethodDelete)
	if err != nil {
		return
	}

	user, err := common.IsAuthenticated(r)
	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_HTTP_UNAUTHORIZED, http.StatusUnauthorized)
		return
	}

	d, err := io.ReadAll(r.Body)
	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_HTTP_INVALID_REQUEST_BODY, http.StatusBadRequest)
		return
	}

	var input customTypes.DELETE_PRODUCT_INPUT
	err = json.Unmarshal(d, &input)
	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_HTTP_INVALID_REQUEST_BODY, http.StatusBadRequest)
		return
	}

	if input.Id == 0 {
		err := fmt.Errorf(constants.ERROR_DB_UNABLE_TO_DELETE_RECORD)
		common.HandleDbError(err, w, constants.ERROR_DB_UNABLE_TO_DELETE_RECORD, http.StatusInternalServerError)
		return
	}

	rows, err := shopModel.GetShopByOwnerId(user.Id)
	if err != nil {
		common.HandleDbError(err, w, constants.ERROR_DB_UNABLE_TO_DELETE_RECORD, http.StatusInternalServerError)
		return
	}

	shop := customTypes.Shop{}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(
			&shop.Id,
			&shop.OwnerId,
			&shop.ShopName,
			&shop.Email,
			&shop.PhoneNumber,
			&shop.MapLocation,
			&shop.ShopType,
			&shop.ShopDescription,
		)
	}

	if shop.Id == 0 {
		err := fmt.Errorf(constants.ERROR_DB_UNABLE_TO_DELETE_RECORD)
		common.HandleDbError(err, w, constants.ERROR_DB_UNABLE_TO_DELETE_RECORD, http.StatusInternalServerError)
		return
	}

	_, err = productModel.DeleteProduct(input.Id, shop.Id)
	if err != nil {
		common.HandleDbError(err, w, constants.ERROR_DB_UNABLE_TO_DELETE_RECORD, http.StatusInternalServerError)
		return
	}

	var response customTypes.RESPONSE_PARAMETERS
	response.Result = nil
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
