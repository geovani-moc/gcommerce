package actions

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/geovani-moc/gcommerce/controller"
	"github.com/geovani-moc/gcommerce/middleware"

	"github.com/gorilla/mux"
)

//App struct define
type App struct {
	router         *mux.Router
	port           string
	globalTemplate *template.Template
}

var _app *App

//BuildApp instantiate the _app case not instantiate
func BuildApp() *App {
	if _app == nil {
		_app = NewApp()

		_app.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			controller.Home(w, r, _app.globalTemplate)
		})
		_app.router.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
			controller.Product(w, r, _app.globalTemplate)
		})
		_app.router.HandleFunc("/stock", func(w http.ResponseWriter, r *http.Request) {
			controller.Stock(w, r, _app.globalTemplate)
		})

		//API restfull PRODUCT
		_app.router.HandleFunc("/api/product/", middleware.GetAllProducts).Methods("GET", "OPTIONS")
		_app.router.HandleFunc("/api/product{id}", middleware.GetProduct).Methods("GET", "OPTIONS")
		_app.router.HandleFunc("/api/product/{id}", middleware.UpdateProduct).Methods("PUT", "OPTIONS")
		_app.router.HandleFunc("/api/delete-product/{id}", middleware.UpdateProduct).Methods("DELETE", "OPTIONS")

		http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	}
	return _app
}

// NewApp create the app
func NewApp() *App {
	app := &App{
		router:         mux.NewRouter(),
		port:           ":8080",
		globalTemplate: parseTemplates(),
	}
	return app
}

//Run starts the server
func (app *App) Run() error {
	fmt.Println("http://localhost" + app.port)
	http.Handle("/", _app.router)

	return http.ListenAndServe(app.port, nil)
}

// func parseTemplates(directory string) *template.Template {
// 	globalTemplate := template.Must(template.ParseGlob(directory + "/*.html"))
// 	template.Must(globalTemplate.ParseGlob(directory + "/base/*.html"))

// 	return globalTemplate
// }

func getTemplate() *template.Template {
	if nil == _app.globalTemplate {
		_app.globalTemplate = parseTemplates()
	}

	return _app.globalTemplate
}

func parseTemplates() *template.Template {
	templ := template.New("")
	err := filepath.Walk("./view", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = templ.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}

		return err
	})

	if err != nil {
		panic(err)
	}

	return templ
}
