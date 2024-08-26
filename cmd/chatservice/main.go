package main

import (
    "mini-chat/internal/repository"
    "mini-chat/internal/service"
    "mini-chat/pkg/utils"
    "mini-chat/transport"
    "github.com/gocql/gocql"
    "log"
    "net/http"
    "os"
)

func main() {
    cassandraHost := os.Getenv("CASSANDRA_HOST")
    redisHost := os.Getenv("REDIS_HOST")

    // Cassandra session setup
    cluster := gocql.NewCluster(cassandraHost)
    cluster.Keyspace = "chat"
    session, err := cluster.CreateSession()
    if err != nil {
        log.Fatal(err)
    }
    defer session.Close()

    // Redis client setup
    redisClient := utils.NewRedisClientWithHost(redisHost)

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
