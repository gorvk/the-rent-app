package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorvk/rent-app/server/api-services/common"
	"github.com/gorvk/rent-app/server/api-services/common/constants"
	customTypes "github.com/gorvk/rent-app/server/api-services/common/types"
	models "github.com/gorvk/rent-app/server/api-services/models/user"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	var err error
	err = common.CheckHttpResponseType(w, r, http.MethodGet)
	if err != nil {
		return
	}

	_, err = common.IsAuthenticated(r)
	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_HTTP_UNAUTHORIZED, http.StatusUnauthorized)
		return
	}

	rows, err := models.GetAllUsers()
	if err != nil {
		common.HandleDbError(err, w, constants.ERROR_DB_UNABLE_TO_GET_RECORD, http.StatusInternalServerError)
		return
	}

	var users []customTypes.User

	defer rows.Close()
	for rows.Next() {
		userRow := customTypes.User{}
		rows.Scan(
			&userRow.Id,
			&userRow.FirstName,
			&userRow.LastName,
			&userRow.Email,
			&userRow.PhoneNumber,
			&userRow.UserAddress,
			&userRow.IsShopEnabled,
			&userRow.AccountPassword,
		)
		users = append(users, userRow)
	}

	var response customTypes.RESPONSE_PARAMETERS
	response.Result = users
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
