package mydict

import "errors"

// Dicotinary type
type Dictionary map[string]string

var errNotFund = errors.New("not found")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFund
}
