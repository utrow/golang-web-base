package out

type ErrorResponse struct {
	Message string `json:"message"`
}

var (
	ErrInternal = "Internal Error."
	ErrRequire  = "%s is required."
)
