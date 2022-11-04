package cmd

import (
	"encoding/json"
	"fmt"
)

type invalidActionError struct {
	action interface{}
}

func newInvalidActionError(action interface{}) invalidActionError {
	return invalidActionError{action}
}

func (err invalidActionError) Error() string {
	j, _ := json.Marshal(err.action)
	return fmt.Sprintf("invalid action: %s", string(j))
}
