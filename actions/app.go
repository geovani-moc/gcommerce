package actions

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/geovani-moc/gcommerce/i18n"
	"github.com/geovani-moc/gcommerce/util"

	"github.com/gorilla/mux"
)

//App struct define
type App struct {
	router          *mux.Router
	version         string
	root            util.Root
	databaseVersion string
}

var _app *App

//BuildApp instantiate the _app case not instantiate
func BuildApp() *App {
	if nil == _app {
		_app = NewApp()
		routerURL(_app)
	}
	return _app
}

// controle do tempo do cache
func cacheControlWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "max-age=60") // em segundos
		h.ServeHTTP(w, r)
	})
}

// NewApp create the app
func NewApp() *App {
	util.LoadEnvVariables()

	app := &App{
		router:          mux.NewRouter(),
		version:         "0.3",
		databaseVersion: os.Getenv("DATABASE_VERSION"),
		root: util.Root{
			Port:            ":8080",
			Templates:       parseTemplates(),
			CurrentLanguage: "pt-BR",
			NamePages: []string{
				"home",
				"product",
				"stock",
				"profile",
				"login",
			},
			URL: "http://localhost:8080",
		},
	}
	var err error
	app.root.Dictionaries, err = i18n.GetAllDictionaries()

	if nil != err {
		log.Print("Nehum idioma localizado")
	}
	return app
}

//Run starts the server
func (app *App) Run() error {
	fmt.Println("Gcommerce", app.version)
	fmt.Println("Database", app.databaseVersion)
	fmt.Println("Link: http://localhost" + app.root.Port)

	return http.ListenAndServe(app.root.Port, app.router)
}

func parseTemplates() *template.Template {
	templ := template.New("")
	err := filepath.Walk("./view", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = templ.ParseFiles(path)
			if err != nil {
				log.Print("erro ao realizar o parser do template, ", err)
			}
		}
		return err
	})

	if err != nil {
		log.Print("erro na verificação de diretorios do template, ", err)
	}

	return templ
}
