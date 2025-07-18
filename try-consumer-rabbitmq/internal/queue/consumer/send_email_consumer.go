package consumer

import (
	"context"

	"github.com/rahmatrdn/go-skeleton/internal/helper"
	"github.com/rahmatrdn/go-skeleton/internal/usecase/todo_list/entity"
)

type SendEmailConsumer struct {
	ctx context.Context
}

type ISendEmailConsumer interface {
	Process(payload map[string]interface{}) error
}

func NewSendEmailConsumer(
	ctx context.Context,
) *SendEmailConsumer {
	return &SendEmailConsumer{ctx}
}

func (l *SendEmailConsumer) Process(payload map[string]interface{}) error {
	var params entity.SendEmailReq
	params.LoadFromMap(payload)

	helper.Dump(params)

	return nil
}
