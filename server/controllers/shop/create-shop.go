package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gorvk/rent-app/api-services/common"
	"github.com/gorvk/rent-app/api-services/common/constants"
	customTypes "github.com/gorvk/rent-app/api-services/common/types"
	shopModel "github.com/gorvk/rent-app/api-services/models/shop"
	userModel "github.com/gorvk/rent-app/api-services/models/user"
)

func CreateShop(w http.ResponseWriter, r *http.Request) {
	var err error
	err = common.CheckHttpResponseType(w, r, http.MethodPost)
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

	var input customTypes.CREATE_SHOP_INPUT
	err = json.Unmarshal(d, &input)
	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_HTTP_INVALID_REQUEST_BODY, http.StatusBadRequest)
		return
	}

	if input.Email == "" || user.Email == "" {
		common.HandleHttpError(err, w, constants.ERROR_HTTP_INVALID_REQUEST_BODY, http.StatusBadRequest)
		return
	}

	user.IsShopEnabled = true
	var updateUserInput customTypes.UPDATE_USER_INPUT
	err = common.CastStructs(user, &updateUserInput)

	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_HTTP_INVALID_REQUEST_BODY, http.StatusBadRequest)
		return
	}

	_, err = userModel.UpdateUser(customTypes.UPDATE_USER_INPUT(updateUserInput))
	if err != nil {
		common.HandleDbError(err, w, constants.ERROR_DB_UNABLE_TO_CREATE_RECORD, http.StatusInternalServerError)
		return
	}

	input.Email = strings.ToLower(input.Email)

	_, err = shopModel.CreateNewShop(input, user.Id)
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
