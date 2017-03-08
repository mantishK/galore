package apperror

import "net/http"

type Error struct {
	Message    string `json:"string"`
	ID         int    `json:"id"`
	HttpStatus int    `json:"-"`
	Log        string `json:"log"`
}

const (
	CodeDB = iota + 1
	CodeRequired
	CodeInvalidInput
	CodeNotNumeric
	CodeTokenInvalid
	CodeTokenExpired
	CodeUserNameExists
	CodeInvalidUserNamePassword
	CodeInvalidUserName
	CodeInvalidPassword
	CodeResourceNotFound
	CodeFileUpload
	CodeJsonEncode
	CodeJsonDecode
	CodeForbidden
	CodeInternal
)

func DB(message string, err error) *Error {
	if len(message) == 0 {
		message = "Some error occured while querying database. Please try later."
	}
	return &Error{ID: CodeDB, Message: message, HttpStatus: http.StatusInternalServerError, Log: err.Error()}
}

func Required(message string, field string) *Error {
	if len(message) == 0 {
		message = "A required field is missing"
	}
	return &Error{ID: CodeRequired, Message: message, HttpStatus: http.StatusBadRequest}
}

func InvalidInput(message string, field string) *Error {
	if len(message) == 0 {
		message = "A field is invalid"
	}
	return &Error{ID: CodeInvalidInput, Message: message, HttpStatus: http.StatusBadRequest}
}

func NotNumericInput(message string, err error, field string) *Error {
	if len(message) == 0 {
		message = "A field has non-numeric chars"
	}
	return &Error{ID: CodeNotNumeric, Message: message, Log: err.Error(), HttpStatus: http.StatusBadRequest}
}

func TokenInvalid(message string, err error, field string) *Error {
	if len(message) == 0 {
		message = "Invalid token"
	}
	return &Error{ID: CodeTokenInvalid, Message: message, Log: err.Error(), HttpStatus: http.StatusBadRequest}
}

func TokenExpired(message string) *Error {
	if len(message) == 0 {
		message = "Token expired"
	}
	return &Error{ID: CodeTokenExpired, Message: message, HttpStatus: http.StatusBadRequest}
}

func UserNameExists(message string, field string) *Error {
	if len(message) == 0 {
		message = "User Name exists"
	}
	return &Error{ID: CodeUserNameExists, Message: message, HttpStatus: http.StatusBadRequest}
}
func InvalidUserNamePassword(message string) *Error {
	if len(message) == 0 {
		message = "Email ID/Password Invalid"
	}
	return &Error{ID: CodeInvalidUserNamePassword, Message: message, HttpStatus: http.StatusBadRequest}
}
func InvalidUserName(message string) *Error {
	if len(message) == 0 {
		message = "User Name Invalid"
	}
	return &Error{ID: CodeInvalidUserName, Message: message, HttpStatus: http.StatusBadRequest}
}

func InvalidPassword(message string, field string) *Error {
	if len(message) == 0 {
		message = "User Password Invalid"
	}
	return &Error{ID: CodeInvalidPassword, Message: message, HttpStatus: http.StatusBadRequest}
}

func ResourceNotFound(message string) *Error {
	if len(message) == 0 {
		message = "Resource Not Found"
	}
	return &Error{ID: CodeResourceNotFound, Message: message, HttpStatus: http.StatusNotFound}
}

func FileUpload(message string, field string, err error) *Error {
	if len(message) == 0 {
		message = "Unable to upload the file"
	}
	return &Error{ID: CodeFileUpload, Message: message, Log: err.Error(), HttpStatus: http.StatusNotFound}
}

func JsonEncode(message string) *Error {
	if len(message) == 0 {
		message = "Something went wrong while encoding json"
	}
	return &Error{ID: CodeJsonEncode, Message: message, HttpStatus: http.StatusInternalServerError}
}

func JsonDecode(message string) *Error {
	if len(message) == 0 {
		message = "Something went wrong while decoding json"
	}
	return &Error{ID: CodeJsonDecode, Message: message, HttpStatus: http.StatusInternalServerError}
}

func Forbidden(message string) *Error {
	if len(message) == 0 {
		message = "Forbiden/Need to login"
	}
	return &Error{ID: CodeForbidden, Message: message, HttpStatus: http.StatusForbidden}
}

func Internal(message string, err error) *Error {
	if len(message) == 0 {
		message = "Forbiden/Need to login"
	}
	return &Error{ID: CodeInternal, Message: message, Log: err.Error(), HttpStatus: http.StatusForbidden}
}
