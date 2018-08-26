package main

import (
	"flag"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

var routes = flag.Bool("routes", false, "Generate router documentation")

func main() {
	flag.Parse()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})

	// RESTy routes for "articles" resource
	r.Route("/access_keys", func(r chi.Router) {
		r.Get("/", ListAccessKeys)
	})

	http.ListenAndServe(":3000", r)
}

// AccessKeys json config
type AccessKeys struct {
	Server string `json:"server"`
	Port   int64  `json:"port"`
	Key    string `json:"key"`
}

// Render is interface of render
func (ak *AccessKeys) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// ListAccessKeys middleware is used to load an Article object from
func ListAccessKeys(w http.ResponseWriter, r *http.Request) {
	println("list keys =====")
	println("add ci test ====")
	list := []render.Renderer{}
	key1 := &AccessKeys{Server: "192.168.0.1", Port: 1080, Key: "acb4"}
	list = append(list, key1)
	render.RenderList(w, r, list)
}
