package constants

// FIX: Define a custom key type for context
type ContextKey string

// FIX: Define the key to be used for storing user data
const (
	UserLoginKey ContextKey = "userLogin"
	Token        = "token"
)
