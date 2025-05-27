package jwt_service

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

const (
	ErrInvalidToken = "INVALID_TOKEN"
)

type JWTService struct {
	secretKey []byte
}

func New(secretKey string) *JWTService {
	return &JWTService{
		secretKey: []byte(secretKey),
	}
}

func (jwts *JWTService) GetJwtSecretKey() []byte {
	return jwts.secretKey
}

func (jwts *JWTService) CreateToken(payload jwt.MapClaims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	t, err := token.SignedString(jwts.secretKey)
	if err != nil {
		return "", err
	}
	return t, nil
}
func (jwts *JWTService) GetUserID(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwts.secretKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["user_id"].(string)
		return userID, nil
	}

	return "", fmt.Errorf("invalid token")
}
func (jwts *JWTService) GetSessionID(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwts.secretKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sessionID := claims["session_id"].(string)
		return sessionID, nil
	}

	return "", fmt.Errorf("INV")
}
