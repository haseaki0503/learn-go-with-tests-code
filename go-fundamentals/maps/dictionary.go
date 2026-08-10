package main

import (
	"errors"
)

const (
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(w string) (string, error) {

	difinition, ok := d[w]

	if !ok {
		return "", ErrNotFound
	}

	return difinition, nil
}

func (d Dictionary) Add(word, difinition string) error {

	_, err := d.Search(word)

	switch err {
		case ErrNotFound:
			d[word] = difinition
		case nil:
			return ErrWordExists
		default:
			return err
	}

	return nil
}
