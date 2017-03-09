package handler

import (
	"encoding/json"
	"net/http"

	ae "github.com/mantishK/galore/apperror"
)

type ok interface {
	OK() *ae.Error
}

func decode(r *http.Request, v interface{}) *ae.Error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return ae.JsonDecode("", err)
	}
	if validatable, ok := v.(ok); ok {
		return validatable.OK()
	}
	return nil
}
