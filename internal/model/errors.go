package model

import "errors"

var ErrNotFound = errors.New("resource not fount")

var ErrInvalidInput = errors.New("invalid output")

var ErrCiruitOpen = errors.New("circuit breaker is open")

var ErrCacheUnavailable = errors.New("cache unavailable")

var ErrSyncInProgress = errors.New("sync already in progress")

var ErrSFTPConnection = errors.New("sftp connection failed")

var ErrInvalidSwiftCode = errors.New("invalid swift code format")

var ErrRateLimitExceeded = errors.New("rate limit exceeded")

var ErrUnauthorized = errors.New("unauthorized")
