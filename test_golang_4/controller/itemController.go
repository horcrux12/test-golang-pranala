package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"test_golang_4/model"
)

func ItemController(wr http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet :
		GetItems(wr, req)
	case http.MethodPost:
		PostItem(wr, req)
	default:
		wr.WriteHeader(http.StatusNotFound)
		wr.Write([]byte("Not Found"))
	}
}

func PostItem(wr http.ResponseWriter, req *http.Request) {
	var (
		items []model.Item
		payload model.Item
		bodyBytes, data, newData []byte
	)

	//  Get data from database
	data, err := os.ReadFile("./data.json")
	if err != nil {
		writeResponse("Bad gateway", http.StatusBadGateway, wr)
		return
	}

	err = json.Unmarshal(data, &items)
	if err != nil {
		fmt.Println("Unmarshal Data 1", err)
		writeResponse("Bad gateway", http.StatusBadGateway, wr)
		return
	}

	// get payload
	if req.Body != nil {
		bodyBytes, err = ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println(err)
			writeResponse("Bad gateway", http.StatusBadGateway, wr)
			return
		}

		err = json.Unmarshal(bodyBytes, &payload)
		if err != nil {
			fmt.Println("Unmarshal Data 2", err)
			writeResponse("Bad gateway", http.StatusBadGateway, wr)
			return
		}
	}
	
	currID := items[len(items) - 1].Id
	currID += 1

	payload.Id = currID

	items = append(items, payload)
	
	newData, err = json.Marshal(items)
	if err != nil {
		fmt.Println("marshal Data 1", err)
		writeResponse("Bad gateway", http.StatusBadGateway, wr)
		return
	}

	err = os.WriteFile("./data.json", newData, 0644)
	if err != nil {
		fmt.Println(err)
		writeResponse("Bad gateway", http.StatusBadGateway, wr)
		return
	}

	writeResponse(map[string]interface{}{
		"status" : "Success",
		"data" : payload,
	}, http.StatusOK, wr)
}

func GetItems(wr http.ResponseWriter, req *http.Request) {
	var items []model.Item
	data, err := os.ReadFile("./data.json")
	if err != nil {
		writeResponse("Bad gateway", http.StatusBadGateway, wr)
		return
	}

	err = json.Unmarshal(data, &items)
	if err != nil {
		writeResponse("Bad gateway", http.StatusBadGateway, wr)
		return
	}

	writeResponse(map[string]interface{}{
		"data" : items,
	}, http.StatusOK, wr)
}

func writeResponse(data interface{}, status int, wr http.ResponseWriter) {
	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(status)
	res, _ := json.Marshal(&data)
	wr.Write(res)
}
