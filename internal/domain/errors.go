package domain

import "errors"

// собрать написанные ошибки в этом файле и добавить новую ошибку ErrRefreshTokenExpired

var (
	ErrUserNotFound = errors.New("User not found")
	ErrBookNotFound = errors.New("Book not found")
	ErrRefreshTokenExpired = errors.New("The refresh token has expired")
)
