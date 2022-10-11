package services

import (
	"fmt"
	"net/http"
	"path"
)

func IconHandler(w http.ResponseWriter, r *http.Request) {}
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var lang = Lang{Source: "id"}
	var pathJson = path.Join("services", "data.json")

	var res = Response{
		Status{
			Water: getRandomNumberInt(1, 100),
			Wind:  getRandomNumberInt(1, 100),
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

	var water, wind = getTemplateData(lang, res)
	var data = map[string]Template{
		"water": water,
		"wind":  wind,
	}

	fmt.Println(data)
	err = executeHtml(w, path.Join("views", "index.html"), data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
