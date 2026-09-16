package handler

import (
	"encoding/json"
	"go-auth/internal/db"
	"net/http"
)

type RoleBackend struct {
	db db.RoleDB
}

func (rb *RoleBackend) GetRoles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	q := r.URL.Query()
	if q["roleId"] == "" {
	}
	roles, err := rb.db.GetRole(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(role)
}
application/json")
	json.NewEncoder(w).Encode(role)
}
