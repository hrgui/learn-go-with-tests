package main

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (dict Dictionary) Search(key string) (string, error) {
	def, ok := dict[key]

	if !ok {
		return "", ErrNotFound
	}

	return def, nil
}

func (dict Dictionary) Add(key string, value string) error {
	_, err := dict.Search(key)

	switch err {
	case ErrNotFound:
		dict[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (dict Dictionary) Update(key string, value string) error {
	_, err := dict.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		dict[key] = value
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
