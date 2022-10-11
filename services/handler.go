package services

import (
	"fmt"
	"net/http"
	"path"
	"strconv"
)

func IconHandler(w http.ResponseWriter, r *http.Request) {}
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	var lang = Lang{Source: "id"}
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

	var data = map[string]string{
		"WaterValue":  strconv.Itoa(res.Status.Water),
		"WaterStatus": waterRules(lang, res.Status.Water),
		"WindValue":   strconv.Itoa(res.Status.Wind),
		"WindStatus":  windRules(lang, res.Status.Wind),
	}

	fmt.Println(data)
	err = executeHtml(w, pathHtml, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
