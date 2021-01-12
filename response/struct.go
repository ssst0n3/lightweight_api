package response

type Base struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type Err struct {
	Base
	Reason string `json:"reason"`
}

type CreateSuccess struct {
	Base
	Id uint `json:"id"`
}

type Auth struct {
	Base
	Auth bool `json:"auth"`
}
