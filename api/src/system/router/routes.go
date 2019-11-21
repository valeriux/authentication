package router

// Handlers over routes
import (
	"authentication/api/pkg/types/routes"
	HomeHandler "authentication/api/src/controllers/home"
	"log"
	"net/http"

	"github.com/go-xorm/xorm"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Global middleware reached!!!")
		next.ServeHTTP(w, r)
	})
}

func GetRoutes(db *xorm.Engine) routes.Routes {

	HomeHandler.Init(db)
	return routes.Routes{
		routes.Route{"Home", "GET", "/", HomeHandler.Index},
	}
}
