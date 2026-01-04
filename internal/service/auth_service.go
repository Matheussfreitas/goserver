package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"goserver/internal/domain"
	"goserver/internal/repository"
	"goserver/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	db             *sql.DB
	userRepository *repository.UserRepository
}

func NewAuthService(db *sql.DB, userRepository *repository.UserRepository) *AuthService {
	return &AuthService{
		db:             db,
		userRepository: userRepository,
	}
}

var (
	ErrUserNotFound      = errors.New("usuário não encontrado")
	ErrUserAlreadyExists = errors.New("usuário já cadastrado")
)

func (s *AuthService) Login(ctx context.Context, email, password string) (*domain.User, string, error) {
	user, err := s.userRepository.FindUserByEmail(ctx, nil, email)
	if err != nil {
		return nil, "", err
	}
	if user == nil {
		return nil, "", ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", ErrUserNotFound
	}

	token, err := utils.GenerateToken(user.ID.String())
	if err != nil {
		return nil, "", err
	}

	fmt.Println("Usuário autenticado com sucesso")

	return user, token, nil
}

func (s *AuthService) Register(ctx context.Context, email, password string) (*domain.User, error) {
	existingUser, err := s.userRepository.FindUserByEmail(ctx, nil, email)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, ErrUserAlreadyExists
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := domain.User{
		Email:    email,
		Password: string(hashPassword),
		Active:   true,
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	err = s.userRepository.CreateUser(ctx, tx, &newUser)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	fmt.Println("Usuário cadastrado com sucesso")
	return &newUser, nil
}
