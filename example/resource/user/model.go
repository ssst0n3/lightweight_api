package user

type UpdateBasicBody struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
}

/*
need encrypt
*/
type UpdatePasswordBody struct {
	Password string `json:"password"`
}

type Model struct {
	UpdateBasicBody
	UpdatePasswordBody
}

type ModelWithId struct {
	Id uint `json:"id"`
	Model
}

const (
	ColumnNameUsername = "username"
	ColumnNameIsAdmin  = "is_admin"
)

// ======response model========
type ListUserBody struct {
	Id uint `json:"id"`
	UpdateBasicBody
}