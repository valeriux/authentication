package router

import (
	"log"
	"net/http"

	"authentication/api/pkg/types/routes"
	UsersHandler "authentication/api/src/controllers/v1/routes/users"
	StatusHandler "authentication/api/src/controllers/v1/status"

	"github.com/go-xorm/xorm"
)

var db *xorm.Engine

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-App-Token")
		if len(token) < 1 {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		log.Println("Inside V1 Middleware")

		next.ServeHTTP(w, r)
	})
}

func GetRoutes(DB *xorm.Engine) (SubRoute map[string]routes.SubRoutePackage) {
	db = DB

	StatusHandler.Init(DB)
	UsersHandler.Init(DB)

	/* ROUTES */
	SubRoute = map[string]routes.SubRoutePackage{
		"/v1": {
			Routes: routes.Routes{
				routes.Route{"Status", "GET", "/status", StatusHandler.Index},
				// routes.Route{"UsersIndex", "GET", "/users", UsersHandler.Index},
				// routes.Route{"UsersStore", "POST", "/users", UsersHandler.Store},
				// routes.Route{"UsersEdit", "GET", "/users/{id}/edit", UsersHandler.Edit},
				// routes.Route{"UsersUpdate", "PUT", "/users/{id}", UsersHandler.Update},
				// routes.Route{"UsersDestroy", "DELETE", "/users/{id}", UsersHandler.Destroy},
			},
			Middleware: Middleware,
		},
	}

	return
}
