package entities

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Reputation int    `json:"reputation"`
}
