package docker // import "golang.docker.io/go-docker"

import (
	"encoding/json"
	"net/url"

	"golang.docker.io/go-docker/api/types"
	"golang.docker.io/go-docker/api/types/filters"
	"golang.docker.io/go-docker/api/types/swarm"
	"golang.org/x/net/context"
)

// SecretList returns the list of secrets.
func (cli *Client) SecretList(ctx context.Context, options types.SecretListOptions) ([]swarm.Secret, error) {
	query := url.Values{}

	if options.Filters.Len() > 0 {
		filterJSON, err := filters.ToParam(options.Filters)
		if err != nil {
			return nil, err
		}

		query.Set("filters", filterJSON)
	}

	resp, err := cli.get(ctx, "/secrets", query, nil)
	if err != nil {
		return nil, err
	}

	var secrets []swarm.Secret
	err = json.NewDecoder(resp.body).Decode(&secrets)
	ensureReaderClosed(resp)
	return secrets, err
}