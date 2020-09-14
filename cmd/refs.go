package cmd

import (
	"strings"

	"github.com/go-openapi/spec"
)

type refs map[string]struct{}

func newRefs() refs {
	return refs{}
}

func findRefs(specif *spec.Swagger, s string, operation *spec.Operation, refs refs) {
	if operation == nil {
		return
	}

	schemas := make([]*spec.Schema, 0, len(operation.Responses.ResponsesProps.StatusCodeResponses))

	for _, v := range operation.Responses.ResponsesProps.StatusCodeResponses {
		schemas = append(schemas, v.ResponseProps.Schema)
	}

	for _, v := range operation.Parameters {
		schemas = append(schemas, v.ParamProps.Schema)
	}

	for _, schema := range schemas {
		if schema == nil {
			continue
		}
		findRefsInSchema(specif, schema, refs)
	}

}

func findRefsInSchema(specif *spec.Swagger, schema *spec.Schema, refs refs) {

	// Recurse in Properties
	for _, property := range schema.SchemaProps.Properties {
		findRefsInSchema(specif, &property, refs)
	}

	// Follow Ref
	url := schema.Ref.GetURL()
	if url != nil {
		fragment := url.Fragment
		prefix := "/definitions/"
		if !strings.HasPrefix(fragment, "/definitions/") {
			panic("no prefix " + prefix)
		}
		defkey := strings.Trim(url.Fragment[len(prefix):], "\"")
		refs[defkey] = struct{}{}

		def := specif.SwaggerProps.Definitions[defkey]
		findRefsInSchema(specif, &def, refs)
	}

	items := schema.SchemaProps.Items
	if items != nil {
		findRefsInSchema(specif, items.Schema, refs)
	}
}
