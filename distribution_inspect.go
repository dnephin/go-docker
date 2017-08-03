package docker // import "golang.docker.io/go-docker"

import (
	"encoding/json"
	"net/url"

	registrytypes "golang.docker.io/go-docker/api/types/registry"
	"golang.org/x/net/context"
)

// DistributionInspect returns the image digest with full Manifest
func (cli *Client) DistributionInspect(ctx context.Context, image, encodedRegistryAuth string) (registrytypes.DistributionInspect, error) {
	var headers map[string][]string

	if encodedRegistryAuth != "" {
		headers = map[string][]string{
			"X-Registry-Auth": {encodedRegistryAuth},
		}
	}

	// Contact the registry to retrieve digest and platform information
	var distributionInspect registrytypes.DistributionInspect
	resp, err := cli.get(ctx, "/distribution/"+image+"/json", url.Values{}, headers)
	if err != nil {
		return distributionInspect, err
	}

	err = json.NewDecoder(resp.body).Decode(&distributionInspect)
	ensureReaderClosed(resp)
	return distributionInspect, err
}