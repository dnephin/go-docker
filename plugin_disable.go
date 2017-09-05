package docker // import "golang.docker.com/go-docker"

import (
	"net/url"

	"golang.docker.com/go-docker/api/types"
	"golang.org/x/net/context"
)

// PluginDisable disables a plugin
func (cli *Client) PluginDisable(ctx context.Context, name string, options types.PluginDisableOptions) error {
	query := url.Values{}
	if options.Force {
		query.Set("force", "1")
	}
	resp, err := cli.post(ctx, "/plugins/"+name+"/disable", query, nil, nil)
	ensureReaderClosed(resp)
	return err
}
