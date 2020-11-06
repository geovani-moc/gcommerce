package actions

import (
	"net/http"

	"github.com/geovani-moc/gcommerce/controller"
	"github.com/geovani-moc/gcommerce/middleware"
)

func routerURL(app *App) {
	//pages
	app.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controller.Home(w, r, &app.root)
	})
	app.router.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		controller.Home(w, r, &app.root)
	})
	app.router.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		controller.Product(w, r, &app.root)
	})
	app.router.HandleFunc("/stock", func(w http.ResponseWriter, r *http.Request) {
		controller.Stock(w, r, &app.root)
	})
	app.router.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		controller.Profile(w, r, &app.root)
	})
	app.router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		controller.Login(w, r, &app.root)
	})

	//API restfull PRODUCT
	app.router.HandleFunc("/api/product/", middleware.GetAllProducts).Methods("GET", "OPTIONS")
	app.router.HandleFunc("/api/product{id}", middleware.GetProduct).Methods("GET", "OPTIONS")
	app.router.HandleFunc("/api/product/{id}", middleware.UpdateProduct).Methods("PUT", "OPTIONS")
	app.router.HandleFunc("/api/delete-product/{id}", middleware.UpdateProduct).Methods("DELETE", "OPTIONS")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	//static files
	app.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		cacheControlWrapper(http.FileServer(http.Dir("static"))))) //static files with cache control

	//not found page
	app.router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controller.Code404(w, r, &app.root)
	})
}
