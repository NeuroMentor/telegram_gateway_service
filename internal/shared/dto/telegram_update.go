package dto

import (
	"telegram_gateway_service/internal/domain"
	"telegram_gateway_service/models"
)

func TelegramUpdateFromModels(update *models.TelegramUpdate) *domain.TelegramUpdate {
	if update == nil {
		return nil
	}

	return &domain.TelegramUpdate{
		UpdateID:      update.UpdateID,
		Message:       TelegramMessageFromModels(update.Message),
		CallbackQuery: TelegramCallbackQueryFromModels(update.CallbackQuery),
	}
}

func TelegramUpdateToModels(update *domain.TelegramUpdate) *models.TelegramUpdate {
	if update == nil {
		return nil
	}

	return &models.TelegramUpdate{
		UpdateID:      update.UpdateID,
		Message:       TelegramMessageToModels(update.Message),
		CallbackQuery: TelegramCallbackQueryToModels(update.CallbackQuery),
	}
}
