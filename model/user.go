package model

import (
	"crypto/sha512"
	"encoding/base64"
	"strings"
	"time"

	"github.com/mantishK/galore/config"
)

type User struct {
	UserId   int       `json:"user_id" id:TRUE auto:TRUE`
	UserName string    `json:"user_name"`
	Password string    `json:"password"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

func (u User) Public() interface{} {
	userFormat := make(map[string]User)
	u.Password = ""
	userFormat["user"] = u
	return userFormat
}

func (u *User) Get() error {
	err := config.DB.QueryRow("SELECT * FROM users WHERE user_id = $1", u.UserId).Scan(&u.UserId, &u.UserName, &u.Password, &u.Created, &u.Modified)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetUserFromUserName() error {
	err := config.DB.QueryRow("SELECT * FROM users WHERE user_name = $1", u.UserName).Scan(&u.UserId, &u.UserName, &u.Password, &u.Created, &u.Modified)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Save() error {
	u.Created = time.Now()
	u.Modified = time.Now()
	err := insertUser(u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Update() error {
	u.Created = time.Now()
	u.Modified = time.Now()
	err := updateUser(u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) UserNameExists() (bool, error) {
	count := 0
	err := config.DB.QueryRow("SELECT count(*) FROM users WHERE user_name = $1", u.UserName).Scan(&count)
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	u.Password = ""
	return true, nil
}

func (u *User) IsValidUser() (bool, error) {
	err := config.DB.QueryRow("SELECT * FROM users WHERE user_name = $1 AND password = $2", u.UserName, u.Password).Scan(&u.UserId, &u.UserName, &u.Password, &u.Created, &u.Modified)
	if err != nil {
		return false, err
	}
	if u.UserId == 0 {
		return false, nil
	}
	return true, nil
}

func (u *User) IsValidPassword() (bool, error) {
	err := config.DB.QueryRow("SELECT * FROM users WHERE user_id = $1 AND password = $2", u.UserId, u.Password).Scan(&u.UserId, &u.UserName, &u.Password, &u.Created, &u.Modified)
	if err != nil {
		return false, err
	}
	if u.UserId == 0 {
		return false, nil
	}
	return true, nil
}

func (u *User) UpdatePassword() error {
	u.Modified = time.Now()
	err := updateUser(u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetPasswordSalt() (string, error) {
	password := ""
	err := config.DB.QueryRow("SELECT password FROM users WHERE user_id = $1", u.UserId).Scan(&password)
	if err != nil {
		return "", nil
	}
	passwordSlice := strings.Split(password, ".")
	salt := passwordSlice[len(passwordSlice)-1]
	return salt, nil
}

func (u *User) HashPassword(salt string) {
	password := u.Password
	sha := sha512.New()
	if len(salt) == 0 {
		salt = generateRandomString(8)
	}
	for i := 0; i < 4; i++ {
		result := base64.StdEncoding.EncodeToString(sha.Sum([]byte(password + salt)))
		password = string(result)
	}
	u.Password = password + "." + salt
}

func insertUser(u *User) error {
	err := config.DB.QueryRow("INSERT INTO users (user_name,password,created,modified) VALUES ($1,$2,$3,$4) returning user_id", u.UserName, u.Password, u.Created, u.Modified).Scan(&u.UserId)
	if err != nil {
		return err
	}
	return nil
}
func updateUser(u *User) error {
	_, err := config.DB.Exec("UPDATE users SET user_name = $1,password = $2,created = $3,modified = $4 WHERE user_id = $5", u.UserName, u.Password, u.Created, u.Modified, u.UserId)
	if err != nil {
		return err
	}
	return nil
}
