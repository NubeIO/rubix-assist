package tasks

import "errors"

type TransactionType int

// go generate ./...

//go:generate stringer -type=TransactionType
const (
	Fail TransactionType = iota
	Completed
)

func CheckTransaction(s string) error {
	switch s {
	case Fail.String():
		return nil
	case Completed.String():
		return nil
	}
	return errors.New("invalid action type, try Fail or Completed")
}
