package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorvk/rent-app/api-services/common"
	"github.com/gorvk/rent-app/api-services/common/constants"
	customTypes "github.com/gorvk/rent-app/api-services/common/types"
	orderModels "github.com/gorvk/rent-app/api-services/models/order"
	productModels "github.com/gorvk/rent-app/api-services/models/product"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
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

	var input customTypes.PLACE_ORDER_INPUT
	err = json.Unmarshal(d, &input)
	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_HTTP_INVALID_REQUEST_BODY, http.StatusBadRequest)
		return
	}

	// check if user placing the order is actual user that is loggedin
	if user.Id != input.BuyerId {
		common.HandleHttpError(err, w, constants.ERROR_HTTP_UNAUTHORIZED, http.StatusBadRequest)
		return
	}

	rows, err := productModels.GetProductById(input.ProductId)
	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_DB_UNABLE_TO_GET_RECORD, http.StatusBadRequest)
		return
	}

	var product customTypes.GET_PRODUCT_BY_ID_OUTPUT
	defer rows.Close()
	for rows.Next() {
		rows.Scan(
			&product.Id,
			&product.ProductName,
			&product.ProductType,
			&product.ProductCondition,
			&product.Price,
			&product.OriginalPurchasedDate,
			&product.OriginalPurchaisingRecieptNo,
			&product.ProductDescription,
			&product.ShopId,
			&product.ShopName,
			&product.City,
			&product.Country,
			&product.Email,
			&product.PhoneNumber,
			&product.MapLocation,
			&product.ShopType,
			&product.ShopDescription,
		)
	}

	// check if product is valid or not using ProductId
	// check if quantity in input is <= quantity of product
	if product.Id != 0 || product.Id != input.ProductId || product.Quantity != input.Quantity {
		common.HandleHttpError(err, w, constants.ERROR_DB_UNABLE_TO_GET_RECORD, http.StatusBadRequest)
		return
	}

	// create a record for placed order in orders table using input
	err = orderModels.CreateOrder(input)
	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_DB_UNABLE_TO_CREATE_RECORD, http.StatusInternalServerError)
		return
	}

	// decrease the count of product's quantity using ProductId in Products detail
	product.Quantity = product.Quantity - input.Quantity
	updatedProduct := customTypes.Product{
		Id:                           product.Id,
		ProductName:                  product.ProductName,
		ShopId:                       product.ShopId,
		ProductType:                  product.ProductType,
		ProductCondition:             product.ProductCondition,
		ProductDescription:           product.ProductDescription,
		Price:                        product.Price,
		OriginalPurchasedDate:        product.OriginalPurchasedDate,
		OriginalPurchaisingRecieptNo: product.OriginalPurchaisingRecieptNo,
	}
	err = productModels.UpdateProduct(updatedProduct)
	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_DB_UNABLE_TO_UPDATE_RECORD, http.StatusInternalServerError)
		return
	}
}
