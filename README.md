# Introduction

**This package allows you to quickly and easily use the Stomshield SMC API via Go.**

This package is based on the [OpenAPI specification](https://github.com/trois-six/smc/generator/source/smc-support-3.6-docs-api.zip) of the SMC API, provided by Stormshield.

As you can see in the [generator/Makefile](generator/Makefile), the package is excluding a lot of the tags (firewalls|objects|proxy|rules|rule-sets) of the API right now, as some part of the spec should be revised to be compatible with kin-openapi and oapi-codegen.

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
    client, err := smc.NewSMCClientWithResponses("https://w.x.y.z/papi/v1", "YOUR_API_KEY")
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
# this builds a small go binary that flattens the swagger spec
make build
# this executes the binary
make swagger
```

The specification contains multiple issues, that need to be fixed manually:
1. The `definitions_folders_FolderMember` definition is circular, so it needs to be removed from the spec and replaced by `#/components/schemas/definitions_folders_FolderMember` in the `definitions_folders_FolderMember` definition.
2. In the POST to the path `/api/config/initial/cloud/{cloudName}` the `cloudName` parameter is not defined in the spec, so it needs to be added to the `parameters` section.
```yaml
      parameters:
        - description: Cloud Name
          in: path
          name: cloudName
          required: true
          schema:
            type: string
```
3. In the POST to the path `/api/nsrpc/script`, there is a path parameter `scriptName` that is not defined in the spec, the POST must be completely moved to the `/api/nsrpc/script/{scriptname}` path.

Not resolved yet:
1. The DELETE operations in the paths `/api/network/interfaces/bulk`, `/api/qos/ifaces-assignations/bulk`, `/api/qos/queues/bulk`, `/api/qos/traffic-shapers/bulk`, `/api/rules/{uuid}`, `/api/rules/bulk` contain requestBody. This is not allowed in the OpenAPI 3.0.1 specification and not supported by the generators. This is supported in OpenAPI 3.1.0, but the generators do not support it yet.
2. In the GET to the path `/proxy/{uuid}/admin/{filename}`, two parameters are defined, but the parameters are not set after.

Now, you can generate the client:

```bash
make generate
```

# Documentation

If you need to check the spec, you can use the [Swagger Editor](https://editor.swagger.io/?url=https://raw.githubusercontent.com/trois-six/smc/refs/heads/main/generator/swagger_flat.yaml) with the spec in [swagger_flat.yaml](generator/swagger_flat.yaml).
