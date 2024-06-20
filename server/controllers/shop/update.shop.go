package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorvk/rent-app/api-services/common"
	"github.com/gorvk/rent-app/api-services/common/constants"
	customTypes "github.com/gorvk/rent-app/api-services/common/types"
	shopModel "github.com/gorvk/rent-app/api-services/models/shop"
)

func UpdateShop(w http.ResponseWriter, r *http.Request) {

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

	var input customTypes.UPDATE_SHOP_INPUT
	err = json.Unmarshal(d, &input)
	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_HTTP_INVALID_REQUEST_BODY, http.StatusBadRequest)
		return
	}

	if input.Email == "" || user.Id == 0 {
		err := fmt.Errorf(constants.ERROR_HTTP_INVALID_REQUEST_BODY)
		common.HandleHttpError(err, w, constants.ERROR_HTTP_INVALID_REQUEST_BODY, http.StatusBadRequest)
		return
	}

	_, err = shopModel.UpdateShop(input, user.Id)
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
