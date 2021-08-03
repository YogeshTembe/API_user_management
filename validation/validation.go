package validation

import (
	"strconv"

	li "github.com/YogeshTembe/golang_project/logwrapper"
	"github.com/YogeshTembe/golang_project/model"
	uuid "github.com/satori/go.uuid"
)

var UserIds = make(map[string]struct{})

func New(text string) error {
	return &errorString{text}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

var StandardLogger = li.NewLogger()

func Validate(user *model.User) bool {
	isValid := true
	if user.Name == "" {
		StandardLogger.InvalidArg("name", user.Id.String())
		isValid = false
	}
	if user.Email == "" {
		StandardLogger.InvalidArg("Email-id", user.Id.String())
		isValid = false
	}
	if len(strconv.Itoa(user.PhoneNumber)) != 10 {
		StandardLogger.InvalidArg("phone number", user.Id.String())
		isValid = false
	}
	if user.Id.String() == "00000000-0000-0000-0000-000000000000" {
		user.Id = uuid.NewV4()
	}

	_, isFound := UserIds[user.Id.String()]
	if isFound {
		StandardLogger.InvalidArg("user-id already present for username-"+user.Name, user.Id.String())
		isValid = false
	}

	return (isValid)
}
