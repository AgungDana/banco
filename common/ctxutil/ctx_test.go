package ctxutil_test

import (
	"banco/common/ctxutil"
	"banco/common/jwt"
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestCtxutil(t *testing.T) {
	ctx := context.Background()
	payload := jwt.Payload{
		UserID:    1,
		Email:     "email@gmail.com",
		IssuedAt:  time.Now().Add(24 * time.Hour),
		ExpiredAt: time.Now().Add(48 * time.Hour),
	}
	log.Println("set payload")
	ctx = ctxutil.NewRequestPayload(ctx, payload)
	log.Println("set id")
	ctx = ctxutil.NewRequestUserID(ctx, payload.UserID)
	log.Println("set transaction id")
	ctx = ctxutil.NewRequestTransactionId(ctx)

	log.Println("get transaction id")
	fmt.Println(ctxutil.GetTransactionId(ctx))
	log.Println("get user id")
	fmt.Println(ctxutil.GetUserIdFromCtx(ctx))
	log.Println("get payload id")
	fmt.Println(ctxutil.GetUserPayloadFromCtx(ctx))
}

func TestCtxutil2(t *testing.T) {
	ctx := &gin.Context{}
	ctx1 := ctxutil.NewRequestTransactionId(ctx.Request.Context())
	// data, ok := ctx1.Value(ctxutil.TRANSACTION_ID).(uuid.UUID)
	// if !ok {
	// 	log.Fatalf("error")
	// }
	// ctx.Set(string(ctxutil.TRANSACTION_ID), data)
	ctx.Request = ctx.Request.WithContext(ctx1)

	id, ok := ctxutil.GetTransactionId(ctx.Request.Context())
	if !ok {
		log.Fatalf("error")
	}

	fmt.Println(id)
}
