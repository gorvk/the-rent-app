package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorvk/rent-app/api-services/common"
	"github.com/gorvk/rent-app/api-services/common/constants"
	customTypes "github.com/gorvk/rent-app/api-services/common/types"
	productModel "github.com/gorvk/rent-app/api-services/models/product"
	shopModel "github.com/gorvk/rent-app/api-services/models/shop"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	var err error
	err = common.CheckHttpResponseType(w, r, http.MethodPut)
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

	var input customTypes.Product
	err = json.Unmarshal(d, &input)
	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_HTTP_INVALID_REQUEST_BODY, http.StatusBadRequest)
		return
	}

	row, err := shopModel.GetShopByOwnerId(user.Id)
	if err != nil {
		common.HandleDbError(err, w, constants.ERROR_DB_UNABLE_TO_UPDATE_RECORD, http.StatusInternalServerError)
		return
	}

	shop := customTypes.Shop{}
	defer row.Close()
	for row.Next() {
		row.Scan(
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

	if shop.OwnerId != user.Id {
		common.HandleDbError(err, w, constants.ERROR_HTTP_ACCESS_DENIED, http.StatusForbidden)
		return
	}

	err = productModel.UpdateProduct(input)
	if err != nil {
		common.HandleDbError(err, w, constants.ERROR_DB_UNABLE_TO_CREATE_RECORD, http.StatusInternalServerError)
		return
	}

	data, err := common.ConstructResponse(true, nil)
	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_HTTP_UNABLE_TO_PARSE_RESPONSE, http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
