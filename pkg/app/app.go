package app

import (
	"github.com/gorilla/mux"
	"github.com/kosha/pinecone-connector/pkg/config"
	"github.com/kosha/pinecone-connector/pkg/logger"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	Log    logger.Logger
	Cfg    *config.Config
}

func router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	return router
}

// Initialize creates the necessary scaffolding of the app
func (a *App) Initialize(log logger.Logger) {

	cfg := config.Get()

	a.Cfg = cfg
	a.Log = log

	a.Router = router()
}

// Run starts the app and serves on the specified addr
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
