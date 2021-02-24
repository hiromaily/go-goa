package jwt

import (
	"fmt"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"

	c "github.com/hiromaily/go-goa/ext/context"
)

// GenerateToken is to generate token
func GenerateToken(c *c.Ctx, userID int) (string, error) {
	token := jwtgo.New(jwtgo.SigningMethodHS512)
	expires := time.Now().Add(time.Duration(60) * time.Minute).Unix()

	token.Claims = jwtgo.MapClaims{
		"iss":    "Go-Goa API",      // who creates the token and signs it
		"aud":    "audience(s)",     // to whom the token is intended to be sent
		"exp":    expires,           // time when the token will expire (10 minutes from now)
		"jti":    uuid.NewString(),  // a unique identifier for the token
		"iat":    time.Now().Unix(), // when the token was issued/created (now)
		"nbf":    2,                 // time before which the token is not yet valid (2 minutes ago)
		"sub":    "subject",         // the subject/principal is whom the token is about
		"scopes": "api:access",      // token scope - not a standard claim
		"user":   userID,            // userID
	}
	//TODO: key
	signedToken, err := token.SignedString([]byte(c.Conf.Jwt.Key))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %s", err) // internal error
	}
	return signedToken, nil
}
