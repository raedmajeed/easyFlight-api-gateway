package utitlity

import "fmt"

type Error struct {
	Code		int
	Message string
}

func (e Error) ReturnError() (int, string) {
	return e.Code, e.Message
}

func (e Error) Error() (string) {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}