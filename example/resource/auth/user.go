package auth

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

type UserWithId struct {
	Id uint `json:"id"`
	User
}

const (
	TableNameUser          = "user"
	ColumnNameUserUsername = "username"
)

