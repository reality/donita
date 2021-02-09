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
    "log"
    "net/http"
    "time"
)

func Logger(inner http.Handler, name string) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        inner.ServeHTTP(w, r)

        log.Printf(
            "%s %s %s %s",
            r.Method,
            r.RequestURI,
            name,
            time.Since(start),
        )
    })
}
