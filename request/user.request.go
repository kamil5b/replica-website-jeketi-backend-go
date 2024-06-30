package request

type LoginRequest struct {
	Email    string `json:"Email"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Password string `json:"password"`
	Email    string `json:"Email"`
}
type DeleteUserRequest struct {
	Password string
}
type ChangePasswordRequest struct {
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}
