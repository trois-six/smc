# Introduction

**This package allows you to quickly and easily use the Stomshield SMC API via Go.**

This package is based on the [OpenAPI specification](https://github.com/trois-six/smc/generator/source/smc-support-3.6-docs-api.zip) of the SMC API, provided by Stormshield.

As you can see in the [generator/Makefile](generator/Makefile), the package is excluding a lot of the tags (cli|config|deployment|firewalls|misc|networkInterfaces|nsrpc|objects|proxy|QoS|routing|rules|rule-sets) of the API right now, as some part of the spec should be revised to be compatible with kin-openapi and oapi-codegen.

# Installation

## Install Package

`go get github.com/trois-six/smc`

## Dependencies

- [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen), this package is basically the generated code from oapi-codegen for the SMC API spec.

# Quick Start

## Hello SMC

The following is the minimum needed code to call the SMS API.

```go
package main

import (
    "context"
    "io"
    "log"
    "net/http"

    "github.com/trois-six/smc"
)

func main() {
    client, err := smc.NewSMCClientWithResponses("https://w.x.y.z/papi/v1", "d22a93be6094d3f7421f37895c0d3646b6741753")
    if err != nil {
        log.Fatal(err)
    }

    resp, err := client.GetAPIAccounts(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    if resp.StatusCode() == http.StatusOK && resp.JSON200 != nil && resp.JSON200.Result != nil {
        for _, result := range *resp.JSON200.Result {
            log.Println(result)
        }
    } else {
        log.Fatalf("Status code: %d", resp.StatusCode())
    }
}
```

## Documentation

The [documentation](https://pkg.go.dev/github.com/trois-six/smc) is available.

## Regenerate the client

If you need to regenerate the client, you can use the following command:

```bash
cd generator
make build
make swagger
# remove the definitions_folders___FolderMember definition from the swagger_flat.yaml file and replace its usage in definitions_folders_FolderMember by '#/components/schemas/definitions_folders_FolderMember', it's a circular reference in the spec
make generate
```

# Documentation

If you need to check the spec, you can use the [Swagger Editor](https://editor.swagger.io/?url=https://raw.githubusercontent.com/trois-six/smc/refs/heads/main/generator/swagger_flat.yaml) with the spec in [swagger_flat.yaml](generator/swagger_flat.yaml).
