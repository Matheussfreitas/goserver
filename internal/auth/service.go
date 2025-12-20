package auth

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound      = errors.New("usuário não encontrado")
	ErrUserAlreadyExists = errors.New("usuário já cadastrado")
)

type AuthService struct {
	userRepository []User
}

type User struct {
	Email    string
	Password string
}

func NewAuthService() *AuthService {
	return &AuthService{
		userRepository: []User{},
	}
}

func (s *AuthService) Login(email, password string) error {
	for _, user := range s.userRepository {
		if user.Email == email {
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
			if err != nil {
				return ErrUserNotFound
			}
			fmt.Println("Usuário autenticado com sucesso")
			return nil
		}
	}
	return ErrUserNotFound
}

func (s *AuthService) Register(email, password string) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	for _, user := range s.userRepository {
		if user.Email == email {
			return ErrUserAlreadyExists
		}
	}

	s.userRepository = append(s.userRepository, User{Email: email, Password: string(hashPassword)})
	fmt.Println("Usuário cadastrado com sucesso")
	return nil
}
