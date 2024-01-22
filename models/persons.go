package models

type Person struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Nationality string `json:"nationality"`
	Gender      string `json:"gender"`
	Age         int    `json:"age"`
}
