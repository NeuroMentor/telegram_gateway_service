package dto

import (
	"telegram_gateway_service/internal/domain"
	"telegram_gateway_service/models"
)

func TelegramMessageFromModels(message *models.TelegramMessage) *domain.TelegramMessage {
	if message == nil {
		return nil
	}

	return &domain.TelegramMessage{
		MessageID: message.MessageID,
		From:      TelegramUserFromModels(message.From),
		Date:      message.Date,
		Text:      message.Text,
		Chat:      TelegramChatFromModels(message.Chat),
	}
}

func TelegramMessageToModels(message *domain.TelegramMessage) *models.TelegramMessage {
	if message == nil {
		return nil
	}

	return &models.TelegramMessage{
		MessageID: message.MessageID,
		From:      TelegramUserToModels(message.From),
		Date:      message.Date,
		Text:      message.Text,
		Chat:      TelegramChatToModels(message.Chat),
	}
}
