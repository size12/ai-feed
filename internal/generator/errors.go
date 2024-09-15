package generator

import "errors"

var (
	ErrNoChoices = errors.New("no available choices, can't generate article")
)
