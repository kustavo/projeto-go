// jwt handles the JWT token
package jwt

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	gjwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/kustavo/benchmark/go/application/interfaces"
	"github.com/kustavo/benchmark/go/domain/message"
)

type tokenType int

const (
	AccessToken tokenType = iota
	RefreshToken
)

type jwt struct {
	cache interfaces.Cache
}

// NewJwt instantiates a new jwt token
func NewJwt(cache interfaces.Cache) *jwt {
	return &jwt{cache: cache}
}

// CreateAuth creates a new auth token and refresh token and stores them in the cache
func (j *jwt) CreateAuth(ctx context.Context, userid uint64) (*interfaces.Tokens, error) {
	now := time.Now()
	refreshUUID := uuid.New().String()

	accessToken, err := createAccessToken(userid, refreshUUID)
	if err != nil {
		return nil, err
	}

	attime := time.Unix(accessToken.Expires, 0)
	err = j.cache.Set(ctx, accessToken.UUID, strconv.Itoa(int(userid)), attime.Sub(now))
	if err != nil {
		return nil, err
	}

	refreshToken, err := createRefreshToken(userid, refreshUUID)
	if err != nil {
		return nil, err
	}

	rttime := time.Unix(refreshToken.Expires, 0)
	err = j.cache.Set(ctx, refreshToken.UUID, strconv.Itoa(int(userid)), rttime.Sub(now))
	if err != nil {
		return nil, err
	}

	return &interfaces.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// createAccessToken creates a new access token
func createAccessToken(userID uint64, refreshUUID string) (*interfaces.Token, error) {
	td := &interfaces.Token{}
	td.Expires = time.Now().Add(time.Minute * 5).Unix()
	td.UUID = uuid.New().String()

	atClaims := gjwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.UUID
	atClaims["refresh_uuid"] = refreshUUID
	atClaims["user_id"] = userID
	atClaims["exp"] = td.Expires
	at := gjwt.NewWithClaims(gjwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(os.Getenv("JWT_ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	td.Token = token
	return td, nil
}

// createRefreshToken creates a new refresh token
func createRefreshToken(userID uint64, UUID string) (*interfaces.Token, error) {
	td := &interfaces.Token{}
	td.Expires = time.Now().Add(time.Hour * 4).Unix()
	td.UUID = UUID

	rtClaims := gjwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.UUID
	rtClaims["user_id"] = userID
	rtClaims["exp"] = td.Expires
	rt := gjwt.NewWithClaims(gjwt.SigningMethodHS256, rtClaims)

	token, err := rt.SignedString([]byte(os.Getenv("JWT_REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}

	td.Token = token
	return td, nil
}

// verifyToken verifies the token and returns the gjwt.Token
func verifyToken(token string, tokenType tokenType, expirationValidation bool) (*gjwt.Token, error) {
	var secret string
	if tokenType == AccessToken {
		secret = os.Getenv("JWT_ACCESS_SECRET")
	} else if tokenType == RefreshToken {
		secret = os.Getenv("JWT_REFRESH_SECRET")
	}

	jwttoken, err := gjwt.Parse(token, func(token *gjwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*gjwt.SigningMethodHMAC); !ok {
			return nil, message.ErrUnexpectedSigningMethod
		}
		return []byte(secret), nil
	})

	if err != nil {
		if errors.Is(err, gjwt.ErrTokenExpired) {
			err = message.ErrExpiredToken
			if !expirationValidation {
				err = nil
			}
		} else if errors.Is(err, gjwt.ErrSignatureInvalid) {
			err = message.ErrInvalidTokenSignature
		}
	}

	if err != nil {
		return nil, err
	}

	return jwttoken, nil
}

// extractTokenMetadata extracts the metadata from the token and returns the claims
func extractTokenMetadata(token string, tokenType tokenType, expirationValidation bool) (*interfaces.Claims, error) {
	jwttoken, err := verifyToken(token, tokenType, expirationValidation)
	if err != nil {
		return nil, err
	}

	claims, ok := jwttoken.Claims.(gjwt.MapClaims)
	if !ok {
		return nil, message.ErrInvalidToken
	}

	var accessUuid string
	if tokenType == AccessToken {
		accessUuid = claims["access_uuid"].(string)
	}
	refreshUuid := claims["refresh_uuid"].(string)

	userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	if err != nil {
		return nil, err
	}

	return &interfaces.Claims{
		AccessUUID:  accessUuid,
		RefreshUUID: refreshUuid,
		UserID:      userId,
	}, nil
}

// FetchAuth fetches the auth token from the cache
func (j *jwt) FetchAuth(ctx context.Context, token string) (uint64, error) {
	claims, err := extractTokenMetadata(token, AccessToken, true)
	if err != nil {
		return 0, err
	}

	userid, err := j.cache.Get(ctx, claims.AccessUUID)
	if err != nil {
		if errors.Is(err, message.ErrItemNotFound) {
			return 0, message.ErrInvalidToken
		}
		return 0, err
	}

	userID, err := strconv.ParseUint(userid, 10, 64)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

// DeleteAuth deletes the auth and refresh tokens from the cache
func (j *jwt) DeleteAuth(ctx context.Context, token string) error {
	claims, err := extractTokenMetadata(token, AccessToken, false)
	if err != nil {
		return err
	}

	err = j.cache.Del(ctx, claims.AccessUUID)
	if err != nil {
		return err
	}

	err = j.cache.Del(ctx, claims.RefreshUUID)
	return err
}

// RefreshAuth deletes the old auth and refresh tokens from the cache, creates the new ones and saves them in the cache and returns them
func (j *jwt) RefreshAuth(ctx context.Context, accessToken string, refreshToken string) (*interfaces.Tokens, error) {
	claimsAccessToken, err := extractTokenMetadata(accessToken, AccessToken, false)
	if err != nil {
		return nil, err
	}

	claimsRefreshToken, err := extractTokenMetadata(refreshToken, RefreshToken, true)
	if err != nil {
		return nil, err
	}

	if claimsAccessToken.RefreshUUID != claimsRefreshToken.RefreshUUID {
		return nil, message.ErrInvalidRefreshToken
	}

	_, err = j.cache.Get(ctx, claimsRefreshToken.RefreshUUID)
	if err != nil {
		if errors.Is(err, message.ErrItemNotFound) {
			return nil, message.ErrInvalidToken
		}
		return nil, err
	}

	err = j.DeleteAuth(ctx, accessToken)
	if err != nil {
		return nil, err
	}
	return j.CreateAuth(ctx, claimsAccessToken.UserID)
}
