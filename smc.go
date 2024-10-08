package smc

import (
	"context"
	"net/http"

	"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
)

func NewSMCClient(hostname, apiKey string) (*Client, error) {
	apiKeyAuthorization, err := securityprovider.NewSecurityProviderBearerToken(apiKey)
	if err != nil {
		return nil, err
	}

	return NewClient(
		hostname,
		WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Accept", "application/json")
			return nil
		}),
		WithRequestEditorFn(apiKeyAuthorization.Intercept),
		// WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		// 	dump, err := httputil.DumpRequestOut(req, true)

		// 	if err != nil {
		// 		return err
		// 	}

		// 	log.Printf("\n\n%q", dump)

		// 	return nil
		// }),
	)
}

func NewSMCClientWithResponses(hostname, apiKey string) (*ClientWithResponses, error) {
	apiKeyAuthorization, err := securityprovider.NewSecurityProviderBearerToken(apiKey)
	if err != nil {
		return nil, err
	}

	return NewClientWithResponses(
		hostname,
		WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Accept", "application/json")
			return nil
		}),
		WithRequestEditorFn(apiKeyAuthorization.Intercept),
		// WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		// 	dump, err := httputil.DumpRequestOut(req, true)

		// 	if err != nil {
		// 		return err
		// 	}

		// 	log.Printf("\n\n%q", dump)

		// 	return nil
		// }),
	)
}
