package reply

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"

	ae "github.com/mantishK/galore/apperror"
	"github.com/mantishK/galore/config"
	"github.com/mantishK/galore/log"
)

type public interface {
	Public() interface{}
}

func OK(w http.ResponseWriter, data interface{}) {
	if obj, ok := data.(public); ok {
		data = obj.Public()
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		Err(w, ae.JsonEncode(""))
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonData))
}

func Err(w http.ResponseWriter, e *ae.Error) {
	log.Err(*e)
	log.Err(string(debug.Stack()))
	// Unset error logging if prod
	if env, ok := config.GetString("app_env"); ok && env != "prod" {
		e.Log = ""
	}
	jsonData, err := json.Marshal(e)
	if err != nil {
		fmt.Fprint(w, "Error")
	}
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, string(jsonData), e.HttpStatus)
}
