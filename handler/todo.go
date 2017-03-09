package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	ae "github.com/mantishK/galore/apperror"
	"github.com/mantishK/galore/model"
	"github.com/mantishK/galore/reply"
)

func GetUserTodos(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)
	todo := model.Todo{}

	todo.UserID = userID
	todos, err := todo.GetUserTodos()
	if err != nil {
		reply.Err(w, ae.DB("", err))
		return
	}

	todosFormat := struct {
		Todos []model.Todo `json:"todos"`
	}{todos}

	reply.OK(w, todosFormat)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)
	todo := model.Todo{}

	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		reply.Err(w, ae.Required("id required", "id"))
		return
	}

	var err error
	todo.ID, err = strconv.Atoi(id)
	if err != nil {
		reply.Err(w, ae.NotNumericInput("", err, "id"))
		return
	}

	err = todo.Get()
	if err == sql.ErrNoRows {
		reply.Err(w, ae.ResourceNotFound(""))
		return
	}
	if err != nil {
		reply.Err(w, ae.DB("", err))
		return
	}
	if todo.UserID != userID {
		reply.Err(w, ae.Forbidden(""))
	}

	result := make(map[string]interface{})
	result["todo"] = todo

	reply.OK(w, todo)
}

// The request body for PostTodo handler that satisfies ok interface
type todoReqBody struct {
	Content string `json:"content"`
}

// The mentod that does intput validation for req body
func (rb *todoReqBody) OK() *ae.Error {
	if len(rb.Content) == 0 {
		return ae.Required("", "content")
	}
	return nil
}

func PostTodo(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)

	var rb todoReqBody

	aErr := decode(r, &rb)
	if aErr != nil {
		reply.Err(w, aErr)
	}

	todo := model.Todo{UserID: userID, Content: rb.Content}

	err := todo.Insert()
	if err != nil {
		reply.Err(w, ae.DB("", err))
		return
	}

	result := make(map[string]interface{})
	result["todo"] = todo

	reply.OK(w, result)
}

func PutTodo(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)

	var rb todoReqBody

	aErr := decode(r, &rb)
	if aErr != nil {
		reply.Err(w, aErr)
	}

	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		reply.Err(w, ae.Required("id required", "id"))
		return
	}

	todoID, err := strconv.Atoi(id)
	if err != nil {
		reply.Err(w, ae.NotNumericInput("", err, "id"))
		return
	}

	todo := model.Todo{ID: todoID}

	err = todo.Get()
	if err == sql.ErrNoRows {
		reply.Err(w, ae.ResourceNotFound(""))
		return
	}
	if err != nil {
		reply.Err(w, ae.DB("", err))
		return
	}
	if todo.UserID != userID {
		reply.Err(w, ae.Forbidden(""))
		return
	}
	todo.Content = rb.Content
	err = todo.Update()
	if err != nil {
		reply.Err(w, ae.DB("", err))
		return
	}

	result := make(map[string]interface{})
	result["todo"] = todo

	reply.OK(w, result)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		reply.Err(w, ae.Required("id required", "id"))
		return
	}

	todoID, err := strconv.Atoi(id)
	if err != nil {
		reply.Err(w, ae.NotNumericInput("", err, "id"))
		return
	}

	todo := model.Todo{ID: todoID}

	err = todo.Get()
	if err == sql.ErrNoRows {
		reply.Err(w, ae.ResourceNotFound(""))
		return
	}
	if err != nil {
		reply.Err(w, ae.DB("", err))
		return
	}
	if todo.UserID != userID {
		reply.Err(w, ae.Forbidden(""))
		return
	}

	err = todo.Delete()
	if err != nil {
		reply.Err(w, ae.DB("", err))
		return
	}

	result := struct {
		Todo model.Todo `json:"todo"`
	}{todo}

	reply.OK(w, result)
}
