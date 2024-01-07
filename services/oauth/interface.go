package oauth

import "gorm.io/gorm"

func NewOauthService(db *gorm.DB) OauthService {
	return &oauthService{
		db: db,
	}
}

type OauthService interface {
	UserExists(email string) bool
	CreateUser(email, password string) (*OauthUser, error)
	GetAuthUser(email, password string) (*OauthUser, error)
	GrantAuthorizationCode(*OauthUser) (string, error)
}

type oauthService struct {
	db *gorm.DB
}
