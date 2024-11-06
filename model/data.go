package model

type Status struct {
	Status Data `json:"status"`
}

type Data struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}
