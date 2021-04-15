package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/saviobarr/golang-microservices/mvc/services"
	"github.com/saviobarr/golang-microservices/mvc/utils"
)

//fetch user
func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)

	resp.Header().Add("Content-Type", "application/json")

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)

		return
	}

	user, apiErr := services.GetUser(userId)

	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)

	resp.Write(jsonValue)

}
