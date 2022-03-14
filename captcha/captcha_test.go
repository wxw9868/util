package captcha

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
)

//start a net/http server
func TestWeb(t *testing.T) {
	//serve Vuejs+ElementUI+Axios Web Application
	http.Handle("/", http.FileServer(http.Dir("./static")))

	//api for create captcha
	http.HandleFunc("/api/getCaptcha", generateCaptchaHandler)

	//api for verify captcha
	http.HandleFunc("/api/verifyCaptcha", captchaVerifyHandle)

	fmt.Println("Server is at :8777")
	if err := http.ListenAndServe(":8777", nil); err != nil {
		log.Fatal(err)
	}
}

func generateCaptchaHandler(w http.ResponseWriter, r *http.Request) {
	body := GetCaptcha("string")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)
}

func captchaVerifyHandle(w http.ResponseWriter, r *http.Request) {
	body := Verify(r.FormValue("id"), r.FormValue("value"))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)
}

func TestGetCaptcha(t *testing.T) {
	body := GetCaptcha("string")
	fmt.Println(body)
}

func TestCaptchaVerify(t *testing.T) {
	result := Verify("", "")
	fmt.Println(result)
}
