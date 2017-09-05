package docker // import "golang.docker.com/go-docker"

import (
	"io"
	"net/url"
	"time"

	"golang.org/x/net/context"

	"golang.docker.com/go-docker/api/types"
	timetypes "golang.docker.com/go-docker/api/types/time"
)

// ContainerLogs returns the logs generated by a container in an io.ReadCloser.
// It's up to the caller to close the stream.
//
// The stream format on the response will be in one of two formats:
//
// If the container is using a TTY, there is only a single stream (stdout), and
// data is copied directly from the container output stream, no extra
// multiplexing or headers.
//
// If the container is *not* using a TTY, streams for stdout and stderr are
// multiplexed.
// The format of the multiplexed stream is as follows:
//
//    [8]byte{STREAM_TYPE, 0, 0, 0, SIZE1, SIZE2, SIZE3, SIZE4}[]byte{OUTPUT}
//
// STREAM_TYPE can be 1 for stdout and 2 for stderr
//
// SIZE1, SIZE2, SIZE3, and SIZE4 are four bytes of uint32 encoded as big endian.
// This is the size of OUTPUT.
//
// You can use github.com/docker/docker/pkg/stdcopy.StdCopy to demultiplex this
// stream.
func (cli *Client) ContainerLogs(ctx context.Context, container string, options types.ContainerLogsOptions) (io.ReadCloser, error) {
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

	resp, err := cli.get(ctx, "/containers/"+container+"/logs", query, nil)
	if err != nil {
		return nil, err
	}
	return resp.body, nil
}
