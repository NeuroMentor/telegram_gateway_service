package srverr

type serverError struct {
	error   string
	message string
	details string

	servError Error
}

func NewServerError(error string, servError Error) ServerError {
	return &serverError{
		error:     error,
		servError: servError,
	}
}

func (s *serverError) Error() string {
	return s.error
}

func (s *serverError) SetMessage(msg string) ServerError {
	s.message = msg
	return s
}

func (s *serverError) SetDetails(details string) ServerError {
	s.details = details
	return s
}

func (s *serverError) GetServerError() Error {
	return s.servError
}

func (s *serverError) GetMessage() string {
	return s.message
}

func (s *serverError) GetDetails() string {
	return s.details
}
