package jwt_test

import (
	"context"
	"testing"
	"time"

	"github.com/kustavo/benchmark/go/domain/message"
	"github.com/kustavo/benchmark/go/infrastructure/jwt"
)

type cache struct {
	inMemory map[string]string
}

func (c cache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	c.inMemory[key] = value.(string)
	return nil
}
func (c cache) Get(ctx context.Context, key string) (string, error) {
	value, found := c.inMemory[key]
	if !found {
		return "", message.ErrItemNotFound
	}
	return value, nil
}
func (c cache) Del(ctx context.Context, key string) error {
	delete(c.inMemory, key)
	return nil
}
func (c cache) Close() {

}

func TestCreateAuth(t *testing.T) {

	cache := cache{}
	jwt := jwt.NewJwt(cache)

	tokens, err := jwt.CreateAuth(context.Background(), 555)
	if err != nil {
		t.Error(err)
	}

	if tokens.AccessToken == nil {
		t.Error("Access token is nil")
	}

	if tokens.RefreshToken == nil {
		t.Error("Refresh token is nil")
	}
}

func TestFetchAuth(t *testing.T) {

	// cache := cache{}
	// jwt := jwt.NewJwt(cache)

	// tokens, err := jwt.CreateAuth(context.Background(), 555)

	// userID, err := jwt.FetchAuth(context.Background(), tokens.AccessToken.Token)

	// tokens, err := jwt.CreateAuth(context.Background(), 0)
	// if err != nil {
	// 	t.Error(err)
	// }

	// if tokens.AccessToken == nil {
	// 	t.Error("Access token is nil")
	// }

	// if tokens.RefreshToken == nil {
	// 	t.Error("Refresh token is nil")
	// }
}
