package services

var StatusLanguage = map[string]map[string]string{
	"id": {
		"safe":    "aman",
		"standby": "siaga",
		"danger":  "bahaya",
	},
	"en": {
		"safe":    "safe",
		"standby": "standby",
		"danger":  "danger",
	},
}

type Lang struct {
	Source string
}

func (lang *Lang) GetStatusLang(key string) string {
	return StatusLanguage[lang.Source][key]
}
