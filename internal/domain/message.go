package domain

import "time"

type Message struct {
    ID          string
    SenderID    string
    RecipientID string
    Content     string
    Timestamp   time.Time
}
