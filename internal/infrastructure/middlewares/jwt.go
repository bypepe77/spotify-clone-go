package middlewares

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type musicClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "Test",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) GenerateToken(userID int) (string, error) {
	claims := &musicClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, errors.New("invalid token")
		}
		return []byte(service.secretKey), nil
	})

}

func FromMapClaims(claims jwt.MapClaims) (*musicClaims, error) {
	b, err := json.Marshal(claims)
	if err != nil {
		return nil, err
	}
	mcClaims := musicClaims{}
	err = json.Unmarshal(b, &mcClaims)
	return &mcClaims, err
}

func GetClaims(c *gin.Context) (*musicClaims, error) {
	anyClaims, ok := c.Get("claims")
	if !ok {
		return nil, errors.New("error with claims in context")
	}
	claims, ok := anyClaims.(*musicClaims)
	if !ok || claims == nil {
		return nil, errors.New("error type of claims")
	}
	return claims, nil
}
