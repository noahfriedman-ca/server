// Routes requests to the correct API functions.
package api

import "github.com/gorilla/mux"

func Router() *mux.Router {
	return mux.NewRouter()
}
