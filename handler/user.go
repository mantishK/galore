package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	ae "github.com/mantishK/galore/apperror"
	"github.com/mantishK/galore/model"
	"github.com/mantishK/galore/reply"
	"github.com/mantishK/galore/validate"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)
	var err error

	id := r.URL.Query().Get("id")
	if len(id) != 0 {
		userID, err = strconv.Atoi(id)
		if err != nil {
			reply.Err(w, ae.InvalidInput("id is not a number", "id"))
			return
		}
	}

	user := model.User{}
	user.UserId = userID
	err = user.Get()
	if err != nil {
		reply.Err(w, ae.DB("", err))
		return
	}

	reply.OK(w, user)
}

func UserNameExists(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Query().Get("user_name")
	if len(userName) == 0 {
		reply.Err(w, ae.Required("", "user_name"))
		return
	}
	user := model.User{}
	user.UserName = userName
	exists, err := user.UserNameExists()
	if err != nil {
		reply.Err(w, ae.DB("", err))
		return
	}

	result := make(map[string]interface{})
	result["exists"] = exists
	reply.OK(w, result)
}

type userReqBody struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (r *userReqBody) OK() *ae.Error {
	if len(r.UserName) == 0 {
		return ae.Required("", "user_name")
	} else if len(r.Password) == 0 {
		return ae.Required("", "password")
	} else if !validate.UserName(r.UserName) {
		return ae.InvalidInput("", "user_name")
	} else if !validate.Password(r.Password) {
		return ae.InvalidInput("", "password")
	}
	return nil
}

func SaveUser(w http.ResponseWriter, r *http.Request) {
	reqBody := userReqBody{}
	appErr := decode(r, &reqBody)
	if appErr != nil {
		reply.Err(w, appErr)
		return
	}
	user := model.User{}
	user.UserName = reqBody.UserName
	user.Password = reqBody.Password
	exists, err := user.UserNameExists()
	if exists {
		reply.Err(w, ae.UserNameExists("", "user_name"))
		return
	} else if err != nil {
		reply.Err(w, ae.DB("", err))
		return
	}
	user.HashPassword("")
	err = user.Save()
	if err != nil {
		reply.Err(w, ae.DB("", err))
		return
	}

	reply.OK(w, user)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	reqBody := userReqBody{}
	appErr := decode(r, &reqBody)
	if appErr != nil {
		reply.Err(w, appErr)
		return
	}

	userToken, appErr := signinWithUserName(reqBody)
	if appErr != nil {
		reply.Err(w, appErr)
		return
	}
	result := make(map[string]interface{})
	result["user_token"] = userToken
	reply.OK(w, result)
}

func signinWithUserName(reqBody userReqBody) (string, *ae.Error) {
	user := model.User{}
	user.UserName = reqBody.UserName
	err := user.GetUserFromUserName()
	if err == sql.ErrNoRows {
		return "", ae.InvalidUserNamePassword("")
	}
	if err != nil {
		return "", ae.DB("", err)
	}
	salt, err := user.GetPasswordSalt()
	if err != nil {
		return "", ae.DB("", err)
	}
	user.Password = reqBody.Password
	user.HashPassword(salt)
	exists, err := user.IsValidUser()
	if err != nil || !exists {
		return "", ae.InvalidUserNamePassword("")
	}

	userToken := model.UserToken{}
	userToken.UserId = user.UserId
	err = userToken.Add()
	if err != nil {
		return "", ae.DB("", err)
	}
	return userToken.Token, nil
}

func SignOut(w http.ResponseWriter, r *http.Request) {
	userToken := model.UserToken{}
	userToken.Token = r.Context().Value("user_token").(string)
	err := userToken.Delete()
	if err != nil {
		reply.Err(w, ae.DB("", err))
		return
	}
	result := make(map[string]interface{})
	result["user_token"] = userToken.Token
	reply.OK(w, result)
}
