package docker // import "golang.docker.com/go-docker"

import (
	"encoding/json"
	"net/url"

	"golang.docker.com/go-docker/api/types"
	"golang.docker.com/go-docker/api/types/filters"
	"golang.docker.com/go-docker/api/types/swarm"
	"golang.org/x/net/context"
)

// ServiceList returns the list of services.
func (cli *Client) ServiceList(ctx context.Context, options types.ServiceListOptions) ([]swarm.Service, error) {
	query := url.Values{}

	if options.Filters.Len() > 0 {
		filterJSON, err := filters.ToParam(options.Filters)
		if err != nil {
			return nil, err
		}

		query.Set("filters", filterJSON)
	}

	resp, err := cli.get(ctx, "/services", query, nil)
	if err != nil {
		return nil, err
	}

	var services []swarm.Service
	err = json.NewDecoder(resp.body).Decode(&services)
	ensureReaderClosed(resp)
	return services, err
}
