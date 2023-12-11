package tokenexchange

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/dghubble/sling"
)

// TokenExchangeService handles exchange of OpenID Connect ID tokens for Octopus access tokens
type TokenExchangeService struct {
	services.Service
}

type OpenIdConfigurationResponse struct {
	Issuer        string `json:"issuer"`
	TokenEndpoint string `json:"token_endpoint"`
}

type TokenExchangeRequest struct {
	GrantType        string `json:"grant_type"`
	Audience         string `json:"audience"`
	SubjectTokenType string `json:"subject_token_type"`
	SubjectToken     string `json:"subject_token"`
}

type TokenExchangeResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int32  `json:"expires_in"`
}

type TokenExchangeErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// NewAPIKeyService returns an TokenExchangeService with a preconfigured client.
func NewTokenExchangeService(sling *sling.Sling) *TokenExchangeService {
	return &TokenExchangeService{
		Service: services.NewService("TokenExchangeService", sling, ""),
	}
}

// ExchangeOpenIdConnectIdTokenForAccessToken exchanges an OpenID Connect ID token for an Octopus access token
func (s *TokenExchangeService) ExchangeOpenIdConnectIdTokenForAccessToken(host string, serviceAccountId string, idToken string) (*TokenExchangeResponse, error) {
	if internal.IsEmpty(serviceAccountId) {
		return nil, internal.CreateInvalidParameterError("ExchangeOpenIdConnectIdTokenForAccessToken", "serviceAccountId")
	}

	if internal.IsEmpty(idToken) {
		return nil, internal.CreateInvalidParameterError("ExchangeOpenIdConnectIdTokenForAccessToken", "idToken")
	}

	if err := services.ValidateInternalState(s); err != nil {
		return nil, err
	}

	base := s.Sling.Base(host)
	resp, err := api.ApiGet(base, new(OpenIdConfigurationResponse), "/.well-known/openid-configuration")

	if err != nil {
		return nil, err
	}

	openIdConfigurationResponse := resp.(*OpenIdConfigurationResponse)

	tokenExchangeData := TokenExchangeRequest{
		GrantType:        "urn:ietf:params:oauth:grant-type:token-exchange",
		Audience:         serviceAccountId,
		SubjectTokenType: "urn:ietf:params:oauth:token-type:jwt",
		SubjectToken:     idToken,
	}

	resp, err = services.ApiPost(base, tokenExchangeData, new(TokenExchangeResponse), openIdConfigurationResponse.TokenEndpoint)

	if err != nil {
		return nil, err
	}

	tokenExchangeResponse := resp.(*TokenExchangeResponse)

	return tokenExchangeResponse, nil
}
