package errors

type ErrorCode string

const (
    ErrUnauthorized        ErrorCode = "UNAUTHORIZED"
    ErrTokenExpired        ErrorCode = "TOKEN_EXPIRED"
    ErrTokenInvalid        ErrorCode = "TOKEN_INVALID"

    ErrNotFound            ErrorCode = "NOT_FOUND"
    ErrRecipeAlreadyExists ErrorCode = "RECIPE_ALREADY_EXISTS"

    ErrValidation          ErrorCode = "VALIDATION_ERROR"

    ErrInternal            ErrorCode = "INTERNAL_ERROR"
)