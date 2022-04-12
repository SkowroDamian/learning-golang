package main

type Dictionary map[string]string

const (
	ErrNotFound     = DictionaryErr("could not find the word in dictionary")
	ErrWordExists   = DictionaryErr("key already exists")
	ErrWordNotFound = DictionaryErr("Word doesnt exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word] //mapa moze zwrci 2 wartosci, druga to bool wskazujcy czy udalo sie odnalezc wartosc
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key) //wbudowana funkcja usuwania dziaa na mapach
}
func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordNotFound
	case nil:
		d[key] = value
	default:
		return err
	}
	return nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
