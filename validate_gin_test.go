package util

import (
	"errors"
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	v := NewValidate("json")
	v.initValidateGin()
	s := v.GinError(errors.New("error"))
	fmt.Println(s)
	myEmail := "wxw9868"
	f := v.GinVar(myEmail, "required,email")
	fmt.Println(f)
}
