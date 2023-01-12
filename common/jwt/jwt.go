package jwt

import (
	"banco/common/werror"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	TOKEN_EXPIRED = werror.Error{
		Code:    "TOKENEXPIRED",
		Message: "Token Expired",
	}
)

type JwtTokenMaker interface {
	GeneratedToken(p Payload) (string, error)
	ValidateToken(token string) (*Payload, error)
}

func NewJwtToken(secretKey string) JwtTokenMaker {
	return &jwtTokenMaker{
		SecretKey: secretKey,
	}
}

type Payload struct {
	UserID    uint
	Email     string
	IssuedAt  time.Time
	ExpiredAt time.Time
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return TOKEN_EXPIRED
	}
	return nil
}

type jwtTokenMaker struct {
	SecretKey string
}

// GeneratedToken implements JwtTokenMaker
func (j *jwtTokenMaker) GeneratedToken(p Payload) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &p)
	return jwtToken.SignedString([]byte(j.SecretKey))
}

// ValidateToken implements JwtTokenMaker
func (j *jwtTokenMaker) ValidateToken(token string) (*Payload, error) {

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(j.SecretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		ver, ok := err.(jwt.ValidationError)
		if ok && errors.Is(ver.Inner, TOKEN_EXPIRED) {
			return nil, TOKEN_EXPIRED
		}
		return nil, err
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return payload, nil
}
