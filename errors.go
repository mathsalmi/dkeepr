package dkeepr

import "errors"

// General ORM errors
var (
	ErrDriverNotSupported = errors.New("ORM: driver not supported")
	ErrOpenConnection     = errors.New("ORM: error opening connection with DB")
	ErrNotStruct          = errors.New("ORM: only structs can be persisted")
)
