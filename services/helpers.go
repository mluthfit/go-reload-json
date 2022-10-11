package services

import (
	"encoding/json"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getRandomNumberInt(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func writeJson(res Response, path string) error {
	var result, err = json.MarshalIndent(res, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(path, result, 0644)
	if err != nil {
		return err
	}

	return nil
}

func readJson(res *Response, path string) error {
	var content, err = os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, res)
	if err != nil {
		return err
	}

	return nil
}

func executeHtml(w http.ResponseWriter, path string, data map[string]string) error {
	var template, err = template.ParseFiles(path)
	if err != nil {
		return err
	}

	err = template.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}

func windRules(lang Lang, value int) (result string) {
	switch {
	case value > 15:
		result = lang.GetStatusLang("danger")
	case value > 6:
		result = lang.GetStatusLang("standby")
	default:
		result = lang.GetStatusLang("safe")
	}

	return
}

func waterRules(lang Lang, value int) (result string) {
	switch {
	case value > 8:
		result = lang.GetStatusLang("danger")
	case value > 5:
		result = lang.GetStatusLang("standby")
	default:
		result = lang.GetStatusLang("safe")
	}

	return
}
