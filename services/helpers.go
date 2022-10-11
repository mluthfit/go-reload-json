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

func executeHtml(w http.ResponseWriter, path string, data map[string]Template) error {
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

func waterRules(value int) string {
	var status = "safe"

	if value > 8 {
		status = "danger"
	} else if value > 5 {
		status = "standby"
	}

	return status
}

func windRules(value int) string {
	var status = "safe"

	if value > 15 {
		status = "danger"
	} else if value > 6 {
		status = "standby"
	}

	return status
}

func statusRules(category string, value int) (status string) {
	if category == "water" {
		return waterRules(value)
	}

	return windRules(value)
}

func getTemplateData(lang Lang, res Response) (Template, Template) {
	var statusWater = statusRules("water", res.Status.Water)
	var statusWind = statusRules("wind", res.Status.Wind)

	var water = Template{
		Value:  res.Status.Water,
		Status: lang.GetStatusLang(statusWater),
		Class:  statusWater,
	}

	var wind = Template{
		Value:  res.Status.Wind,
		Status: lang.GetStatusLang(statusWind),
		Class:  statusWind,
	}

	return water, wind
}
