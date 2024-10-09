package main

import (
	"bytes"
	"context"
	"log"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
	"gopkg.in/yaml.v3"
)

type FlatSwagger struct {
	OpenAPI    string                         `yaml:"openapi"`
	Info       map[string]interface{}         `yaml:"info"`
	Paths      map[string]interface{}         `yaml:"paths"`
	Components map[string]interface{}         `yaml:"components"`
	Security   []openapi3.SecurityRequirement `yaml:"security"`
}

func main() {
	sl := openapi3.NewLoader()
	sl.IsExternalRefsAllowed = true

	doc, err := sl.LoadFromFile("swagger.yaml")
	if err != nil {
		log.Fatalf("could not load swagger.yaml: %v", err)
	}

	// Resolve external references, to flatten the swagger.
	doc.InternalizeRefs(context.Background(), nil)

	// Add the security scheme to the flat swagger.
	doc.Components.SecuritySchemes = make(map[string]*openapi3.SecuritySchemeRef)
	doc.Components.SecuritySchemes["apiKey"] = &openapi3.SecuritySchemeRef{
		Value: &openapi3.SecurityScheme{
			Type:   "http",
			Scheme: "bearer",
		},
	}

	// Add the security requirement to the flat swagger.
	doc.Security = *openapi3.NewSecurityRequirements().With(openapi3.SecurityRequirement{
		"apiKey": []string{},
	})

	bytesOpenAPIFlatSwagger, err := yaml.Marshal(doc)
	if err != nil {
		log.Fatalf("could not marshal to YAML the flat swagger: %v", err)
	}

	// Due to a bug in a kin-openapi dependency, the order of the OpenAPI spec is not preserved
	// So we are doing an unmasrhal and a marshal again to fix it.
	// Else we would have written bytesOpenAPIFlatSwagger directly to the file.

	// err = os.WriteFile("swagger_flat.yaml", bytesOpenAPIFlatSwagger, 0644)
	// if err != nil {
	// 	log.Fatalf("could not write modified swagger data to file: %v", err)
	// }

	var flatSwagger FlatSwagger
	err = yaml.Unmarshal(bytesOpenAPIFlatSwagger, &flatSwagger)
	if err != nil {
		log.Fatalf("could not unmarshal to YAML the flat swagger: %v", err)
	}

	// In yaml.v3 the default indent is 4, we need to set it to 2 when we marshal the data.
	var b bytes.Buffer
	e := yaml.NewEncoder(&b)
	e.SetIndent(2)
	if err = e.Encode(flatSwagger); err != nil {
		log.Fatalf("could not marshal to YAML the flat swagger: %v", err)
	}

	err = os.WriteFile("swagger_flat.yaml", b.Bytes(), 0644)
	if err != nil {
		log.Fatalf("could not write modified swagger data to file: %v", err)
	}
}
