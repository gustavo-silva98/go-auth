package handler

import (
	"encoding/json"
	"go-auth/internal/db"
	"net/http"
)

type RoleBackend struct {
	db db.RoleRepo
}

func (rb *RoleBackend) GetRoles(w http.ResponseWriter, r *http.Request) {
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
