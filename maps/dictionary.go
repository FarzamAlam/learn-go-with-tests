package maps

type Dictionary map[string]string

const (
	ErrNotFound      = DictionaryErr("Key is not found in the Dictionary.")
	ErrExists        = DictionaryErr("Key exists.")
	ErrDoesNotExists = DictionaryErr("Key does not exists.")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

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

func (d Dictionary) Update(word, defination string) error {
	_, ok := d[word]
	if !ok {
		return ErrDoesNotExists
	}
	d[word] = defination
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
