package util

import "github.com/golang-jwt/jwt"

func GenerateToken(username string) string {
	type customClaims struct {
		Username string `json:username`
		jwt.StandardClaims
	}
	claims := customClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			Issuer: "noteify",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString([]byte("noteify"))

	return string(signedToken)
}
