package dkeepr

// Orm represents an ORM with common CRUD operations
//
// TODO: will it be useful? perhaps its not necessary as there is
// no standard like Java JPA
type Orm interface {
	Begin() (Transaction, error)
	Save(o interface{}) (interface{}, error)
	Delete(o interface{}) error
	Find(id interface{}) (interface{}, error)
	FindAll(field string, value interface{}) (interface{}, error)
}

// Transaction is the representation of a transaction
type Transaction interface {
	Begin() (Transaction, error)
	Add(fn func() error) Transaction
	CommitOrRollback() error
}

// Save saves an entity
func (d *Dkeepr) Save(o interface{}) (interface{}, error) {

	pe, err := parseEntity(o)
	if err != nil {
		return nil, err
	}

	id, err := d.driver.Save(pe)
	if err != nil {
		return nil, err
	}

	return id, nil
}

// Delete deletes an entity given its ID
func (d *Dkeepr) Delete(o interface{}) error {
	pe, err := parseEntity(o)
	if err != nil {
		return err
	}

	return d.driver.Delete(pe)
}

// Find finds an entity given its ID
func (d *Dkeepr) Find(id interface{}) (interface{}, error) {
	return nil, nil
}

// FindAll finds one or more entities given a field and its value
func (d *Dkeepr) FindAll(field string, value interface{}) (interface{}, error) {
	return nil, nil
}

// Begin starts a new transaction
func (d *Dkeepr) Begin() (Transaction, error) {
	return d.Begin()
}
