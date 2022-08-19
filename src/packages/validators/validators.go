package validatorpkg

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"myclass/src/app"
	"strings"
)

type Message = map[string]string

func Exec(data interface{}) error {

	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		return err
	}

	return nil

}

func ExecWithMessage(data interface{}, message Message) error {
	validate := validator.New()
	err := validate.Struct(data)

	if err == nil {
		return nil
	}

	for _, e := range err.(validator.ValidationErrors) {
		result, ok := message[fmt.Sprintf("%v.%v", e.Field(), e.Tag())]
		if !ok {
			return e
		}
		return errors.New(result)
	}
	return nil

}

func ExecWithKeyError(data interface{}, key string) error {
	validate := validator.New()
	err := validate.Struct(data)

	if err == nil {
		return nil
	}

	for _, e := range err.(validator.ValidationErrors) {
		key = fmt.Sprintf("%v.%v_%v", key, strings.ToLower(e.Field()), e.Tag())
		if len(e.Param()) > 0 {
			key = fmt.Sprintf("%v_%v", key, e.Param())
		}
		return errors.New(app.NewError(key).Message)
	}

	return nil
}
