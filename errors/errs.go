package errs

import "errors"

// General ORM errors
var (
	// driver
	ErrDriverNotSupported = errors.New("ORM: driver not supported")
	ErrDriverNotChosen    = errors.New("ORM: driver not chosen yet. Use NewDkeepr func to get a proper instance")
	ErrMethodNotImpl      = errors.New("ORM: method not implemented")

	// connection
	ErrOpenConnection = errors.New("ORM: error opening connection with DB")
	ErrConnNotOpen    = errors.New("ORM: connection not open")

	// reflection
	ErrNotStruct   = errors.New("ORM: only structs can be persisted")
	ErrUnknownType = errors.New("ORM: unknown type") // TODO: show type name

	// data
	ErrNoResult = errors.New("ORM: no result")
	ErrNoPk     = errors.New("ORM: obligatory IDs not given")

	// transaction
	ErrBegin  = errors.New("ORM: error opening transaction")
	ErrCommit = errors.New("ORM: there was a problem with something. Rolling back transaction... ")
)
