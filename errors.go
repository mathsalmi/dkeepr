package dkeepr

import "errors"

// General ORM errors
var (
	// driver
	ErrDriverNotSupported = errors.New("ORM: driver not supported")
	ErrDriverNotChosen    = errors.New("ORM: driver not chosen yet. Use NewDkeepr func to get a proper instance")

	// connection
	ErrOpenConnection = errors.New("ORM: error opening connection with DB")
	ErrConnNotOpen    = errors.New("ORM: connection not open")

	// reflection
	ErrNotStruct   = errors.New("ORM: only structs can be persisted")
	ErrUnknownType = errors.New("ORM: unknown type")
)
