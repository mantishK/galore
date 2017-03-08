package model

import (
	"time"

	"github.com/mantishK/galore/config"
)

type UserToken struct {
	UserId   int       `db:"user_id" json:"user_id" id:"TRUE"`
	Token    string    `db:"token" json:"token"`
	Created  time.Time `db:"created" json:"created"`
	Modified time.Time `db:"modified" json:"modified"`
}

func (ut *UserToken) Public() interface{} {
	token := make(map[string]string)
	token["token"] = ut.Token
	return token
}

func (ut *UserToken) GetUserIdFromToken() error {
	err := config.DB.QueryRow("SELECT * FROM user_token WHERE token = $1", ut.Token).Scan(&ut.UserId, &ut.Token, &ut.Created, &ut.Modified)
	return err
}

func (ut *UserToken) Add() error {
	ut.Created = time.Now()
	ut.Modified = time.Now()
	ut.Token = generateRandomToken()
	err := insertUserToken(ut)
	if err != nil {
		return err
	}
	return nil
}

func (ut *UserToken) Delete() error {
	_, err := config.DB.Exec("DELETE from user_token where token = $1", ut.Token)
	if err != nil {
		return err
	}
	return nil
}

func (ut *UserToken) Update() error {
	ut.Modified = time.Now()
	token := generateRandomToken()
	_, err := config.DB.Exec("UPDATE user_token SET token = $1, modified = $2 WHERE user_id= $3 AND token = $4", token, ut.Modified, ut.UserId, ut.Token)
	if err != nil {
		return err
	}
	ut.Token = token
	return nil
}

func generateRandomToken() string {
	r := ""
	for {
		r = generateRandomString(80)
		user_token_test := UserToken{}
		user_token_test.Token = r
		err := user_token_test.GetUserIdFromToken()
		if err != nil && user_token_test.UserId == 0 {
			break
		}
	}
	return r
}

func insertUserToken(ut *UserToken) error {
	_, err := config.DB.Exec("INSERT INTO user_token (user_id,token,created,modified) VALUES ($1,$2,$3,$4)", ut.UserId, ut.Token, ut.Created, ut.Modified)
	return err
}
func updateUserToken(ut *UserToken) error {
	_, err := config.DB.Exec("UPDATE user_token SET token = $1,created = $2,modified = $3 WHERE user_id = $4", ut.Token, ut.Created, ut.Modified, ut.UserId)
	return err
}
