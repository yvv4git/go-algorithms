package extract_ints_numbers_from_str

import "errors"

var (
	ErrEmptyQuery        = errors.New("empty query")
	ErrInvalidPortNumber = errors.New("invalid port number")
	ErrPortsNoFound      = errors.New("ports no found")
)
