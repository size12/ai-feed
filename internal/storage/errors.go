package storage

import "errors"

var (
	ErrNotCreated    = errors.New("not created")
	ErrNotFound      = errors.New("not found")
	ErrFailedUpdate  = errors.New("failed update")
	ErrAlreadyExists = errors.New("already exists")
)
