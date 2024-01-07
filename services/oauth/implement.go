package oauth

import (
	"errors"
	"identify-provider/utils"
	"math/rand"
	"strings"
	"time"
)

var (
	ErrUserNotFound           = errors.New("User not found")
	ErrInvalidUserPassword    = errors.New("Invalid user password")
	ErrCannotSetEmptyUsername = errors.New("Cannot set empty username")
	ErrUserPasswordNotSet     = errors.New("User password not set")
	ErrEmailTaken             = errors.New("Email taken")
	ErrMissingPassword        = errors.New("Missing password")
)

type OauthUser struct {
	ID       int
	Email    string
	Password string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (s oauthService) UserExists(email string) bool {
	user, err := s.FindUserByEmail(email)
	return err == nil && user.Email != ""
}

// TODO: move logic with database to repository
func (s oauthService) FindUserByEmail(email string) (*OauthUser, error) {
	user := new(OauthUser)
	result := s.db.Table("db_local.users").Where("email = LOWER(?)", email).Find(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (s oauthService) CreateUser(email, password string) (*OauthUser, error) {
	user := &OauthUser{
		ID:        rand.Intn(100),
		Email:     strings.ToLower(email),
		CreatedAt: time.Now().UTC(),
	}

	if password == "" {
		return nil, ErrMissingPassword
	}

	if password != "" {
		passwordHash, err := utils.HashPassword(password)
		if err != nil {
			return nil, err
		}
		user.Password = string(passwordHash)
	}

	if s.UserExists(user.Email) {
		return nil, ErrEmailTaken
	}

	if err := s.db.Table("db_local.users").Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s oauthService) GetAuthUser(email, password string) (*OauthUser, error) {
	user, err := s.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	// Verify the password
	if utils.VerifyPassword(user.Password, password) != nil {
		return nil, ErrInvalidUserPassword
	}

	return user, nil
}
