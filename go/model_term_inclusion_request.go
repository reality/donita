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

type TermInclusionRequest struct {

	// The IRI of the external ontology term to import.
	SourceIri string `json:"sourceIri"`

	// The ID of the external ontology from which to import the class, as it appears in BioPortal.
	SourceOntology string `json:"sourceOntology"`

	// (optional) The IRI of the suggested direct parent class for the requested term.
	ParentClass string `json:"parentClass,omitempty"`

	// Should the direct parent class of the external ontology term also be included?
	IncludeDirectParent bool `json:"includeDirectParent,omitempty"`

	// Name of the person requesting the term addition.
	CreatorName string `json:"creatorName"`

	// Email address of the person requesting the term addition.
	CreatorEmail string `json:"creatorEmail"`
}