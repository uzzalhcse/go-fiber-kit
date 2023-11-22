package Services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/uzzalhcse/amadeus-go/app/Models"
	"time"
)

type JWTServiceImpl struct {
	SecretKey string
}

func NewJWTService(secretKey string) *JWTServiceImpl {
	return &JWTServiceImpl{SecretKey: secretKey}
}

func (s *JWTServiceImpl) GenerateToken(user *Models.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":   user.ID,
		"iss":   "your-issuer",                         // Replace with your desired issuer
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (adjust as needed)
		"iat":   time.Now().Unix(),
		"email": user.Email,
		// Add other custom claims as needed
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
