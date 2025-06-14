package srverr

type ServerError interface {
	Error() string

	SetMessage(msg string) ServerError
	SetDetails(details string) ServerError

	GetServerError() Error

	GetMessage() string
	GetDetails() string
}
