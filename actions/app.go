package actions

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//App struct define
type App struct {
	router *mux.Router
	port   string
}

var _app *App

//StartApp instantiate the _app case not instantiate
func StartApp() *App {
	if _app == nil {
		_app = NewApp()

		staticDir := "/static/"
		var dir string

		flag.StringVar(&dir, "dir", ".", "http://localhost"+_app.port)

		_app.router.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir(dir))))

	}
	return _app
}

// NewApp create the app
func NewApp() *App {
	app := &App{
		router: mux.NewRouter(),
		port:   ":8080",
	}
	return app
}

//Run starts the server
func (app *App) Run() error {
	fmt.Println("http://localhost" + app.port)
	http.Handle("/", _app.router)

	return http.ListenAndServe(app.port, nil)
}
