package easy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

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
	bytes, err:=ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bytes, &placeholder)
	if err != nil {
		panic(err)
	}
}
