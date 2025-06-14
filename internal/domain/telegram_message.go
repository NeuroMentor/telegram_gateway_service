package domain

type TelegramMessage struct {
	MessageID int64
	From      *TelegramUser
	Chat      *TelegramChat
	Date      int64
	Text      string
}
