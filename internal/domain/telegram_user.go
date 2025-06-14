package domain

type TelegramUser struct {
	ID           int64
	IsBot        bool
	FirstName    string
	LastName     string
	Username     string
	LanguageCode string
}
