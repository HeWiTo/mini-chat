package service

import (
	"errors"
	"mini-chat/internal/domain"
    "mini-chat/internal/repository"
    "github.com/dgrijalva/jwt-go"
    "time"
)

type AuthService interface {
    Register(username, password string) (string, error)
    Login(username, password string) (string, error)
}

type authService struct {
    repo repository.UserRepository
    secretKey string
}

func NewAuthService(repo repository.UserRepository, secretKey string) AuthService {
    return &authService{repo: repo, secretKey: secretKey}
}

func (s *authService) Register(username, password string) (string, error) {
    user := domain.User{
        ID:       "generate-unique-id", // This will be replaced by UUID generator or similar
        Username: username,
        Password: hashPassword(password),
    }
    err := s.repo.CreateUser(user)
    if err != nil {
        return "", err
    }
    token, err := s.generateToken(user)
    if err != nil {
        return "", err
    }
    return token, nil
}

func (s *authService) Login(username, password string) (string, error) {
    user, err := s.repo.GetUserByUsername(username)
    if err != nil {
        return "", err
    }
    if !checkPasswordHash(password, user.Password) {
        return "", errors.New("invalid credentials")
    }
    token, err := s.generateToken(user)
    if err != nil {
        return "", err
    }
    return token, nil
}

func (s *authService) generateToken(user domain.User) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    })
    tokenString, err := token.SignedString([]byte(s.secretKey))
    if err != nil {
        return "", err
    }
    return tokenString, nil
}