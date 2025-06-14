package domain

type TelegramUpdate struct {
	UpdateID      int64
	Message       *TelegramMessage
	CallbackQuery *TelegramCallbackQuery
}
