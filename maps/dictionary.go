package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("Key is not found in the Dictionary.")
var ErrExists = errors.New("Key exists")

func (d Dictionary) Search(word string) (string, error) {
	defination, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return defination, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, ok := d[word]
	if ok {
		return ErrExists
	}
	d[word] = definition
	return nil
}
