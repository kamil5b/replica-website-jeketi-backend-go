package request

type DeleteUserRequest struct {
	Password string
}
type ChangePasswordRequest struct {
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}
