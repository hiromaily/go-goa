package middleware

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/hiromaily/go-goa/pkg/jwts"
)

// Middlewarer interface
type Middlewarer interface {
	CheckJWT() func(http.Handler) http.Handler
}

// middleware object
type middleware struct {
	logger *zap.Logger
	jwter  jwts.JWTer
}

// NewMiddleware returns Server interface
func NewMiddleware(
	logger *zap.Logger,
	jwter jwts.JWTer,
) Middlewarer {
	return &middleware{
		logger: logger,
		jwter:  jwter,
	}
}

// FIXME: This code is outdated
func (m *middleware) CheckJWT() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		// A HTTP handler is a function.
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//token := jwt.ContextJWT(ctx)
			//
			//if val, ok := token.Claims.(jwtgo.MapClaims)["scopes"].(string); !ok {
			//	w.WriteHeader(401)
			//	return jwt.ErrJWTError("Token is expired or not valid")
			//} else if val != "api:access" {
			//	return jwt.ErrJWTError("you are not allowed to access api")
			//}
			//if val, ok := token.Claims.(jwtgo.MapClaims)["user"].(float64); ok {
			//	//fmt.Println("token.Claims[user]:", val)
			//	ctx = context.WithValue(ctx, "user.jwt", int(val))
			//}

			m.logger.Info("middleware CheckJWT")

			auth := r.Header.Get("Authorization")
			if auth == "" {
				m.logger.Error("http.StatusBadRequest", zap.Error(errors.New("Authorization header is missing")))
				return
			}

			// Bearer token
			authParts := strings.Split(auth, " ")
			if len(authParts) != 2 || authParts[0] != "Bearer" {
				m.logger.Error("http.StatusBadRequest", zap.Error(errors.New("Authorization header is invalid")))
				return
			}
			if err := m.jwter.ValidateToken(authParts[1]); err != nil {
				m.logger.Error("http.StatusBadRequest", zap.Error(err))
				return
			}

			h.ServeHTTP(w, r)
		})
	}
}
