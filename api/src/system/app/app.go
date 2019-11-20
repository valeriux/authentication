package app

import (
	"log"
	"net/http"

	"github.com/go-xorm/xorm"
	"github.com/gorilla/mux"
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

func HomeHandler(w http.ResponseWriter, r *http.ReadRequest) {
	w.Write([]byte("Hello World"))

}

// Start the server
func (s *Server) Start() {
	log.Println("Starting server on port!" + s.port)

	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)

	http.ListenAndServe(s.port, router)
}
