package errors

type ErrorCode string

const (
    // Auth
	ErrUnauthorized ErrorCode = "UNAUTHORIZED"
	ErrTokenExpired ErrorCode = "TOKEN_EXPIRED"
	ErrTokenInvalid ErrorCode = "TOKEN_INVALID"

    // Resource
    ErrBadRequest          ErrorCode = "BAD_REQUEST"
    ErrForbidden           ErrorCode = "FORBIDDEN"
	ErrNotFound            ErrorCode = "NOT_FOUND"
	ErrRecipeAlreadyExists ErrorCode = "RECIPE_ALREADY_EXISTS"

    // Validation
	ErrValidation ErrorCode = "VALIDATION_ERROR"

    // Conflict
    ErrConflict ErrorCode = "CONFLICT"

    // Internal
	ErrInternal    ErrorCode = "INTERNAL_ERROR"
	ErrUnavailable ErrorCode = "SERVICE_UNAVAILABLE"
)
