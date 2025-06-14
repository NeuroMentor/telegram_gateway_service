package dto

import (
	"telegram_gateway_service/internal/domain"
	"telegram_gateway_service/models"
)

func TelegramCallbackQueryFromModels(callbackQuery *models.TelegramCallbackQuery) *domain.TelegramCallbackQuery {
	if callbackQuery == nil {
		return nil
	}

	return &domain.TelegramCallbackQuery{
		ID:      callbackQuery.ID,
		From:    TelegramUserFromModels(callbackQuery.From),
		Data:    callbackQuery.Data,
		Message: TelegramMessageFromModels(callbackQuery.Message),
	}
}

func TelegramCallbackQueryToModels(callbackQuery *domain.TelegramCallbackQuery) *models.TelegramCallbackQuery {
	if callbackQuery == nil {
		return nil
	}

	return &models.TelegramCallbackQuery{
		ID:      callbackQuery.ID,
		From:    TelegramUserToModels(callbackQuery.From),
		Data:    callbackQuery.Data,
		Message: TelegramMessageToModels(callbackQuery.Message),
	}
}
