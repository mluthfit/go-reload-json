package services

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type Response struct {
	Status Status `json:"status"`
}

type Template struct {
	Value  int
	Status string
	Class  string
}
