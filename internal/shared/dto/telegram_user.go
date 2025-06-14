package dto

import (
	"telegram_gateway_service/internal/domain"
	"telegram_gateway_service/models"
)

func TelegramUserFromModels(user *models.TelegramUser) *domain.TelegramUser {
	if user == nil {
		return nil
	}

	return &domain.TelegramUser{
		ID:           user.ID,
		IsBot:        user.IsBot,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Username:     user.Username,
		LanguageCode: user.LanguageCode,
	}
}

func TelegramUserToModels(user *domain.TelegramUser) *models.TelegramUser {
	if user == nil {
		return nil
	}

	return &models.TelegramUser{
		ID:           user.ID,
		IsBot:        user.IsBot,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Username:     user.Username,
		LanguageCode: user.LanguageCode,
	}
}
