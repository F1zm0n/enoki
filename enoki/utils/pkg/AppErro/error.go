package apperror

import "errors"

var (
	ErrNoSuchPkg = errors.New("no such package in repository")
	BreakErr     = errors.New("break error")
)
