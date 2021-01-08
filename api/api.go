// Routes requests to the correct API functions.
package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go/doc"
	"go/parser"
	"go/token"
	"net/http"
	"reflect"
	"runtime"
	"strings"
	"sync"
)

// Map containing all available API functions. Functions are added to this map via init() functions in their declaration files.
var routedFuncs = make(map[string]http.HandlerFunc)

func init() {
	routedFuncs[""] = ListAvailable
}

// Create a router that dynamically routes all API function paths.
func Subrouter(r *mux.Router) {
	sub := r.PathPrefix("/api").Subrouter()

	for k, v := range routedFuncs {
		sub.Path(fmt.Sprintf("/%s", k)).HandlerFunc(v)
	}
	sub.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("ERROR: API function does not exist"))
	})
}

// A struct used in creating the JSON response for the ListAvailable API function.
type jsonFunc struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Description string `json:"description"`
}

// List the available API functions based on the routedFuncs map.
func ListAvailable(w http.ResponseWriter, _ *http.Request) {
	// Get the descriptions of the functions from the documentation comments.
	var (
		funcs    []jsonFunc
		funcsMtx sync.RWMutex
		fset     = token.NewFileSet()
	)

	if d, e := parser.ParseDir(fset, "./api/", nil, parser.ParseComments); e != nil {
		http.Error(w, fmt.Sprintf("ERROR: got error '%e' while parsing", e), http.StatusInternalServerError)
		return
	} else {
		if v, ok := d["api"]; !ok {
			http.Error(w, "ERROR: failed to access function descriptions", http.StatusInternalServerError)
			return
		} else {
			var (
				pkg    = doc.New(v, "./api/", doc.AllMethods)
				wg     sync.WaitGroup
				pkgMtx sync.RWMutex
			)

			for k, v := range routedFuncs {
				name := strings.TrimPrefix(runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name(), "github.com/noahfriedman-ca/server/api.")

				wg.Add(1)
				go func(name, k string) {
					defer wg.Done()
					desc := "*No description found.*"

					for i := 0; i < len(pkg.Funcs); i++ {
						pkgMtx.RLock()
						f := *pkg.Funcs[i]
						pkgMtx.RUnlock()

						if f.Name == name {
							desc = strings.Trim(f.Doc, "\n")
							break
						}
					}

					funcsMtx.Lock()
					funcs = append(funcs, jsonFunc{
						Name:        name,
						Path:        "/api/" + k,
						Description: desc,
					})
					funcsMtx.Unlock()
				}(name, k)
			}

			wg.Wait()
			if b, e := json.MarshalIndent(struct {
				Funcs []jsonFunc `json:"functions"`
			}{funcs}, "", "    "); e != nil {
				http.Error(w, "ERROR: failed to generate JSON", http.StatusInternalServerError)
				return
			} else {
				_, _ = w.Write(b)
			}
		}
	}
}
