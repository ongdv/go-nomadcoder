package mydict

import "errors"

// Dicotinary type
type Dictionary map[string]string

var errNotFund = errors.New("not found")
var errWordExists = errors.New("invalidate data")

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
