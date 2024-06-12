package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gorvk/rent-app/api-services/common"
	"github.com/gorvk/rent-app/api-services/common/constants"
	customTypes "github.com/gorvk/rent-app/api-services/common/types"
	productModel "github.com/gorvk/rent-app/api-services/models/product"
)

func SearchProducts(w http.ResponseWriter, r *http.Request) {

	var err error
	err = common.CheckHttpResponseType(w, r, http.MethodPost)
	if err != nil {
		return
	}

	d, err := io.ReadAll(r.Body)
	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_HTTP_INVALID_REQUEST_BODY, http.StatusBadRequest)
		return
	}

	var input customTypes.SEARCH_PRODUCTS_INPUT
	err = json.Unmarshal(d, &input)
	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_HTTP_INVALID_REQUEST_BODY, http.StatusBadRequest)
		return
	}

	input.SearchTerm = strings.Replace(input.SearchTerm, " ", "&", -1)

	rows, err := productModel.SearchProducts(input.SearchTerm)
	if err != nil {
		common.HandleDbError(err, w, constants.ERROR_DB_UNABLE_TO_GET_RECORD, http.StatusInternalServerError)
		return
	}

	var products []customTypes.GET_SEARCHED_PRODUCTS_OUTPUT = []customTypes.GET_SEARCHED_PRODUCTS_OUTPUT{}

	defer rows.Close()
	for rows.Next() {
		row := customTypes.GET_SEARCHED_PRODUCTS_OUTPUT{}
		rows.Scan(
			&row.ProductName,
			&row.ProductType,
			&row.ProductCondition,
			&row.Price,
			&row.ProductDescription,
			&row.ShopId,
			&row.ShopName,
			&row.City,
			&row.Country,
			&row.Document,
			&row.SsQuery,
			&row.RankProductName,
			&row.Similarity,
		)
		products = append(products, row)
	}

	var response customTypes.RESPONSE_PARAMETERS
	if products[0].Similarity == "0" && products[0].RankProductName == "0" {
		var r = [0]customTypes.GET_SEARCHED_PRODUCTS_OUTPUT{}
		response.Result = r
	} else {
		response.Result = products
	}
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
