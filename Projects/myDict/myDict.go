package myDict

import (
	"errors"
	"fmt"
)

// Type alias
type Money int

// var money Money = 2

// just 'alias(가명)', structure 아님!! - key, value가 string
type Dictionary map[string]string

var errNotFound = errors.New("can not found the word")
var errWordExist = errors.New("that word already exists")
var errCantUpdate = errors.New("can't update non-existing word")
var errCantDelete = errors.New("can't delete non-existing word")

//# 모든 type은 receiver를 이용하여 methods를 가질 수 있다!!

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	//# Map은 해당 key에 value가 존재하는지의 여부를 반환해준다!!!
	value, exists := d[word]

	if exists {
		return value, nil
	}

	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word string, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		fmt.Println("exist")
		return errWordExist
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		// map docs 참고
		delete(d, word)
	case errNotFound:
		return errCantDelete
	}

	return nil
}
