package docker // import "golang.docker.io/go-docker"

import (
	"io"
	"net/http"
	"net/url"

	"golang.docker.io/go-docker/api/types"
	"golang.org/x/net/context"
)

// PluginCreate creates a plugin
func (cli *Client) PluginCreate(ctx context.Context, createContext io.Reader, createOptions types.PluginCreateOptions) error {
	headers := http.Header(make(map[string][]string))
	headers.Set("Content-Type", "application/x-tar")

	query := url.Values{}
	query.Set("name", createOptions.RepoName)

	resp, err := cli.postRaw(ctx, "/plugins/create", query, createContext, headers)
	if err != nil {
		return err
	}
	ensureReaderClosed(resp)
	return err
}
