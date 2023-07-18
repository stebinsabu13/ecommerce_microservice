package support

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stebinsabu13/ecommerce_microservice/auth_service/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	ID uint
	jwt.RegisteredClaims
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// password authorization
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}

func ReferalCodeGenerator() string {
	uuidObj := uuid.New()
	uuidString := strings.Replace(uuidObj.String(), "-", "", -1)
	return uuidString[:6]
}

func GenerateJWT(id uint) (string, error) {

	expireTime := time.Now().Add(1 * time.Hour)

	// create token with expire time and claims id as user id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	})

	// conver the token into signed string
	tokenString, err := token.SignedString([]byte(config.GetCofig().Jwt_secret_key))

	if err != nil {
		return "", err
	}
	// refresh token add next time
	return tokenString, nil
}

func ValidateToken(tokenString string) (Claims, error) {
	claims := Claims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetCofig().Jwt_secret_key), nil
		},
	)
	//checking the expiry of the token
	if time.Now().Unix() > claims.ExpiresAt.Unix() {
		return claims, errors.New("token expired re-login")
	}
	if err != nil || !token.Valid {
		return claims, errors.New("not valid token")
	}
	return claims, nil
}
