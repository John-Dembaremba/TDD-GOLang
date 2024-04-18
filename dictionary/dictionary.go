package dictionary

import "errors"

type Dictionary map[string]string

var (
	errNotFound = errors.New("could not find the word you are looking for")
	ErrWordExists = errors.New("word already exists")
	ErrorWordDoesntExists = errors.New("cannot update word because it does not exist")
)


func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]

	if !ok {
		return "", errNotFound
	}
	return value, nil
}

func (d Dictionary)Add(word, definition string)error{
	_, err := d.Search(word)

	switch err{
	case errNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return nil
	}
	return nil
}

func (d Dictionary)Update(word, newDefinition string)error{
	_, err := d.Search(word)

	switch err{
	case errNotFound:
		return ErrorWordDoesntExists
	case nil:
		d[word]=newDefinition
	}
	return nil
}

func (d Dictionary)Delete(word string)error{
	_, err := d.Search(word)

	if err != nil{
		return ErrorWordDoesntExists
	}
	return nil
}


