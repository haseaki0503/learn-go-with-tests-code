package main

import (
	"errors"
)

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (d Dictionary) Search(w string) (string, error) {

	difinition, ok := d[w]

	if !ok {
		return "", ErrNotFound
	}

	return difinition, nil
}

func (d Dictionary) Add(word, difinition string) {
	d[word] = difinition
}
