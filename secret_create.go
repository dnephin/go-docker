package docker // import "golang.docker.com/go-docker"

import (
	"encoding/json"

	"golang.docker.com/go-docker/api/types"
	"golang.docker.com/go-docker/api/types/swarm"
	"golang.org/x/net/context"
)

// SecretCreate creates a new Secret.
func (cli *Client) SecretCreate(ctx context.Context, secret swarm.SecretSpec) (types.SecretCreateResponse, error) {
	var response types.SecretCreateResponse
	if err := cli.NewVersionError("1.25", "secret create"); err != nil {
		return response, err
	}
	resp, err := cli.post(ctx, "/secrets/create", nil, secret, nil)
	if err != nil {
		return response, err
	}

	err = json.NewDecoder(resp.body).Decode(&response)
	ensureReaderClosed(resp)
	return response, err
}
