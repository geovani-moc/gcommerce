package actions

import (
	"fmt"
	"net/http"

	"github.com/geovani-moc/gcommerce/controller"
	"github.com/geovani-moc/gcommerce/middleware"

	"github.com/gorilla/mux"
)

//App struct define
type App struct {
	router *mux.Router
	port   string
}

var _app *App

//BuildApp instantiate the _app case not instantiate
func BuildApp() *App {
	if _app == nil {
		_app = NewApp()

		_app.router.HandleFunc("/", controller.Home)
		_app.router.HandleFunc("/product", controller.Product)
		_app.router.HandleFunc("/stock", controller.Stock)

		//API restfull
		_app.router.HandleFunc("/api/product/", middleware.GetAllProducts).Methods("GET", "OPTIONS")

		http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
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
