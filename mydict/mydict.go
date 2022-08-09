package mydict

import "errors"

// Dicotinary type
type Dictionary map[string]string

var (
	errNotFund    = errors.New("not found")
	errWordExists = errors.New("invalidate data")
	errCantUpdate = errors.New("can't update non-existing word")
	errCantDelete = errors.New("can't delete non-existing word")
)

//  Find a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFund
}

// Add a word
func (d Dictionary) Add(word string, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFund:
		d[word] = def
	case nil:
		return errWordExists
	}
	// if err == errNotFund {
	// 	d[word] = def
	// } else if err == nil {
	// 	return errWordExists
	// }

	return nil
}

// Update word
func (d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFund:
		return errCantUpdate
	}

	return nil

}

// Delete
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case errNotFund:
		return errCantDelete
	}
	return nil
}
