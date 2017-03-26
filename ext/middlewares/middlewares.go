package middlewares

import (
	jwtgo "github.com/dgrijalva/jwt-go"
	//"context"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	"golang.org/x/net/context"
	"net/http"
)

func AuthMiddlewareHandler(nextHandler goa.Handler) goa.Handler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		token := jwt.ContextJWT(ctx)

		if val, ok := token.Claims.(jwtgo.MapClaims)["scopes"].(string); !ok {
			w.WriteHeader(401)
			return jwt.ErrJWTError("Token is expired or not valid")
		} else if val != "api:access" {
			return jwt.ErrJWTError("you are not allowed to access api")
		}

		if val, ok := token.Claims.(jwtgo.MapClaims)["user"].(map[string]interface{}); ok {
			fmt.Println("token.Claims:", val)
			ctx = context.WithValue(ctx, "user.jwt", val)
		}
		return nextHandler(ctx, w, r)
	}
}
