package constant

const (
	SuccessMessage = "Operation successful"
	ErrorMessage   = "Something went wrong"

	// Role constants
	RoleAdmin = "admin"
	RoleUser  = "user"
)

// Standard Error Messages
var (
	ErrNotFound      = "Record not found"
	ErrUnauthorized  = "Unauthorized access"
	ErrInternalError = "Internal server error"
)
