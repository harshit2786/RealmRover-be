package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"realmrovers/services"
	"strconv"
	"strings"
)

type RequestData struct {
	Code  string `json:"code"`
}

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

func (h *UserHandler) SignUser(w http.ResponseWriter , r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	var requestData RequestData
	if err := json.Unmarshal(body,&requestData); err!=nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	access_token,err := h.Service.SignUser(requestData.Code)
	if err!= nil {
		http.Error(w,"Failed to get access token", http.StatusForbidden)
	}
	type access struct {
		Access_token string
	}
	var token = access{Access_token: access_token}
	response,err := json.Marshal(token)
	if err != nil {
		http.Error(w, "Failed to marshal user", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type" ,"application/json")
	w.WriteHeader(http.StatusOK);
	w.Write(response)
}