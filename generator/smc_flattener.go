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
	OpenAPI    string                 `yaml:"openapi"`
	Info       map[string]interface{} `yaml:"info"`
	Paths      map[string]interface{} `yaml:"paths"`
	Components map[string]interface{} `yaml:"components"`
}

func main() {
	sl := openapi3.NewLoader()
	sl.IsExternalRefsAllowed = true

	doc, err := sl.LoadFromFile("swagger.yaml")
	if err != nil {
		log.Fatalf("could not load swagger.yaml: %v", err)
	}

	doc.InternalizeRefs(context.Background(), nil)
	// if schema, found := doc.Components.Schemas["definitions_folders___FolderMember"]; found {
	// 	// log.Printf("found schema: %#v", schema.Value)
	// 	// schema.Value.Extensions = make(map[string]any)
	// 	// schema.Value.Extensions["x-go-name"] = "DefinitionsFoldersFolderMemberParentParentParent"
	// 	// doc.Components.Schemas["definitions_folders___FolderMember"] = schema
	// 	log.Printf("found schema: %#v", schema.Value)
	// }

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
