package docker // import "golang.docker.com/go-docker"

import "golang.org/x/net/context"

// NetworkRemove removes an existent network from the docker host.
func (cli *Client) NetworkRemove(ctx context.Context, networkID string) error {
	resp, err := cli.delete(ctx, "/networks/"+networkID, nil, nil)
	ensureReaderClosed(resp)
	return err
}
