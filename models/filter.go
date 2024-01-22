package models

type Filter struct {
	Set         int
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Nationality string `json:"nationality"`
	Gender      string `json:"gender"`
	Age         string `json:"age"`
	Offset      string
	Limit       string
}
