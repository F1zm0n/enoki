package apperror

import "errors"

var (
	ErrNoSuchPkg      = errors.New("no such package in repository")
	ErrGettingPkgInfo = errors.New("error getting package info")
	ErrReadingConfig  = errors.New(
		"error reading configuration file ~/.config/enoki/enoki.conf, check configuration file",
	)
)
