package domain

import "encoding/json"

type Error struct {
	NameFile string
	Err      error
	Msg      string
}

func (e *Error) Error() string {
	jbyte, err := json.Marshal(e)
	if err != nil {
		return "Failed to marshal error"
	}
	return string(jbyte)
}
