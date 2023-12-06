package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func listAll(user string) map[string][]string {
	return map[string][]string{
		"tonybai": {
			"xxx.jpg",
			"yyy.jpg",
			"zzz.jpg",
		},
	}

}

func userInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("userInfoHandler:", *r)

	// check access_token

	// get user info by access_token
	user := "tonybai"
	w.Write([]byte(user))
	return
}

func photosHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("photosHandler:", *r)

	// check access_token

	// get user info by access_token
	user := "tonybai"

	method := r.FormValue("method")

	if method == "listall" {
		pl := listAll(user)
		v, err := json.Marshal(pl)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(v)
		return
	}

	w.WriteHeader(http.StatusNotImplemented)
}

func main() {
	http.HandleFunc("/photos", photosHandler)
	http.HandleFunc("/userinfo", userInfoHandler)

	fmt.Println("启动云盘服务器，监听8082端口")
	http.ListenAndServe(":8082", nil)
}
