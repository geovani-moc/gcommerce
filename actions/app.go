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
	"github.com/geovani-moc/gcommerce/i18n"
	"github.com/geovani-moc/gcommerce/middleware"
	"github.com/geovani-moc/gcommerce/util"

	"github.com/gorilla/mux"
)

//App struct define
type App struct {
	router          *mux.Router
	version         string
	databaseVersion float64
	root            util.Root
}

var _app *App

//BuildApp instantiate the _app case not instantiate
func BuildApp() *App {
	if nil == _app {
		_app = NewApp()

		//pages
		_app.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			controller.Home(w, r, &_app.root)
		})
		_app.router.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
			controller.Home(w, r, &_app.root)
		})
		_app.router.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
			controller.Product(w, r, &_app.root)
		})
		_app.router.HandleFunc("/stock", func(w http.ResponseWriter, r *http.Request) {
			controller.Stock(w, r, &_app.root)
		})
		_app.router.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
			controller.Profile(w, r, &_app.root)
		})

		//API restfull PRODUCT
		_app.router.HandleFunc("/api/product/", middleware.GetAllProducts).Methods("GET", "OPTIONS")
		_app.router.HandleFunc("/api/product{id}", middleware.GetProduct).Methods("GET", "OPTIONS")
		_app.router.HandleFunc("/api/product/{id}", middleware.UpdateProduct).Methods("PUT", "OPTIONS")
		_app.router.HandleFunc("/api/delete-product/{id}", middleware.UpdateProduct).Methods("DELETE", "OPTIONS")
		http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

		//static files
		_app.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
			cacheControlWrapper(http.FileServer(http.Dir("static"))))) //static files with cache control

		//not found page
		_app.router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			controller.Code404(w, r, &_app.root)
		})
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
	log.Print("Criando App ...")
	app := &App{
		router:          mux.NewRouter(),
		version:         "0.1 alpha",
		databaseVersion: 0.1, //ler e gravar no arquivo env, e comparar com o banco de dados
		root: util.Root{
			Port:            ":8080",
			Templates:       parseTemplates(),
			CurrentLanguage: "pt-BR",
			NamePages: []string{
				"home",
				"product",
				"stock",
				"profile",
			},
		},
	}
	var err error
	app.root.Dictionaries, err = i18n.GetAllDictionaries()

	if nil != err {
		log.Print("Nehum idioma localizado")
	}

	log.Print("App criado ...")
	return app
}

//Run starts the server
func (app *App) Run() error {
	log.Print("App iniciado ...")
	fmt.Println("Version: Gcommerce ", app.version)
	fmt.Println("Aplicação disponivel em: http://localhost" + app.root.Port)

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
