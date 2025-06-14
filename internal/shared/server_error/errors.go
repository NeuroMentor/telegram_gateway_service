package srverr

type ErrorTypeInternalServerError string

func (err ErrorTypeInternalServerError) String() string {
	return string(err)
}

const (
	ErrInternalServerError ErrorTypeInternalServerError = "internal server error"
)

type ErrorTypeBadRequest string

func (err ErrorTypeBadRequest) String() string {
	return string(err)
}

const (
	ErrBadRequest ErrorTypeBadRequest = "bad request"
)

type ErrorTypeNotFound string

func (err ErrorTypeNotFound) String() string {
	return string(err)
}

const (
	ErrNotExists ErrorTypeNotFound = "not exists"
)
