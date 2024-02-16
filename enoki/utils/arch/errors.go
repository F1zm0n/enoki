package arch

import "errors"

var (
	ErrGettingPkgInfo = errors.New("error getting package info")
	ErrNoSuchPackage  = errors.New("no such package found")
	ErrSoPackage      = errors.New("this is not .so package")
)
