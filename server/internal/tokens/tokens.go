package tokens

import (
	"fmt"
	"mtsaudio/internal/entities"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateNewJwts(user entities.User) (string, string, error) {
	op := "usecase.token.GenerateNewJwts()"
	key := []byte("boba")
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":        user.Id,
			"username":  user.Username,
			"createdAt": int(time.Now().Unix()),
			"expIn":     108000,
		})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":        user.Id,
			"username":  user.Username,
			"createdAt": int(time.Now().Unix()),
			"expIn":     2592000,
		})

	strAccessToken, err := accessToken.SignedString(key)
	strRefreshToken, err := refreshToken.SignedString(key)
	if err != nil {
		return "", "", fmt.Errorf("%s: failed to sign jwt: %w", op, err)
	}
	return strAccessToken, strRefreshToken, nil
}

func ParseToken(tokenString string) (entities.Token, error) {
	op := "tokens.ParseToken()"
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%s: unexpected signing method: %v", op, t.Header["alg"])
		}
		return []byte("boba"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return entities.Token{
			Id:        uint(claims["id"].(float64)),
			Username:  claims["username"].(string),
			CreatedAt: int(claims["createdAt"].(float64)),
			ExpIn:     int(claims["expIn"].(float64)),
		}, nil
	}

	return entities.Token{}, fmt.Errorf("%s: failed to parse token: %w", op, err)
}

// black list
var blackList map[string]bool

func InitBlackList() {
	blackList = make(map[string]bool)
}

func RemoveToken(token string) {
	blackList[token] = true
}

func IsInBlackList(token string) bool {
	if val, ok := blackList[token]; val && ok {
		return true
	}
	return false
}
