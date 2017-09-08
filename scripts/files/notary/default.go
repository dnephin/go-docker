package notary

import (
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/docker/go-connections/sockets"
	"github.com/docker/notary/trustpinning"
	"github.com/docker/notary/tuf/data"
	"golang.docker.com/go-docker/api/types"
	"golang.docker.com/go-docker/registry/auth"
	"golang.docker.com/go-docker/registry/auth/challenge"
	"golang.docker.com/go-docker/registry/transport"
)

const authClientID = "docker"

func makeHubTransport(server, image string, authConfig *types.AuthConfig) http.RoundTripper {
	base := http.DefaultTransport
	modifiers := []transport.RequestModifier{
		transport.NewHeaderRequestModifier(http.Header{
			"User-Agent": []string{"go-docker-v1"},
		}),
	}

	authTransport := transport.NewTransport(base, modifiers...)
	pingClient := &http.Client{
		Transport: authTransport,
		Timeout:   5 * time.Second,
	}
	req, err := http.NewRequest("GET", server+"/v2/", nil)
	if err != nil {
		panic(err)
	}

	challengeManager := challenge.NewSimpleManager()
	resp, err := pingClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if err := challengeManager.AddResponse(resp); err != nil {
		panic(err)
	}
	tokenHandler := auth.NewTokenHandler(base, nil, image, "pull")
	modifiers = append(modifiers, auth.NewAuthorizer(challengeManager, tokenHandler, auth.NewBasicHandler(nil)))

	return transport.NewTransport(base, modifiers...)
}

// NewRepositoryWithDefaults is the simplest way to instantiate a notary Repository.
// Useful when talking to server "https://notary.docker.io".
// Note that image repository has to be fully qualified, for example: "docker.io/library/alpine"
func NewRepositoryWithDefaults(server, image string) (Repository, error) {
	rootDir := ".trust"
	return NewFileCachedRepository(rootDir, data.GUN(image), server, makeHubTransport(server, image), nil, trustpinning.TrustPinConfig{})
}

func defaultTransport(server, image string, authConfig *types.AuthConfig) (http.RoundTripper, error) {
	direct := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}

	// TODO(dmcgowan): Call close idle connections when complete, use keep alive
	base := &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		Dial:                direct.Dial,
		TLSHandshakeTimeout: 10 * time.Second,
		TLSClientConfig:     nil,
		// TODO(dmcgowan): Call close idle connections when complete and use keep alive
		DisableKeepAlives: true,
	}

	proxyDialer, err := sockets.DialerFromEnvironment(direct)
	if err == nil {
		base.Dial = proxyDialer.Dial
	}

	modifiers := []transport.RequestModifier{
		transport.NewHeaderRequestModifier(http.Header{
			"User-Agent": []string{"go-docker-v1"},
		}),
	}
	authTransport := transport.NewTransport(base, modifiers...)

	pingClient := &http.Client{
		Transport: authTransport,
		Timeout:   5 * time.Second,
	}
	req, err := http.NewRequest("GET", server+"/v2/", nil)
	if err != nil {
		return nil, err
	}
	challengeManager := challenge.NewSimpleManager()
	resp, err := pingClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := challengeManager.AddResponse(resp); err != nil {
		return nil, err
	}
	//tokenHandler := auth.NewTokenHandler(base, nil, image, "pull")

	tokenHandlerOptions := auth.TokenHandlerOptions{
		Transport: authTransport,
		Scopes:    []auth.Scope{"pull"},
		ClientID:  authClientID,
	}

	if authConfig != nil {
		tokenHandlerOptions.Scopes = append(tokenHandlerOptions.Scopes, "push")
		tokenHandlerOptions.Credentials = staticCredentialStore{authConfig}
	}

	tokenHandler := auth.NewTokenHandlerWithOptions(tokenHandlerOptions)
	modifiers = append(modifiers, auth.NewAuthorizer(challengeManager, tokenHandler, auth.NewBasicHandler(nil)))

	/*
		if authConfig.RegistryToken != "" {
			passThruTokenHandler := &existingTokenHandler{token: authConfig.RegistryToken}
			modifiers = append(modifiers, auth.NewAuthorizer(challengeManager, passThruTokenHandler))
		} else {
	*/
	scope := auth.RepositoryScope{
		Repository: repoName,
		Actions:    actions,
		Class:      repoInfo.Class,
	}

	basicHandler := auth.NewBasicHandler(creds)
	modifiers = append(modifiers, auth.NewAuthorizer(challengeManager, tokenHandler, basicHandler))
	//}
	return transport.NewTransport(base, modifiers...), nil
}

type staticCredentialStore struct {
	auth *types.AuthConfig
}

func (scs staticCredentialStore) Basic(*url.URL) (string, string) {
	if scs.auth == nil {
		return "", ""
	}
	return scs.auth.Username, scs.auth.Password
}

func (scs staticCredentialStore) RefreshToken(*url.URL, string) string {
	if scs.auth == nil {
		return ""
	}
	return scs.auth.IdentityToken
}

func (scs staticCredentialStore) SetRefreshToken(*url.URL, string, string) {
}
