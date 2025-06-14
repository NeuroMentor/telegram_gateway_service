package dto

import (
	"telegram_gateway_service/internal/domain"
	"telegram_gateway_service/models"
)

func TelegramChatFromModels(chat *models.TelegramChat) *domain.TelegramChat {
	if chat == nil {
		return nil
	}

	return &domain.TelegramChat{
		ID:        chat.ID,
		Type:      chat.Type,
		Title:     chat.Title,
		Username:  chat.Username,
		FirstName: chat.FirstName,
		LastName:  chat.LastName,
	}
}

func TelegramChatToModels(chat *domain.TelegramChat) *models.TelegramChat {
	if chat == nil {
		return nil
	}

	return &models.TelegramChat{
		ID:        chat.ID,
		Type:      chat.Type,
		Title:     chat.Title,
		Username:  chat.Username,
		FirstName: chat.FirstName,
		LastName:  chat.LastName,
	}
}
