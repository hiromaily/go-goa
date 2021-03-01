package jwts

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type SimpleJWTer interface {
	GenerateToken(userID int) (string, error)
}

type simpleJWT struct {
	key string
}

func NewSimpleJWT(key string) SimpleJWTer {
	return &simpleJWT{
		key: key,
	}
}

// GenerateToken generates token
func (j *simpleJWT) GenerateToken(userID int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)
	expires := time.Now().Add(time.Duration(60) * time.Minute).Unix()

	uid, err := uuid.NewUUID()
	if err != nil {
		return "", errors.New("fail to call uuid.NewUUID()")
	}

	token.Claims = jwt.MapClaims{
		"iss":    "go-goa API",      // who creates the token and signs it
		"aud":    "audience(s)",     // to whom the token is intended to be sent
		"exp":    expires,           // time when the token will expire (10 minutes from now)
		"jti":    uid,               // a unique identifier for the token
		"iat":    time.Now().Unix(), // when the token was issued/created (now)
		"nbf":    2,                 // time before which the token is not yet valid (2 minutes ago)
		"sub":    "subject",         // the subject/principal is whom the token is about
		"scopes": "api:access",      // token scope - not a standard claim
		"user":   userID,            // userID
	}
	signedToken, err := token.SignedString([]byte(j.key))
	if err != nil {
		return "", errors.Wrapf(err, "failed to call token.SignedString()") // internal error
	}
	return signedToken, nil
}
