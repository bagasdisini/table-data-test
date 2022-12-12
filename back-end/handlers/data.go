package handlers

import (
	datadto "backend/dto/data"
	dto "backend/dto/result"
	"backend/models"
	"backend/repositories"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerData struct {
	DataRepository repositories.DataRepository
}

func HandlerData(DataRepository repositories.DataRepository) *handlerData {
	return &handlerData{DataRepository}
}

func (h *handlerData) ShowData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := h.DataRepository.ShowData()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{StatusCode: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	status, err := h.DataRepository.ShowStatus()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{StatusCode: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: status, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerData) GetDataByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var data models.Data
	data, err := h.DataRepository.GetDataByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{StatusCode: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var status models.Status
	status, err2 := h.DataRepository.GetStatusByID(id)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{StatusCode: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: status, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerData) CreateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	x := r.FormValue("amount")
	y := r.FormValue("statusID")

	amount, _ := strconv.Atoi(x)
	statusID, _ := strconv.Atoi(y)

	request := datadto.DataRequest{
		ProductName:  r.FormValue("productName"),
		Amount:       amount,
		CustomerName: r.FormValue("customerName"),
		StatusID:     statusID,
		CreateBy:     r.FormValue("createBy"),
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{StatusCode: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	productID := time.Now().UnixNano() / int64(time.Millisecond)
	transactionDate := time.Now()

	data := models.Data{
		ProductID:       productID,
		ProductName:     request.ProductName,
		Amount:          request.Amount,
		CustomerName:    request.CustomerName,
		StatusID:        request.StatusID,
		TransactionDate: transactionDate,
		CreateBy:        request.CreateBy,
		CreateOn:        transactionDate,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{StatusCode: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err = h.DataRepository.CreateData(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{StatusCode: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, _ = h.DataRepository.GetDataByID(data.ID)

	var status models.Status
	status, err2 := h.DataRepository.GetStatusByID(data.StatusID)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{StatusCode: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: status, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerData) UpdateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	x := r.FormValue("amount")
	y := r.FormValue("statusID")

	amount, _ := strconv.Atoi(x)
	statusID, _ := strconv.Atoi(y)

	request := datadto.DataRequest{
		ProductName:  r.FormValue("productName"),
		Amount:       amount,
		CustomerName: r.FormValue("customerName"),
		StatusID:     statusID,
		CreateBy:     r.FormValue("createBy"),
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{StatusCode: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data := models.Data{}

	if request.ProductName != "" {
		data.ProductName = request.ProductName
	}

	if request.Amount != 0 {
		data.Amount = request.Amount
	}

	if request.CustomerName != "" {
		data.CustomerName = request.CustomerName
	}

	if request.StatusID != 0 {
		data.StatusID = request.StatusID
	}

	if request.CreateBy != "" {
		data.CreateBy = request.CreateBy
	}

	data, err := h.DataRepository.UpdateData(data, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{StatusCode: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var status models.Status
	status, err2 := h.DataRepository.GetStatusByID(id)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{StatusCode: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: status, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerData) DeleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	data, err := h.DataRepository.GetDataByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{StatusCode: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err2 := h.DataRepository.DeleteData(data, id)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{StatusCode: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var status models.Status
	status, err3 := h.DataRepository.GetStatusByID(data.StatusID)
	if err3 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{StatusCode: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: status, Data: data}
	json.NewEncoder(w).Encode(response)
}
