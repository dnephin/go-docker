package docker // import "golang.docker.com/go-docker"

import (
	"encoding/json"
	"net/url"

	"golang.docker.com/go-docker/api/types"
	"golang.docker.com/go-docker/api/types/filters"
	"golang.docker.com/go-docker/api/types/swarm"
	"golang.org/x/net/context"
)

// NodeList returns the list of nodes.
func (cli *Client) NodeList(ctx context.Context, options types.NodeListOptions) ([]swarm.Node, error) {
	query := url.Values{}

	if options.Filters.Len() > 0 {
		filterJSON, err := filters.ToParam(options.Filters)

		if err != nil {
			return nil, err
		}

		query.Set("filters", filterJSON)
	}

	resp, err := cli.get(ctx, "/nodes", query, nil)
	if err != nil {
		return nil, err
	}

	var nodes []swarm.Node
	err = json.NewDecoder(resp.body).Decode(&nodes)
	ensureReaderClosed(resp)
	return nodes, err
}
