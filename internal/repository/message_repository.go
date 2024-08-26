package repository

import (
    "mini-chat/internal/domain"
    "github.com/gocql/gocql"
    "time"
)

type MessageRepository interface {
    SendMessage(message domain.Message) error
    GetMessages(senderID, recipientID string) ([]domain.Message, error)
}

type cassandraMessageRepository struct {
    session *gocql.Session
}

func NewCassandraMessageRepository(session *gocql.Session) MessageRepository {
    return &cassandraMessageRepository{session: session}
}

func (r *cassandraMessageRepository) SendMessage(message domain.Message) error {
    query := `INSERT INTO messages (id, sender_id, recipient_id, content, timestamp) VALUES (?, ?, ?, ?, ?)`
    return r.session.Query(query, message.ID, message.SenderID, message.RecipientID, message.Content, message.Timestamp).Exec()
}

func (r *cassandraMessageRepository) GetMessages(senderID, recipientID string) ([]domain.Message, error) {
    var messages []domain.Message
    query := `SELECT id, sender_id, recipient_id, content, timestamp FROM messages WHERE sender_id = ? AND recipient_id = ?`
    iter := r.session.Query(query, senderID, recipientID).Iter()

    var msg domain.Message
    for iter.Scan(&msg.ID, &msg.SenderID, &msg.RecipientID, &msg.Content, &msg.Timestamp) {
        messages = append(messages, msg)
    }
    if err := iter.Close(); err != nil {
        return nil, err
    }
    return messages, nil
}
