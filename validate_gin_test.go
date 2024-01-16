package util

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewValidate_Gin(t *testing.T) {
	v := NewValidate("json")
	v.InitValidateGin()
	s := v.GinError(errors.New("error"))
	fmt.Println(s)
	myEmail := "wxw9868"
	f := v.GinVar(myEmail, "required,email")
	fmt.Println(f)
}
