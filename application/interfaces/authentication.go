package interfaces

import "context"

type Token struct {
	Token   string
	UUID    string
	Expires int64
}

type Tokens struct {
	AccessToken  *Token
	RefreshToken *Token
}

type Claims struct {
	AccessUUID  string
	RefreshUUID string
	UserID      uint64
}

type Authentication interface {
	CreateAuth(ctx context.Context, userid uint64) (*Tokens, error)
	RefreshAuth(ctx context.Context, accessToken string, refreshToken string) (*Tokens, error)
	FetchAuth(ctx context.Context, token string) (uint64, error)
	DeleteAuth(ctx context.Context, token string) error
}
