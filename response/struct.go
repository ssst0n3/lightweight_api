package response

type Base struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type CreateSuccess struct {
	Base
	Id uint `json:"id"`
}

type Auth struct {
	Base
	Auth bool `json:"auth"`
}
