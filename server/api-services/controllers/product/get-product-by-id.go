package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorvk/rent-app/server/api-services/common"
	"github.com/gorvk/rent-app/server/api-services/common/constants"
	customTypes "github.com/gorvk/rent-app/server/api-services/common/types"
	productModel "github.com/gorvk/rent-app/server/api-services/models/product"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {

	var err error
	err = common.CheckHttpResponseType(w, r, http.MethodDelete)
	if err != nil {
		return
	}

	d, err := io.ReadAll(r.Body)
	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_HTTP_INVALID_REQUEST_BODY, http.StatusBadRequest)
		return
	}

	var input customTypes.GET_PRODUCT_BY_ID_INPUT
	err = json.Unmarshal(d, &input)
	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_HTTP_INVALID_REQUEST_BODY, http.StatusBadRequest)
		return
	}

	rows, err := productModel.GetProductById(input.Id)
	if err != nil {
		common.HandleDbError(err, w, constants.ERROR_DB_UNABLE_TO_GET_RECORD, http.StatusInternalServerError)
		return
	}

	product := customTypes.Product{}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(
			&product.Id,
			&product.ProductName,
			&product.ShopId,
			&product.ProductType,
			&product.ProductCondition,
			&product.Price,
			&product.OriginalPurchasedDate,
			&product.OriginalPurchaisingRecieptNo,
			&product.ProductDescription,
		)
	}

	var response customTypes.RESPONSE_PARAMETERS
	response.Result = product
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
