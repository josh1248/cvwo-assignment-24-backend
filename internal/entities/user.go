package entities

// db tag marshalls SQL data into Golang data with sqlx.Select,
// json tag marshalls Golang data to JSON with go-gin.
type User struct {
	ID         int    `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Reputation int    `db:"reputation" json:"reputation"`
	Password   string `db:"password" json: "password"`
}
