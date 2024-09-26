package send

import (
	"dev11/lib/http/response"
	"encoding/json"
	"net/http"
)

func ErrorJSON(w http.ResponseWriter, resp response.Response) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(resp)
}

func OkJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response.OK())
}
