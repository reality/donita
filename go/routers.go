/*
 * Provisional Term Management API
 *
 * DRAFT API specification for provisional term request management for an ontology. An instance of the API is attached to a particular ontology, and configured on the back-end to interface with the tooling surrounding that ontology. Therefore, information does not have to be given through the API as to location of Ontology/GitHub/Wiki instances for term addition.
 *
 * API version: 1.0.0
 * Contact: l.slater.1@bham.ac.uk
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v1/",
		Index,
	},

	Route{
		"RequestTermAddition",
		strings.ToUpper("Post"),
		"/v1/requestTermAddition",
		RequestTermAddition,
	},

	Route{
		"RequestTermInclusion",
		strings.ToUpper("Post"),
		"/v1/requestTermInclusion",
		RequestTermInclusion,
	},
}
