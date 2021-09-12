package jwt

import (
	"fmt"
	"testing"
)

func TestCreateToken(t *testing.T) {
	token, err := CreateToken("part_time_job", "18201108862", uint(1), false)
	fmt.Println(token)
	fmt.Println(err)
}

func TestParseToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTQ2NTE2NjcsImlzcyI6IjE4MjAxMTA4ODYyIiwidWlkIjoxLCJhZG1pbiI6ZmFsc2V9.h_UFhxRrOhX1nPrMzS1L2QbY_uRvwqqPGXSiqIF8FNE"
	r, err := ParseToken(token, "part_time_job")
	if err != nil {
		if err.Error() == "Token is expired" {
			fmt.Println("授权已过期")
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println(r)
}
