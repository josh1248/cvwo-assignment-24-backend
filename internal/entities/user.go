package entities

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// placeholder DB.
var Users []User = []User{
	{ID: 123, Name: "Jojo"},
	{ID: 456, Name: "Hoho"},
	{ID: 789, Name: "Lolo"},
}
