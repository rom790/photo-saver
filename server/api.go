package server

import (
	"fmt"
	"net/http"
	"photo-saver/store"

	"github.com/gorilla/mux"
)

type APIserver struct {
	config *Config
	router *mux.Router
	store  *store.DataStore
}

func NewServer(config *Config) *APIserver {
	return &APIserver{
		config: config,
		router: mux.NewRouter(),
	}
}

func (api *APIserver) Start() error {
	fmt.Println("Starting server on ", api.config.Addr)

	api.configureRouter()

	err := api.configureStore()
	if err != nil {
		return fmt.Errorf("db config error %w", err)
	}

	server := &http.Server{
		Addr:    api.config.Addr,
		Handler: api.router,
	}

	return server.ListenAndServe()
}

func (api *APIserver) configureRouter() {

}

func (api *APIserver) configureStore() error {
	st := store.NewDataStore(api.config.Store)
	// fmt.Println(api.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	api.store = &st

	return nil
}
