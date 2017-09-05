package docker // import "golang.docker.com/go-docker"

import (
	"net/url"
	"strconv"

	"golang.docker.com/go-docker/api/types/swarm"
	"golang.org/x/net/context"
)

// ConfigUpdate attempts to update a Config
func (cli *Client) ConfigUpdate(ctx context.Context, id string, version swarm.Version, config swarm.ConfigSpec) error {
	query := url.Values{}
	query.Set("version", strconv.FormatUint(version.Index, 10))
	resp, err := cli.post(ctx, "/configs/"+id+"/update", query, config, nil)
	ensureReaderClosed(resp)
	return err
}
