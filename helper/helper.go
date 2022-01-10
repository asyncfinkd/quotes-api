package helper

import (
	"errors"
)

func CreateError(submission string) error {
	return errors.New(submission)
}
