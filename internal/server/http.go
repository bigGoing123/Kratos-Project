package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	v1 "kratosTestApp/api/user/v1"
	"kratosTestApp/internal/conf"
	"kratosTestApp/internal/server/self_middle"
	"kratosTestApp/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewWhiteListMatcher() selector.MatchFunc {

	whiteList := make(map[string]struct{})
	whiteList["/*"] = struct{}{}
	whiteList["/shop.interface.v1.ShopInterface/Register"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, ac *conf.Auth, user *service.UserService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			//selector.Server(
			//	jwt.Server(func(token *jwt5.Token) (interface{}, error) {
			//		return []byte(ac.ApiKey), nil
			//	}, jwt.WithSigningMethod(jwt5.SigningMethodHS256), jwt.WithClaims(func() jwt5.Claims {
			//		return &jwt5.MapClaims{}
			//	})),
			//).Match(NewWhiteListMatcher()).Build(),
		),
	}
	// 添加自定义的 ResponseEncoder
	opts = append(opts, http.ResponseEncoder(self_middle.CustomResponseEncoder))

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
