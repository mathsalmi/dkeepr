package dkeepr

import "reflect"

// returns a value of a reflected object
func getReflectedValue(t reflect.Kind, val reflect.Value) (value interface{}, err error) {
	switch t {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value = val.Int()
	case reflect.Float32, reflect.Float64:
		value = val.Float()
	case reflect.String:
		value = val.String()
	// TODO: how to handle a struct?
	default:
		err = ErrUnknownType
	}

	return
}
