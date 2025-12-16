package api

import (
	"encoding/json"
	"go/femm_server/data"
	"net/http"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// api/exhibitons?id=34
	id := r.URL.Query()["id"]
	if id != nil { // we will try to serve only one
		finalId, err := strconv.Atoi(id[0])
		if err == nil && finalId < len(data.GetAll()) {
			json.NewEncoder(w).Encode(data.GetAll()[finalId])

		} else {
			http.Error(w, "Invalid Exhibitions", http.StatusBadRequest)
		}

	} else { // we want all
		json.NewEncoder(w).Encode(data.GetAll())
	}
}
