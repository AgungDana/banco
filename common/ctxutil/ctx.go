package ctxutil

import (
	"context"
	"encoding/json"
	"errors"

	"banco/common/jwt"

	"github.com/google/uuid"
)

type key string

const (
	USER_PAYLOAD   key = "X-Banco-PAYLOAD"
	USER_ID        key = "X-Banco-ID"
	TRANSACTION_ID key = "X-Banco-Txn-ID"
)

func getFromContext(ctx context.Context, key key) (any, bool) {
	data := ctx.Value(key)
	switch u := data.(type) {
	case string:
		v, err := uuid.Parse(u)
		if err != nil {
			return nil, false
		}
		return v, true
	case uint:
		return u, true
	case uuid.UUID:
		return u.String(), true
	case jwt.Payload:
		data, err := json.Marshal(u)
		if err != nil {
			return "", false
		}
		return string(data), true
	case *jwt.Payload:
		data, err := json.Marshal(u)
		if err != nil {
			return "", false
		}
		return string(data), true
	default:
		return "", false
	}
}

func GetUserPayloadFromCtx(ctx context.Context) (*jwt.Payload, error) {
	data, ok := getFromContext(ctx, USER_PAYLOAD)
	if !ok {
		return nil, errors.New("internal server error")
	}
	payload := jwt.Payload{}
	err := json.Unmarshal([]byte(data.(string)), &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func GetUserIdFromCtx(ctx context.Context) (*uint, error) {
	data, ok := getFromContext(ctx, USER_ID)
	if !ok {
		return nil, errors.New("internal server error")
	}
	id := data.(uint)
	return &id, nil
}

func GetTransactionId(ctx context.Context) (uuid.UUID, bool) {
	id, ok := getFromContext(ctx, TRANSACTION_ID)
	idUuid, _ := uuid.Parse(id.(string))
	return idUuid, ok
}

func NewRequestTransactionId(ctx context.Context) context.Context {
	return context.WithValue(ctx, TRANSACTION_ID, uuid.New())
}

func NewRequestPayload(ctx context.Context, payload jwt.Payload) context.Context {
	return context.WithValue(ctx, USER_PAYLOAD, payload)
}
func NewRequestUserID(ctx context.Context, userID uint) context.Context {
	return context.WithValue(ctx, USER_ID, userID)
}
