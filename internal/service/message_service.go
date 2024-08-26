package service

import (
    "mini-chat/internal/domain"
    "mini-chat/internal/repository"
    "mini-chat/pkg/utils"
    "encoding/json"
    "github.com/go-redis/redis/v8"
    "time"
    "github.com/google/uuid"
)

type messageService struct {
    repo repository.MessageRepository
    redisClient *redis.Client
}

func NewMessageService(repo repository.MessageRepository, redisClient *redis.Client) MessageService {
    return &messageService{repo: repo, redisClient: redisClient}
}

func (s *messageService) SendMessage(senderID, recipientID, content string) error {
    message := domain.Message{
        ID:          uuid.New().String(),
        SenderID:    senderID,
        RecipientID: recipientID,
        Content:     content,
        Timestamp:   time.Now(),
    }
    err := s.repo.SendMessage(message)
    if err != nil {
        return err
    }

    cacheKey := s.getCacheKey(senderID, recipientID)
    s.redisClient.Del(utils.Ctx, cacheKey)
    
    return nil
}

func (s *messageService) GetMessages(senderID, recipientID string) ([]domain.Message, error) {
    cacheKey := s.getCacheKey(senderID, recipientID)
    cachedMessages, err := s.redisClient.Get(utils.Ctx, cacheKey).Result()
    if err == nil && cachedMessages != "" {
        var messages []domain.Message
        err = json.Unmarshal([]byte(cachedMessages), &messages)
        if err == nil {
            return messages, nil
        }
    }

    messages, err := s.repo.GetMessages(senderID, recipientID)
    if err != nil {
        return nil, err
    }

    messagesJSON, err := json.Marshal(messages)
    if err == nil {
        s.redisClient.Set(utils.Ctx, cacheKey, messagesJSON, time.Hour).Err()
    }

    return messages, nil
}

func (s *messageService) getCacheKey(senderID, recipientID string) string {
    return "chat:" + senderID + ":" + recipientID
}
