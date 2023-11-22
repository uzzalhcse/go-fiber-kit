// Services/auth_service.go

package Services

import "github.com/uzzalhcse/amadeus-go/app/Models"

type AuthService interface {
	Authenticate(username, password string) (bool, error)
	GetUserByUsername(username string) (*Models.User, error)
	Register(user *Models.User) error
	UpdateProfile(username string, updatedUser *Models.User) error
	ForgetPassword(username string) (string, error)
	// Add other authentication-related methods as needed
}
