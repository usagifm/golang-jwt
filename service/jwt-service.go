package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWT Service is a contract of what jwtservice can do
type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

// NewJWTService method creates a new instance of JWTService
func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "usagifm",
		secretKey: getSecretKey(),
	}

}

func getSecretKey() string {
	secretKey := "wdape12o3eo1je912"
	// if secretKey != "" {
	// 	secretKey = "usagifm"

	// }

	return secretKey
}

func (j *jwtService) GenerateToken(UserID string) string {
	claims := jwtCustomClaim{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("Unexpected signing method %w", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil

	})

}
