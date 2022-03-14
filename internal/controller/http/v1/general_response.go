package v1

import (
	"encoding/json"
	"net/http"
)

/*
type GeneralResponse struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"status"`
	ErrMsg     string      `json:"err_msg"`
	Data       interface{} `json:"data"`
}
*/

type GeneralResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type ResponseErrorCodeAndMessage struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_msg,omitempty"`
}

func SendResponse(w http.ResponseWriter, data interface{}, errCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if errCode == 0 {
		errCode = http.StatusOK
	}
	w.WriteHeader(errCode)

	errorMsg := http.StatusText(errCode)

	var resp GeneralResponse
	if errCode == http.StatusOK {
		if data == nil {
			resp = GeneralResponse{
				Success: true,
			}
		} else {
			resp = GeneralResponse{
				Success: true,
				Data:    data,
			}
		}
	} else {
		resp = GeneralResponse{
			Success: false,
			Data: ResponseErrorCodeAndMessage{
				ErrorCode:    errCode,
				ErrorMessage: errorMsg,
			},
		}
	}

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		panic(err)
	}
}
