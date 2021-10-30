package main

import (
	"bri-rece/api/connect"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type briApp struct {
	connect connect.Connect
	router  *mux.Router
}

func (app *briApp) run() {
	//headersOk := handlers.AllowedHeaders([]string{"*"})
	//originsOk := handlers.AllowedOrigins([]string{"*"})
	//methodsOk := handlers.AllowedMethods([]string{"*"})
	h := app.connect.ApiServer([]string{})
	handlerOk := cors.AllowAll().Handler(app.router)
	log.Println("Listening on", h)
	NewAppRouter(app).InitMainRoutes()
	err := http.ListenAndServe(h, handlerOk)
	if err != nil {
		log.Fatalln(err)
	}

}

func NewBriApp() *briApp {
	r := mux.NewRouter()
	var appConnect = connect.NewConnect()
	return &briApp{
		connect: appConnect,
		router:  r,
	}
}

func main() {
	NewBriApp().run()

}
