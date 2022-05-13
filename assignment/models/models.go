package models

type Patient struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Covid bool   `json:"covid"`
}
