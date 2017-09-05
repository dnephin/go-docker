package docker // import "golang.docker.com/go-docker"

import (
	"net/url"

	"golang.docker.com/go-docker/api/types/versions"
	"golang.org/x/net/context"
)

// VolumeRemove removes a volume from the docker host.
func (cli *Client) VolumeRemove(ctx context.Context, volumeID string, force bool) error {
	query := url.Values{}
	if versions.GreaterThanOrEqualTo(cli.version, "1.25") {
		if force {
			query.Set("force", "1")
		}
	}
	resp, err := cli.delete(ctx, "/volumes/"+volumeID, query, nil)
	ensureReaderClosed(resp)
	return err
}
