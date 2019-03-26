package main

import "fmt"

type InvalidStatusError struct {
	status string
}

func (e InvalidStatusError) Error() string {
	return fmt.Sprintf("invalid status %q", e.status)
}

func NewInvalidStatusError(status string) InvalidStatusError {
	return InvalidStatusError {
		status: status,
	}
}
