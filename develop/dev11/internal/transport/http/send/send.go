package send

import (
	"dev11/lib/http/response"
	"encoding/json"
	"net/http"
)

func SendJSON(w http.ResponseWriter, resp response.Response) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(resp)
}
