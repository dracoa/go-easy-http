package easy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type H map[string]interface{}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func Recover(w http.ResponseWriter) {
	if r := recover(); r != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("%v", r)))
		return
	}
}

func PathParams(req *http.Request) []string {
	return strings.Split(req.URL.Path, "/")[1:]
}

func BearerToken(req *http.Request) string {
	return strings.Replace(req.Header.Get("Authorization"), "Bearer ", "", 0)
}

func UserAgent(req *http.Request) string {
	return req.Header.Get("User-Agent")
}

func ParseJSON(req *http.Request, placeholder interface{}) {
	bytes, err := ioutil.ReadAll(req.Body)
	PanicIf(err)
	err = json.Unmarshal(bytes, &placeholder)
	PanicIf(err)
}

func WriteJSON(w http.ResponseWriter, status int, j map[string]interface{}) {
	bytes, err := json.Marshal(j)
	PanicIf(err)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(bytes)
	PanicIf(err)
}

func WriteText(w http.ResponseWriter, status int, t string) {
	w.WriteHeader(status)
	_, err := w.Write([]byte(t))
	PanicIf(err)
}
