package docker // import "golang.docker.com/go-docker"

import (
	"io"
	"net/url"
	"time"

	"golang.org/x/net/context"

	"golang.docker.com/go-docker/api/types"
	timetypes "golang.docker.com/go-docker/api/types/time"
)

// ServiceLogs returns the logs generated by a service in an io.ReadCloser.
// It's up to the caller to close the stream.
func (cli *Client) ServiceLogs(ctx context.Context, serviceID string, options types.ContainerLogsOptions) (io.ReadCloser, error) {
	query := url.Values{}
	if options.ShowStdout {
		query.Set("stdout", "1")
	}

	if options.ShowStderr {
		query.Set("stderr", "1")
	}

	if options.Since != "" {
		ts, err := timetypes.GetTimestamp(options.Since, time.Now())
		if err != nil {
			return nil, err
		}
		query.Set("since", ts)
	}

	if options.Timestamps {
		query.Set("timestamps", "1")
	}

	if options.Details {
		query.Set("details", "1")
	}

	if options.Follow {
		query.Set("follow", "1")
	}
	query.Set("tail", options.Tail)

	resp, err := cli.get(ctx, "/services/"+serviceID+"/logs", query, nil)
	if err != nil {
		return nil, err
	}
	return resp.body, nil
}
