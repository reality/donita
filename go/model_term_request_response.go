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

type TermRequestResponse struct {

	// URL to to the GitHub issue containing the term request.
	Githubissue string `json:"githubissue,omitempty"`

	// URL to the Wiki page discussing the term request.
	Wikipage string `json:"wikipage,omitempty"`

	// IRI for immediate annotation. If this is a new term request, it will be a provisional identifier for the target ontology. If this is a term inclusion request, it will be the source IRI.
	Identifier *interface{} `json:"identifier,omitempty"`
}