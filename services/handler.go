package services

import (
	"fmt"
	"net/http"
	"path"
)

func IconHandler(w http.ResponseWriter, r *http.Request) {}
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	var pathJson = path.Join("services", "data.json")
	var pathHtml = path.Join("views", "index.html")
	var waterValue = getRandomNumberInt(1, 100)
	var windValue = getRandomNumberInt(1, 100)

	var res = Response{
		Status{
			Water: waterValue,
			Wind:  windValue,
		},
	}

	err = writeJson(res, pathJson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = readJson(&res, pathJson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(res)
	err = executeHtml(w, pathHtml, res.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
