package drivers

// IsDriverSupported tells whether or not the ORM supports a given driver
func IsDriverSupported(driver string) bool {
	drivers := []string{"postgres", "mssql"}

	for _, item := range drivers {
		if driver == item {
			return true
		}
	}

	return false
}

// TODO: define an unexported interface for drivers, to be called "ormDrivers"
type ormDriver interface {
	save(table string, columns []string, values []interface{}) (interface{}, error)
}
