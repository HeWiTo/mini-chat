package repository

import (
    "mini-chat/internal/domain"
    "github.com/gocql/gocql"
)

type UserRepository interface {
    CreateUser(user domain.User) error
    GetUserByUsername(username string) (domain.User, error)
}

type cassandraUserRepository struct {
    session *gocql.Session
}

func NewCassandraUserRepository(session *gocql.Session) UserRepository {
    return &cassandraUserRepository{session: session}
}

func (r *cassandraUserRepository) CreateUser(user domain.User) error {
    query := `INSERT INTO users (id, username, password) VALUES (?, ?, ?)`
    return r.session.Query(query, user.ID, user.Username, user.Password).Exec()
}

func (r *cassandraUserRepository) GetUserByUsername(username string) (domain.User, error) {
    var user domain.User
    query := `SELECT id, username, password FROM users WHERE username = ? LIMIT 1`
    err := r.session.Query(query, username).Scan(&user.ID, &user.Username, &user.Password)
    return user, err
}
