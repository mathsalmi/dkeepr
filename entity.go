package dkeepr

import (
	"reflect"
	"strings"
)

// ParseEntity is the result of a parsed entity (struct)
// It groups the table name, column names and values of that object
type parsedEntity struct {
	tablename string

	pk       []string
	pkvalues []interface{}

	columns []string
	values  []interface{}
}

func (pe *parsedEntity) Tablename() string {
	return pe.tablename
}

func (pe *parsedEntity) Pk() []string {
	return pe.pk
}

func (pe *parsedEntity) Pkvalues() []interface{} {
	return pe.pkvalues
}

func (pe *parsedEntity) Columns() []string {
	return pe.columns
}

func (pe *parsedEntity) Values() []interface{} {
	return pe.values
}

// Parses the entity
//
// It checks for the table name, primary key names, column names and convert values
//
// Struct can contain the following annotations
// - orm_ignore: ORM will not map this Field
// - orm_pk: sets the field as a primary key
// - orm_field_name: sets custom name for the field
// - orm_insert: defines if column should be present on insert sqls or not
func parseEntity(o interface{}) (*parsedEntity, error) {
	t := reflect.TypeOf(o)

	// TODO: check if pointer
	// if t.Elem().Kind() != reflect.Struct {
	// 	return nil, ErrNotStruct
	// }

	pe := new(parsedEntity)

	// table name
	pe.tablename = t.Name()

	// fields
	v := reflect.ValueOf(o)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// ignore
		if field.Tag.Get("orm_ignore") == "true" {
			continue
		}

		// name
		fieldName := field.Tag.Get("orm_field_name")
		if strings.Trim(fieldName, " ") == "" {
			fieldName = field.Name
		}

		// value
		fieldRef := v.FieldByName(field.Name)
		fieldValue, err := getReflectedValue(fieldRef.Kind(), fieldRef)
		if err != nil {
			return nil, err
		}

		// pk
		if field.Tag.Get("orm_pk") == "true" {
			pe.pk = append(pe.pk, fieldName)
			pe.pkvalues = append(pe.pkvalues, fieldValue)
		}

		// columns
		if field.Tag.Get("orm_insert") != "false" {
			pe.columns = append(pe.columns, fieldName)
			pe.values = append(pe.values, fieldValue)
		}

	}

	return pe, nil
}
