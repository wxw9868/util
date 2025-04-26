package jwt

import (
	"testing"
)

func TestCreateToken(t *testing.T) {
	token, err := CreateToken("part_time_job", "18201108862", 1, false)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)

	r, err := ParseToken(token, "part_time_job")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", r)
}
