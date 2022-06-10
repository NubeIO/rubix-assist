package tasks

import "errors"

type TransactionType int

// go generate ./...

//go:generate stringer -type=TransactionType
const (
	Failed TransactionType = iota
	Completed
)

func CheckTransaction(s string) error {
	switch s {
	case Failed.String():
		return nil
	case Completed.String():
		return nil
	}
	return errors.New("invalid action type, try Fail or Completed")
}
