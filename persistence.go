package dkeepr

import "bitbucket.org/mathsalmi/dkeepr/drivers"

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
func (d *Dkeepr) Begin() (drivers.Transaction, error) {
	return d.driver.Begin()
}
