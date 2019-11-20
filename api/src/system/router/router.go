package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Worldddd!!!"))


type Router struct {
	Router *mux.Router
}


func (r *Router) Init() {
	r.Router.HandleFunc("/", HomeHandler)

}

func NewRouter() (r Router) {
	r.Router = mux.NewRouter()

	return
}