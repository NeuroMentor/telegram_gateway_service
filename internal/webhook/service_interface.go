package webhook

import (
	"context"
	"telegram_gateway_service/internal/domain"
	srverr "telegram_gateway_service/internal/shared/server_error"
)

type Service interface {
	ProcessUpdate(ctx context.Context, update *domain.TelegramUpdate) srverr.ServerError
	ProcessMessage(ctx context.Context, message *domain.TelegramMessage) srverr.ServerError
	ProcessCallback(ctx context.Context, callback *domain.TelegramCallbackQuery) srverr.ServerError
	ProcessEditedMessage(ctx context.Context, message *domain.TelegramMessage) srverr.ServerError
}
