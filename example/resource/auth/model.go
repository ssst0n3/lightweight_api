package auth

type LoginSuccessResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
	UserId   uint   `json:"user_id"`
}
