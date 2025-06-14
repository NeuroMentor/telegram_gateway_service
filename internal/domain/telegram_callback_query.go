package domain

type TelegramCallbackQuery struct {
	ID      string
	From    *TelegramUser
	Data    string
	Message *TelegramMessage
}
