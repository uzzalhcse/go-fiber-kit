package auth

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=5,max=50" error_ms:"Username is required and must be between 5 and 50 characters."`
	Password string `json:"password" validate:"required,min=6" error:"Password is required and must be at least 8 characters."`
}
