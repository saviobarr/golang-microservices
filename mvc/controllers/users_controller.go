package controllers

import (
	"encoding/json"
	"github.com/saviobarr/golang-microservices/mvc/services"
	"github.com/saviobarr/golang-microservices/mvc/utils"
	"log"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		log.Println(err)
		return
	}
	user, apiError := services.GetUser(userId)

	if apiError != nil {
		resp.WriteHeader(apiError.StatusCode)
		resp.Write([]byte(apiError.Message))
	}
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
