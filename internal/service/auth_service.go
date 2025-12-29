package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"goserver/internal/domain"
	"goserver/internal/repository"

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

func (s *AuthService) Login(ctx context.Context, email, password string) error {
	user, err := s.userRepository.FindUserByEmail(ctx, nil, email)
	if err != nil {
		return err
	}
	if user == nil {
		return ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return ErrUserNotFound
	}

	fmt.Println("Usuário autenticado com sucesso")
	return nil
}

func (s *AuthService) Register(ctx context.Context, email, password string) error {
	existingUser, err := s.userRepository.FindUserByEmail(ctx, nil, email)
	if err != nil {
		return err
	}

	if existingUser != nil {
		return ErrUserAlreadyExists
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := domain.User{
		Email:    email,
		Password: string(hashPassword),
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = s.userRepository.CreateUser(ctx, tx, newUser)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	fmt.Println("Usuário cadastrado com sucesso")
	return nil
}
