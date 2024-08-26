package transport

import (
    "encoding/json"
    "net/http"
    "mini-chat/internal/service"
)

type MessageHandler struct {
    messageService service.MessageService
}

func NewMessageHandler(svc service.MessageService) *MessageHandler {
    return &MessageHandler{messageService: svc}
}

func (h *MessageHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
    var req struct {
        SenderID    string `json:"sender_id"`
        RecipientID string `json:"recipient_id"`
        Content     string `json:"content"`
    }
    json.NewDecoder(r.Body).Decode(&req)
    err := h.messageService.SendMessage(req.SenderID, req.RecipientID, req.Content)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func (h *MessageHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
    senderID := r.URL.Query().Get("sender_id")
    recipientID := r.URL.Query().Get("recipient_id")

    messages, err := h.messageService.GetMessages(senderID, recipientID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(messages)
}