package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	v1 "kratosTestApp/api/user/v1"
	"kratosTestApp/internal/conf"
	"kratosTestApp/internal/pkg/middlewire"
	"kratosTestApp/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwt5 "github.com/golang-jwt/jwt/v5"
)

func NewWhiteListMatcher() selector.MatchFunc {

	whiteList := make(map[string]struct{})
	whiteList["/user.v1.User/Register"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true //false：所有的都设置为白名单
	}
}

func NewHTTPServer(c *conf.Server, ac *conf.Auth, user *service.UserService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			selector.Server( //用户鉴权
				jwt.Server(func(token *jwt5.Token) (interface{}, error) {
					return []byte(ac.Jwt.Secret), nil
				}, jwt.WithSigningMethod(jwt5.SigningMethodHS256),
					jwt.WithClaims(func() jwt5.Claims {
						return &jwt5.MapClaims{}
					})),
			).Match(NewWhiteListMatcher()).Build(),
			validate.Validator(), //参数校验
		),
	}

	opts = append(opts, http.ResponseEncoder(middlewire.CustomResponseEncoder))

	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterUserHTTPServer(srv, user)
	return srv
}
