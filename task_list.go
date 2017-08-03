package docker // import "golang.docker.io/go-docker"

import (
	"encoding/json"
	"net/url"

	"golang.docker.io/go-docker/api/types"
	"golang.docker.io/go-docker/api/types/filters"
	"golang.docker.io/go-docker/api/types/swarm"
	"golang.org/x/net/context"
)

// TaskList returns the list of tasks.
func (cli *Client) TaskList(ctx context.Context, options types.TaskListOptions) ([]swarm.Task, error) {
	query := url.Values{}

	if options.Filters.Len() > 0 {
		filterJSON, err := filters.ToParam(options.Filters)
		if err != nil {
			return nil, err
		}

		query.Set("filters", filterJSON)
	}

	resp, err := cli.get(ctx, "/tasks", query, nil)
	if err != nil {
		return nil, err
	}

	var tasks []swarm.Task
	err = json.NewDecoder(resp.body).Decode(&tasks)
	ensureReaderClosed(resp)
	return tasks, err
}
