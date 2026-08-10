package main

import (
	"testing"
)

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		difinition := "this is just a test"
		err := dictionary.Add(word, difinition)

		assertError(t, err, nil)
		assertDifinition(t, dictionary, word, difinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		difinition := "this is just a test"
		dictionary := Dictionary{ word: difinition }
		err := dictionary.Add(word, difinition)

		assertError(t, err, ErrWordExists)
		assertDifinition(t, dictionary, word, difinition)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expected to get an error.")
	}
}

func assertDifinition(t *testing.T, dictionary Dictionary, word, difinition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", word)
	}
	if difinition != got {
		t.Errorf("got %q want %q", got, difinition)
	}
}
