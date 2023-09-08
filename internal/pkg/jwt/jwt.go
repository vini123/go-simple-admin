package jwt

import (
	"context"
	"sync"
	"time"

	jwtv2 "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/frame/g"
)

type JwtConfig struct {
	Realm      string
	Key        string
	Timeout    int
	MaxRefresh int
}

var authService *jwtv2.GfJWTMiddleware

var once sync.Once

func NewJwt() *jwtv2.GfJWTMiddleware {
	once.Do(func() {

		config := JwtConfig{
			Realm:      "go simple admin",
			Key:        "this is key",
			Timeout:    60 * 24 * 7,
			MaxRefresh: 60 * 24 * 10,
		}

		auth := jwtv2.New(&jwtv2.GfJWTMiddleware{
			Realm:           config.Realm,
			Key:             []byte(config.Key),
			Timeout:         time.Minute * time.Duration(config.Timeout),
			MaxRefresh:      time.Minute * time.Duration(config.MaxRefresh),
			IdentityKey:     "id",
			TokenLookup:     "header: Authorization, query: token, cookie: jwt",
			TokenHeadName:   "Bearer",
			TimeFunc:        time.Now,
			Authenticator:   Authenticator,
			Unauthorized:    Unauthorized,
			PayloadFunc:     PayloadFunc,
			IdentityHandler: IdentityHandler,
		})
		authService = auth
	})
	return authService
}

func PayloadFunc(data interface{}) jwtv2.MapClaims {
	claims := jwtv2.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

func IdentityHandler(ctx context.Context) interface{} {
	claims := jwtv2.ExtractClaims(ctx)
	return claims[authService.IdentityKey]
}

func Unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	r.Response.WriteJson(g.Map{
		"code":    code,
		"message": message,
	})
	r.ExitAll()
}

// Authenticator is used to validate login parameters.
// It must return user data as user identifier, it will be stored in Claim Array.
// if your identityKey is 'id', your user data must have 'id'
// Check error (e) to determine the appropriate error message.
func Authenticator(ctx context.Context) (user interface{}, err error) {
	return user, nil
}
