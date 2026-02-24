package godserr

import "errors"

var (
	ErrNoItems      error = errors.New("no items found")
	ErrKeyNotFound  error = errors.New("key not found")
	ErrItemNotFound error = errors.New("item not found")
)
