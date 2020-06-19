package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ssst0n3/awesome_libs/cipher"
	"time"
)

const PanicJwtSecretHasNotBeenInited = "JwtSecretHasNotBeenInited"

var JwtSecret []byte

func GenerateToken(userId uint, isAdmin bool, duration time.Duration) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(duration)
	claims := Claims{
		UserId:  userId,
		IsAdmin: isAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret, err := getSecretKey(&jwt.Token{})
	if err != nil {
		return "", err
	}
	token, err := tokenClaims.SignedString(secret)
	return token, err
}

func getSecretKey(_ *jwt.Token) (interface{}, error) {
	if len(JwtSecret) == 0 {
		panic(PanicJwtSecretHasNotBeenInited)
	}
	return JwtSecret, nil
}

func CheckPassword(inputPassword string, password string) (bool, error) {
	// TODO: change aes256 to pbkdf2
	passwordDecrypted, err := cipher.CommonCipher.Decrypt(password)
	if err != nil {
		return false, err
	}
	return inputPassword == passwordDecrypted, nil
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, getSecretKey)
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
