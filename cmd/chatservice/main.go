package main

import (
    "chat-system/internal/repository"
    "chat-system/internal/service"
    "chat-system/pkg/utils"
    "chat-system/transport"
    "github.com/gocql/gocql"
    "log"
    "net/http"
)

func main() {
    // Cassandra session setup
    cluster := gocql.NewCluster("127.0.0.1")
    cluster.Keyspace = "chat"
    session, err := cluster.CreateSession()
    if err != nil {
        log.Fatal(err)
    }
    defer session.Close()

    // Redis client setup
    redisClient := utils.NewRedisClient()

    // Repository and service initialization
    userRepo := repository.NewCassandraUserRepository(session)
    authService := service.NewAuthService(userRepo, "my-secret-key")

    messageRepo := repository.NewCassandraMessageRepository(session)
    messageService := service.NewMessageService(messageRepo, redisClient)

    // Handlers setup
    authHandler := transport.NewAuthHandler(authService)
    messageHandler := transport.NewMessageHandler(messageService)

    // Routes
    http.HandleFunc("/register", authHandler.Register)
    http.HandleFunc("/login", authHandler.Login)
    http.HandleFunc("/send", messageHandler.SendMessage)
    http.HandleFunc("/messages", messageHandler.GetMessages)

    log.Println("Server started at :8080")
    http.ListenAndServe(":8080", nil)
}