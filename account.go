package main

import "errors"

var ErrEmptyID = errors.New("ErrEmptyID")

type BankAccount struct {
	Name   string
	ID     int
	Salary int
	BankID int
}

func OK(b BankAccount) error {
	if b.ID == 0 {
		return ErrEmptyID
	}
	return nil
}
