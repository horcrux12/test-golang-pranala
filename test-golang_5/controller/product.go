package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"test_golang_5/app"
	"test_golang_5/dto/out"
	"test_golang_5/model"

	"github.com/gorilla/mux"
)

const (
	BadRequestMessage = "Bad request"
	BadgatewayMessage = "Bad gateway"
	SuccessMessage = "Success"
)

func GetProducts(wr http.ResponseWriter, req *http.Request) {
	var resp []model.Product
	err := app.AppAtrribute.DB.Model(&model.Product{}).Find(&resp).Error
	if err != nil {
		writeResponse(out.WebResponse{
			Message: BadgatewayMessage,
		}, http.StatusBadGateway, wr)
		return
	}

	writeResponse(out.WebResponse{
		Message: SuccessMessage,
	}, http.StatusOK, wr)
}

func CreateProducts(wr http.ResponseWriter, req *http.Request) {
	var payload model.Product
	err := json.NewDecoder(req.Body).Decode(&payload)
	if err != nil {
		writeResponse(out.WebResponse{
			Message: BadRequestMessage,
		}, http.StatusBadRequest, wr)
		return
	}

	err = app.AppAtrribute.DB.Create(&payload).Error
	if err != nil {
		writeResponse(out.WebResponse{
			Message: BadgatewayMessage,
		}, http.StatusBadGateway, wr)
		return
	}

	writeResponse(out.WebResponse{
		Message: SuccessMessage,
	}, http.StatusOK, wr)
}

func UpdateProducts(wr http.ResponseWriter, req *http.Request) {
	var payload model.Product
	vars := mux.Vars(req)
	err := json.NewDecoder(req.Body).Decode(&payload)
	if err != nil {
		writeResponse(out.WebResponse{
			Message: BadRequestMessage,
		}, http.StatusBadRequest, wr)
		return
	}
	payload.ID, _ = strconv.Atoi(vars["ID"])

	err = app.AppAtrribute.DB.Save(&payload).Error
	if err != nil {
		writeResponse(out.WebResponse{
			Message: BadgatewayMessage,
		}, http.StatusBadGateway, wr)
		return
	}

	writeResponse(out.WebResponse{
		Message: SuccessMessage,
	}, http.StatusOK, wr)
}

func DeleteProducts(wr http.ResponseWriter, req *http.Request) {
	var payload model.Product
	vars := mux.Vars(req)
	payload.ID, _ = strconv.Atoi(vars["ID"])

	err := app.AppAtrribute.DB.Delete(&payload).Error
	if err != nil {
		writeResponse(out.WebResponse{
			Message: BadgatewayMessage,
		}, http.StatusBadGateway, wr)
		return
	}

	writeResponse(out.WebResponse{
		Message: SuccessMessage,
	}, http.StatusOK, wr)
}

func writeResponse(data interface{}, status int, wr http.ResponseWriter) {
	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(status)
	res, _ := json.Marshal(&data)
	wr.Write(res)
}