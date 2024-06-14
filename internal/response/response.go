package response

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func Json(w http.ResponseWriter, status int, resp ApiResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)

}

func SuccessResponse(data interface{}) ApiResponse {
	return ApiResponse{
		Success: true,
		Data:    data,
	}
}
func ErrorResponse(err string) ApiResponse {
	return ApiResponse{
		Success: false,
		Error:   err,
	}
}
