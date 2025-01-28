package handler

import (
	"encoding/json"
	"net/http"
	"realmrovers/services"
	"strconv"
	"strings"
)

type UserHandler struct {
	Service *services.UserService
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	idParam := pathParts[2]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.Service.GetUsers(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Failed to marshal user", http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
