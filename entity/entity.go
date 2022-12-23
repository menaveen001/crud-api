package entity

type User struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Dob         string `json:"dob"`
	Address     string `json:"address"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdat"`
}
