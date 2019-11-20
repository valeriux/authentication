package app

import (
	"authentication/api/src/system/router"
	"log"
	"net/http"

	"github.com/go-xorm/xorm"
)

type Server struct {
	port string
	Db   *xorm.Engine
}

func NewServer() Server {
	return Server{}
}

// Init all vals
func (s *Server) Init(port string, db *xorm.Engine) {
	log.Println("Initializing server...")
	s.port = ":" + port
	s.Db = db
}

// Start the server
func (s *Server) Start() {
	log.Println("Starting server on port!" + s.port)

	r.Init()
	r := router.NewRouter()

	http.ListenAndServe(s.port, r.Router)
}
