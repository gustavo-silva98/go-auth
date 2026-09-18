package handler

import (
	"encoding/json"
	"go-auth/internal/db"
	"io"
	"log"
	"net/http"
)

type RoleBackend struct {
	db db.RoleRepo
}

func (rb *RoleBackend) GetRole(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("ID")
	role, err := rb.db.GetRole(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := getRoleResponse{
		Id:          role.Id,
		Name:        role.Name,
		Permissions: role.Permissions,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
	w.WriteHeader(http.StatusOK)
}

func (rb *RoleBackend) PutRole(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var bodyjson getRoleResponse
	json.Unmarshal(body, &bodyjson)
	if bodyjson.Id == "" || bodyjson.Name == "" || len(bodyjson.Permissions) == 0 {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	role := db.Role{
		Id:          bodyjson.Id,
		Name:        bodyjson.Name,
		Permissions: bodyjson.Permissions,
	}

	if err := rb.db.PutRole(r.Context(), role); err != nil {
		log.Printf("Erro ao inserir Role %v - Err: %v", role.Name, err)
		http.Error(w, "Error Putting Role", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
