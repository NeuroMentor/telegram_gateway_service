package webhook

import (
	"context"
	"telegram_gateway_service/internal/domain"
	srverr "telegram_gateway_service/internal/shared/server_error"
)

type webhookService struct{}

func NewWebhookService() Service {
	return &webhookService{}
}

func (m *webhookService) ProcessUpdate(ctx context.Context, update *domain.TelegramUpdate) srverr.ServerError {
	//TODO implement me
	panic("implement me")
}

func (m *webhookService) ProcessMessage(ctx context.Context, message *domain.TelegramMessage) srverr.ServerError {
	//TODO implement me
	panic("implement me")
}

func (m *webhookService) ProcessCallback(ctx context.Context, callback *domain.TelegramCallbackQuery) srverr.ServerError {
	//TODO implement me
	panic("implement me")
}

func (m *webhookService) ProcessEditedMessage(ctx context.Context, message *domain.TelegramMessage) srverr.ServerError {
	//TODO implement me
	panic("implement me")
}
