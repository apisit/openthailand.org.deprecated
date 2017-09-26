package route

import (
	"net/http"
	"openthailand/middleware"

	"github.com/gorilla/mux"
)

func Use(middlewares ...func(http.HandlerFunc) http.HandlerFunc) []func(http.HandlerFunc) http.HandlerFunc {
	return middlewares
}

type Config struct {
	Path        string
	Func        http.HandlerFunc
	Hook        func(http.HandlerFunc) http.HandlerFunc
	Middlewares []func(http.HandlerFunc) http.HandlerFunc
	Method      string
}

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
	DEL    = "DEL"
)

type MyRoute struct {
	r *mux.Router
}

func Router(r *mux.Router) MyRoute {
	return MyRoute{r}
}

func (m MyRoute) Register(configs ...[]Config) {
	for _, config := range configs {
		for _, route := range config {
			m.r.HandleFunc(route.Path, middleware.UseWithHook(route.Func, route.Hook, route.Middlewares)).Methods(route.Method)
		}
	}
}
