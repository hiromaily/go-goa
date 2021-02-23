package middlewares

//import (
//	"context"
//	"fmt"
//	jwtgo "github.com/dgrijalva/jwt-go"
//	goa "goa.design/goa/v3/pkg"
//	"goa.design/goa/v3/middleware/security/jwt"
//	"net/http"
//	//"reflect"
//)
//
//// AuthMiddlewareHandler is to check jwt
//func AuthMiddlewareHandler(nextHandler goa.Handler) goa.Handler {
//	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
//		fmt.Println("[JWT] AuthMiddlewareHandler()")
//		token := jwt.ContextJWT(ctx)
//
//		if val, ok := token.Claims.(jwtgo.MapClaims)["scopes"].(string); !ok {
//			w.WriteHeader(401)
//			return jwt.ErrJWTError("Token is expired or not valid")
//		} else if val != "api:access" {
//			return jwt.ErrJWTError("you are not allowed to access api")
//		}
//
//		//if val, ok := token.Claims.(jwtgo.MapClaims)["user"].(map[string]interface{}); ok {
//		//fmt.Println(token.Claims.(jwtgo.MapClaims)["user"])
//		//fmt.Println(reflect.TypeOf(token.Claims.(jwtgo.MapClaims)["user"]))
//
//		if val, ok := token.Claims.(jwtgo.MapClaims)["user"].(float64); ok {
//			//fmt.Println("token.Claims[user]:", val)
//			ctx = context.WithValue(ctx, "user.jwt", int(val))
//		}
//		return nextHandler(ctx, w, r)
//	}
//}
