// Controllers/auth_controller.go

package Controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uzzalhcse/amadeus-go/app/Http/Requests"
	"github.com/uzzalhcse/amadeus-go/app/Http/Responses"
	"github.com/uzzalhcse/amadeus-go/app/Models"
	"github.com/uzzalhcse/amadeus-go/app/Repositories"
	"github.com/uzzalhcse/amadeus-go/app/Services"
	"log"
)

type AuthController struct {
	BaseController *BaseController
	AuthService    Services.AuthService
	JWTService     Services.JWTService
}

func NewAuthController() *AuthController {
	that := NewBaseController()
	authRepo := Repositories.NewAuthRepository(that.DB)
	authService := Services.NewAuthService(authRepo)
	jwtService := Services.NewJWTService(that.Config.App.JwtSecret)

	return &AuthController{
		BaseController: that,
		AuthService:    authService,
		JWTService:     jwtService,
	}
}

// Login handles the login route
func (that *AuthController) Login(c *fiber.Ctx) error {
	var request Requests.LoginRequest

	if err := Requests.Validate(c, &request); err != nil {
		return Responses.Error(c, err.Error(), nil)
	}

	// Authenticate user
	authenticated, err := that.AuthService.Authenticate(request.Username, request.Password)
	if err != nil {
		return Responses.Error(c, "Authentication failed", nil)
	}

	if !authenticated {
		return Responses.Error(c, "Invalid credentials", nil)
	}

	// Generate JWT token
	user, _ := that.AuthService.GetUserByUsername(request.Username) // Assuming you have a GetUserByUsername method
	token, err := that.JWTService.GenerateToken(user)
	if err != nil {
		return Responses.Error(c, "Failed to generate token", nil)
	}

	// Send JWT token in the response
	return Responses.Success(c, fiber.Map{
		"message": "Login successful",
		"token":   token,
	})
}

// Register handles the registration route
func (that *AuthController) Register(c *fiber.Ctx) error {
	var request Requests.RegisterRequest

	if err := Requests.Validate(c, &request); err != nil {
		return Responses.Error(c, err.Error(), nil)
	}

	user := &Models.User{
		Name:     request.Name,
		Email:    request.Email,
		Username: request.Username,
		Password: request.Password,
		// Add other user properties as needed
	}

	if err := that.AuthService.Register(user); err != nil {
		return Responses.Error(c, "Registration failed", nil)
	}

	return Responses.Success(c, fiber.Map{"message": "Registration successful"})
}

// UpdateProfile handles the update profile route
func (that *AuthController) UpdateProfile(c *fiber.Ctx) error {
	var request Requests.UpdateProfileRequest

	if err := Requests.Validate(c, &request); err != nil {
		return Responses.Error(c, err.Error(), nil)
	}

	// Assuming you have a way to identify the current user (e.g., from the JWT token)
	username := "current_username"

	updatedUser := &Models.User{
		// Update user properties as needed
	}

	if err := that.AuthService.UpdateProfile(username, updatedUser); err != nil {
		return Responses.Error(c, "Profile update failed", nil)
	}

	return Responses.Success(c, fiber.Map{"message": "Profile updated successfully"})
}

// ForgetPasswordHandler handles the forget password route
func (that *AuthController) ForgetPassword(c *fiber.Ctx) error {
	var request Requests.ForgetPasswordRequest

	if err := Requests.Validate(c, &request); err != nil {
		return Responses.Error(c, err.Error(), nil)
	}

	// Assuming you have a way to identify the current user (e.g., from the JWT token)
	username := "current_username"

	resetToken, err := that.AuthService.ForgetPassword(username)
	if err != nil {
		return Responses.Error(c, "Failed to initiate password reset", nil)
	}
	log.Println(resetToken)
	// Send resetToken to the user (e.g., via email)

	return Responses.Success(c, fiber.Map{"message": "Password reset initiated"})
}
