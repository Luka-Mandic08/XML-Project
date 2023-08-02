package model

import (
	"fmt"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	Username string `json:"username,omitempty"`
	Role     string `json:"roles,omitempty"`
	UserId   string `json:"user_id,omitempty"`
	jwt.StandardClaims
}

func (claims JwtClaims) Valid() error {
	var now = time.Now().UTC().Unix()
	if claims.VerifyExpiresAt(now, true) && claims.VerifyIssuer(issuer, true) {
		return nil
	}
	return fmt.Errorf("Token is invalid")
}

func (claims JwtClaims) VerifyAudience(origin string) bool {
	return strings.Compare(claims.Audience, origin) == 0
}
