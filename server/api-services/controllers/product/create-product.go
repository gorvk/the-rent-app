package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorvk/rent-app/server/api-services/common"
	"github.com/gorvk/rent-app/server/api-services/common/constants"
	customTypes "github.com/gorvk/rent-app/server/api-services/common/types"
	productModel "github.com/gorvk/rent-app/server/api-services/models/product"
	shopModel "github.com/gorvk/rent-app/server/api-services/models/shop"
	userModel "github.com/gorvk/rent-app/server/api-services/models/user"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var err error
	err = common.CheckHttpResponseType(w, r, http.MethodPost)
	if err != nil {
		return
	}

	token, err := common.IsAuthenticated(r)
	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_HTTP_UNAUTHORIZED, http.StatusUnauthorized)
		return
	}

	claims := token.Claims.(*jwt.RegisteredClaims)
	rows, err := userModel.GetUserByEmail(claims.Issuer)
	if err != nil {
		common.HandleDbError(err, w, constants.ERROR_DB_UNABLE_TO_GET_RECORD, http.StatusInternalServerError)
		return
	}

	user := customTypes.User{}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.PhoneNumber,
			&user.UserAddress,
			&user.IsShopEnabled,
			&user.AccountPassword,
		)
	}

	d, err := io.ReadAll(r.Body)
	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_HTTP_INVALID_REQUEST_BODY, http.StatusBadRequest)
		return
	}

	var input customTypes.CREATE_PRODUCT_INPUT
	err = json.Unmarshal(d, &input)
	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_HTTP_INVALID_REQUEST_BODY, http.StatusBadRequest)
		return
	}

	rows, err = shopModel.GetShopByOwnerId(user.Id)
	if err != nil {
		common.HandleDbError(err, w, constants.ERROR_DB_UNABLE_TO_GET_RECORD, http.StatusInternalServerError)
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
			&shop.City,
			&shop.Country,
		)
	}

	if shop.Id == 0 || user.Id == 0 || input.ProductName == "" {
		err := fmt.Errorf(constants.ERROR_DB_UNABLE_TO_CREATE_RECORD)
		common.HandleHttpError(err, w, constants.ERROR_HTTP_INVALID_REQUEST_BODY, http.StatusBadRequest)
		return
	}

	_, err = productModel.CreateProduct(input, shop.Id)
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
