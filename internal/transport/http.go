package transport

import (
    "encoding/json"
    "net/http"
    "mini-chat/internal/service"
)

type AuthHandler struct {
    authService service.AuthService
}

func NewAuthHandler(svc service.AuthService) *AuthHandler {
    return &AuthHandler{authService: svc}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    json.NewDecoder(r.Body).Decode(&req)
    token, err := h.authService.Register(req.Username, req.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    json.NewDecoder(r.Body).Decode(&req)
    token, err := h.authService.Login(req.Username, req.Password)
    if err != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}
